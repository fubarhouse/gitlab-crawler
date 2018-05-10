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

package logging

import "testing"

func TestLogError(t *testing.T) {
	LogError("Test Error")
}

func TestLogInfo(t *testing.T) {
	LogInfo("Test Info")
}

func TestLogWarn(t *testing.T) {
	LogWarn("Test Warning")
}

func TestLogDebug(t *testing.T) {
	LogDebug("Test Debug")
}

func TestColorRed(t *testing.T) {
	LogInfo(ColorRed("TEST RED"))
}

func TestColorGreen(t *testing.T) {
	LogInfo(ColorGreen("TEST GREEN"))
}

func TestColorYellow(t *testing.T) {
	LogInfo(ColorYellow("TEST YELLOW"))
}

func TestColorBlue(t *testing.T) {
	LogInfo(ColorBlue("TEST BLUE"))
}
