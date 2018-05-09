/**
    This file is part of gitlab-crawler.

    Gitlab-crawler is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    Gitlab-crawler is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with gitlab-crawler.  If not, see <http://www.gnu.org/licenses/>.
**/

package crawlconfig

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/fatih/camelcase"
	"github.com/tinyzimmer/gitlab-crawler/logging"
)

const ENV_CONFIG_PREFIX = "CRAWLER"
const DEFAULT_GITLAB_TOKEN = ""
const DEFAULT_GITLAB_SERVER = "http://127.0.0.1"
const DEFAULT_COMPRESSION_TYPE = "gzip"
const DEFAULT_DEBUG = false

type CrawlConfiguration struct {
	GitlabServer    string   `json:"gitlabServer"`
	GitlabToken     string   `json:"gitlabToken"`
	GitlabGroups    []string `json:"gitlabGroups"`
	CompressionType string   `json:"cacheCompressionType"`
	RecurseBranches bool     `json:"recurseBranches"`
	TestMode        bool     `json:"testMode"`
	Debug           bool     `json:"crawlerDebug"`
}

func (c CrawlConfiguration) DumpConfig() (config string) {
	cbytes, _ := json.Marshal(c)
	config = string(cbytes)
	return
}

func GetConfig() (config CrawlConfiguration) {
	logging.LogInfo("Retrieving default configurations")
	config.GitlabServer = DEFAULT_GITLAB_SERVER
	config.GitlabToken = DEFAULT_GITLAB_TOKEN
	config.GitlabGroups = make([]string, 0)
	config.CompressionType = DEFAULT_COMPRESSION_TYPE
	config.Debug = DEFAULT_DEBUG
	logging.LogInfo("Parsing configurations from environment")
	config = getEnvConfigs(config)
	if config.Debug {
		logging.LogDebug(config.DumpConfig())
	}
	return
}

func getEnvConfigs(config CrawlConfiguration) (parsedConfig CrawlConfiguration) {
	var envConfig string
	var envConfigValue string
	parsedConfig = config
	s := reflect.ValueOf(&parsedConfig).Elem()
	typeOfConfig := s.Type()
	for i := 0; i < s.NumField(); i++ {
		envConfig = ENV_CONFIG_PREFIX
		f := s.Field(i)
		for _, val := range camelcase.Split(fmt.Sprint(typeOfConfig.Field(i).Name)) {
			envConfig = envConfig + "_" + strings.ToUpper(val)
		}
		envConfigValue = os.Getenv(envConfig)
		if envConfigValue == "" {
			continue
		} else if fmt.Sprint(f.Type()) == "bool" {
			f.SetBool(checkEnvBool(envConfigValue))
		} else if fmt.Sprint(f.Type()) == "string" {
			f.SetString(envConfigValue)
		} else if typeOfConfig.Field(i).Name == "GitlabGroups" {
			f.Set(checkEnvSlice(envConfigValue))
		}
	}
	return
}

func checkEnvSlice(value string) (values reflect.Value) {
	parsed := strings.Split(value, ",")
	values = reflect.ValueOf(parsed)
	return
}

func checkEnvBool(value string) (res bool) {
	res = false
	if value == "1" || strings.ToLower(value) == "true" {
		res = true
	}
	return
}
