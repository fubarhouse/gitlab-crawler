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
	"os"
	"strings"
	"testing"
)

const TEST_ENV_DEBUG = "CRAWLER_DEBUG"
const TEST_ENV_GROUPS = "CRAWLER_GITLAB_GROUPS"
const TEST_ENV_SERVER = "CRAWLER_GITLAB_SERVER"

func TestGetDefaultConfig(t *testing.T) {
	os.Clearenv()
	config := GetConfig()
	config.DumpConfig()
}

func TestSetBoolValue(t *testing.T) {
	os.Clearenv()
	os.Setenv(TEST_ENV_DEBUG, "true")
	config := GetConfig()
	if !config.Debug {
		t.Errorf("Failed to override boolean config value")
	}
}

func TestSetStringValue(t *testing.T) {
	os.Clearenv()
	os.Setenv(TEST_ENV_SERVER, "test_server")
	config := GetConfig()
	if config.GitlabServer != "test_server" {
		t.Errorf("Failed to override string config value")
	}
}

func TestSetGitlabGroups(t *testing.T) {

	os.Clearenv()
	os.Setenv(TEST_ENV_GROUPS, "")
	emptyGroup := make([]string, 0)
	config := GetConfig()
	if len(config.GitlabGroups) != 0 {
		t.Errorf("Bad parse of gitlab groups. Got %+v, wanted %+v", config.GitlabGroups, emptyGroup)
	}

	os.Clearenv()
	singleGroup := make([]string, 1)
	singleGroup[0] = "test_group"
	os.Setenv(TEST_ENV_GROUPS, singleGroup[0])
	config = GetConfig()
	if len(config.GitlabGroups) != len(singleGroup) || config.GitlabGroups[0] != singleGroup[0] {
		t.Errorf("Bad parse of single gitlab group. Got %+v, wanted %+v", config.GitlabGroups, singleGroup)
	}

	os.Clearenv()
	multiGroup := make([]string, 3)
	multiGroup[0] = "test_group"
	multiGroup[1] = "test_group_2"
	multiGroup[2] = "test_group_3"
	os.Setenv(TEST_ENV_GROUPS, strings.Join(multiGroup, ","))
	config = GetConfig()
	if len(config.GitlabGroups) != len(multiGroup) {
		t.Errorf("Bad length on parsed gitlab groups. Got %+v, wanted %+v", config.GitlabGroups, multiGroup)
	}
	for i, _ := range config.GitlabGroups {
		if config.GitlabGroups[i] != multiGroup[i] {
			t.Errorf("Bad parse of gitlab groups. Got %+v, wanted %+v", config.GitlabGroups, multiGroup)
		}
	}
}
