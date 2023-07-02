package desktop

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"

	"github.com/mitchellh/mapstructure"
)

// TODO add icon finding https://wiki.archlinux.org/title/desktop_entries#Icon_path

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
				Name:     "",
				Icon:     "",
				Exec:     "",
				ExecFunc: func() {},
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
			Name:     "",
			Icon:     "",
			Exec:     "",
			ExecFunc: func() {},
		}, errors.New("wont be displayed")
	}

	result.ExecFunc = func() {
		log.Println("he")
		split := strings.Split(result.Exec, " ")
		end := len(split)
		_ = exec.Command(split[0], split[1:end]...).Run()
	}

	return result, nil
}

func getRegex(key string) *regexp.Regexp {
	str := fmt.Sprintf("(?m)^%v=(.+)", key)
	regex, _ := regexp.Compile(str)
	return regex
}
