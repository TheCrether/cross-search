package desktop

import (
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
)

var xdgDataDirs = func() []string {
	for _, envVar := range os.Environ() {
		if strings.Contains(envVar, "XDG_DATA_DIRS") {
			dirs := strings.Replace(envVar, "XDG_DATA_DIRS=", "", 1)
			return strings.Split(dirs, ":")
		}
	}
	return []string{}
}()

var home, _ = os.UserHomeDir()
var additionalDirs = []string{
	path.Join(home, ".local", "share"),
}
var desktop = path.Join(home, "Desktop")

var dirs = func() []string {
	var temp []string
	for _, dir := range append(xdgDataDirs, additionalDirs...) {
		tempDir := path.Join(dir, "applications")
		if _, err := os.Stat(tempDir); !os.IsNotExist(err) {
			temp = append(temp, tempDir)
		}
	}
	temp = append(temp, desktop)
	return temp
}

var options = Options{
	EXT:  ".desktop",
	DIRS: dirs(),
}

func parse(input string) (Result, error) {
	temp := make(map[string]interface{})

	disqualifying := []string{"NoDisplay", "Hidden"}
	for _, key := range disqualifying {
		regex := getRegex(key)
		match := regex.FindString(input)
		if match != "" {
			return Result{
				Name: "",
				Icon: "",
				Exec: nil,
			}, errors.New("wont be displayed")
		}
	}

	keys := []string{
		"Name",
		"Exec",
		"Icon",
	}

	for _, key := range keys {
		regex := getRegex(key)
		match := regex.FindString(input)
		if match != "" {
			temp[key] = strings.Replace(match, key+"=", "", 1)
		}
	}

	var result Result

	err := mapstructure.Decode(temp, &result)
	if err != nil {
		log.Println("asd")
		return Result{
			Name: "",
			Icon: "",
			Exec: nil,
		}, errors.New("wont be displayed")
	}

	return result, nil
}

func getRegex(key string) *regexp.Regexp {
	str := fmt.Sprintf("(?m)^%v=(.+)", key)
	regex, _ := regexp.Compile(str)
	return regex
}