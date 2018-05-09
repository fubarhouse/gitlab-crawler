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
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/tinyzimmer/gitlab-crawler/crawlconfig"
)

type TestCredsResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	CreatedAt    int64  `json:"created_at"`
}

type TestAccessTokenResponse struct {
	Id            int64    `json:"id"`
	Revoked       bool     `json:"revoked"`
	Scopes        []string `json:"scopes"`
	Token         string   `json:"token"`
	Active        bool     `json:"active"`
	Impersonation bool     `json:"impersonation"`
	Name          string   `json:"name"`
	CreatedAt     string   `json:"created_at"`
	ExpiresAt     string   `json:"expires_at"`
}

const BOGUS_ENDPOINT = "https://127.0.0.1:60000"

func TestBadGitlabAddress(t *testing.T) {
	os.Clearenv()
	os.Setenv("CRAWLER_GITLAB_SERVER", BOGUS_ENDPOINT)
	log.Println("TEST: Testing bogus endpoint")
	config := crawlconfig.GetConfig()
	err := RunEngine(config)
	if err == nil {
		t.Errorf("Allowed bogus endpoint of %s", BOGUS_ENDPOINT)
	}
}

func TestBadRunEngine(t *testing.T) {
	os.Clearenv()
	log.Println("TEST: Testing bad credentials")
	os.Setenv("CRAWLER_GITLAB_TOKEN", "")
	err := RunEngine(crawlconfig.GetConfig())
	if err == nil {
		t.Errorf("Allowed unauthenticated query")
	}
}

func TestGoodRunEngine(t *testing.T) {
	os.Clearenv()
	otoken, err := GetTestOauthToken()
	if err != nil {
		t.Errorf(err.Error())
	}
	atoken, err := GetTestAccessToken(otoken)
	if err != nil {
		t.Errorf(err.Error())
	}
	os.Setenv("CRAWLER_GITLAB_TOKEN", atoken)
	err = RunEngine(crawlconfig.GetConfig())
	if err != nil {
		fmt.Println(err.Error())
		t.Errorf("Failed to create engine with access token, did you run docker-compose?")
	}
}

func GetTestOauthToken() (token string, err error) {
	log.Println("TEST: Retrieving test gitlab oauth token")
	urlString := "http://localhost/oauth/token"
	resp, err := http.PostForm(urlString, url.Values{
		"grant_type": {"password"},
		"username":   {"root"},
		"password":   {"testpassword"},
	})
	if err != nil {
		return
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	credsResponse := &TestCredsResponse{}
	err = decoder.Decode(&credsResponse)
	if err == nil {
		token = credsResponse.AccessToken
		log.Println(fmt.Sprintf("TEST: Succesfully retrieved test gitlab oauth token: %s", token))
	}
	return
}

func GetTestAccessToken(otoken string) (token string, err error) {
	client := &http.Client{}
	log.Println("TEST: Retrieving test impersonation token")
	urlString := "http://localhost/api/v4/users/1/impersonation_tokens"
	var jsonString = []byte(`{"name": "test", "scopes": ["api"]}`)
	req, err := http.NewRequest("POST", urlString, bytes.NewBuffer(jsonString))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", otoken))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	credsResponse := &TestAccessTokenResponse{}
	err = decoder.Decode(&credsResponse)
	if err == nil {
		token = credsResponse.Token
		log.Println(fmt.Sprintf("TEST: Succesfully retrieved test gitlab impersonation token: %s", token))
	} else {
		log.Println(err.Error())
	}
	return
}
