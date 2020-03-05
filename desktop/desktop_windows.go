package desktop

import (
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"log"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"
)

var home, _ = os.UserHomeDir()
var desktop = path.Join(home, "Desktop")

var dirs = func() []string {
	var temp []string
	for _, dir := range append(xdgDataDirs) {
		//tempDir := path.Join(dir, "applications")
		if _, err := os.Stat(dir); !os.IsNotExist(err) {
			temp = append(temp, tempDir)
		}
	}
	temp = append(temp, desktop)
	return temp
}

var options = Options{
	EXT:  ".lnk",
	DIRS: dirs(),
}

func parse(input string) (Result, error) {


	return result, nil
}

func getRegex(key string) *regexp.Regexp {
	str := fmt.Sprintf("(?m)^%v=(.+)", key)
	regex, _ := regexp.Compile(str)
	return regex
}