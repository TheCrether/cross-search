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
	Name string
	Icon string
	Exec interface{}
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

var dirsVisited []string

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
