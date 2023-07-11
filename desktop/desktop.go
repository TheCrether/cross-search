package desktop

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/TheCrether/cross-search/utils"
)

// Options the desktop options
type Options struct {
	EXT  string
	DIRS []string
}

type Result struct {
	Name     string
	Icon     string
	Exec     string
	ExecFunc func()
}

func ContainsResult(array *[]Result, result *Result) bool {
	for _, r := range *array {
		if result.Name == r.Name && result.Icon == r.Icon {
			return true
		}
	}
	return false
}

func GetResults() []Result {
	return getApplications()
}

func getApplications() []Result {
	var results []Result
	for _, dir := range options.DIRS {
		filepath.Walk(dir, visit(&results))
	}
	return results
}

var (
	dirsVisited  []string
	alreadyAdded []string
	cnt          = make(map[string]int)
)

func visit(results *[]Result) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && !utils.Contains(dirsVisited, path) {
			dirsVisited = append(dirsVisited, path)
			filepath.Walk(path, visit(results))
		} else if filepath.Ext(path) == options.EXT {
			data, err := ioutil.ReadFile(path)
			if bad := errorCheck(err); bad {
				return nil
			}
			result, err := parse(string(data))
			if err != nil {
				return nil
			}
			if utils.Contains(alreadyAdded, result.Name) {
				if ContainsResult(results, &result) {
					return nil
				}
				cnt[result.Name]++
				result.Name += " (" + string(cnt[result.Name]) + ")"
			}
			alreadyAdded = append(alreadyAdded, result.Name)
			*results = append(*results, result)
		}
		return nil
	}
}

func errorCheck(e error) bool {
	if e != nil {
		// panic for any errors.
		log.Print(e)
		return true
	}
	return false
}
