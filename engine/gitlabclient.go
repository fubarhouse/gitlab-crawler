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

package engine

import (
	"fmt"
	"strings"

	"github.com/tinyzimmer/gitlab-crawler/crawlconfig"
	"github.com/tinyzimmer/gocibot/logging"
	"github.com/xanzy/go-gitlab"
)

type RuntimeEngine struct {
	BaseUri string
	Config  crawlconfig.CrawlConfiguration
	Client  *gitlab.Client
}

func RunEngine(config crawlconfig.CrawlConfiguration) (err error) {
	newEngine := RuntimeEngine{}
	newEngine.BaseUri = fmt.Sprintf("%s/api/v4/", config.GitlabServer)
	newEngine.Client = gitlab.NewClient(nil, config.GitlabToken)
	newEngine.Client.SetBaseURL(newEngine.BaseUri)
	err = newEngine.TestAuth()
	if err != nil {
		logging.LogError(err.Error())
		return
	}
	logging.LogInfo("Validated credentials")
	return
}

func (e RuntimeEngine) TestAuth() (err error) {
	_, _, err = e.Client.Settings.GetSettings()
	if err != nil {
		if strings.Contains(err.Error(), "json") { // go-gitlab bug, remember to investigate
			err = nil
		}
		return
	}
	return
}
