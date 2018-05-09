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

import (
	"fmt"
	"log"
)

func LogError(msg string) {
	line := fmt.Sprintf("%s %s", ColorRed("ERROR"), msg)
	log.Println(line)
}

func LogInfo(msg string) {
	line := fmt.Sprintf("%s %s", ColorGreen("INFO"), msg)
	log.Println(line)
}

func LogWarn(msg string) {
	line := fmt.Sprintf("%s %s", ColorYellow("WARNING"), msg)
	log.Println(line)
}

func LogDebug(msg string) {
	line := fmt.Sprintf("%s %s", ColorBlue("DEBUG"), msg)
	log.Println(line)
}

func ColorRed(value string) string {
	return fmt.Sprintf("\033[0;31m%s\033[0m", value)
}

func ColorGreen(value string) string {
	return fmt.Sprintf("\033[0;32m%s\033[0m", value)
}

func ColorYellow(value string) string {
	return fmt.Sprintf("\033[0;33m%s\033[0m", value)
}

func ColorBlue(value string) string {
	return fmt.Sprintf("\033[0;34m%s\033[0m", value)
}
