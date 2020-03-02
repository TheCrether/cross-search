package main

import (
	"github.com/gotk3/gotk3/gtk"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func onChanged() {
	text, err := gEntry.GetText()
	errorCheck(err)

	escaped := regexp.QuoteMeta(text)

	beginning, err := regexp.Compile("(?i).*" + escaped + ".*")
	errorCheck(err)

	contains, err := regexp.Compile("(?i).*" + escaped + ".*")
	errorCheck(err)

	gList.GetChildren().Foreach(func(item interface{}) {
		widget := item.(*gtk.Widget)
		property, err := widget.GetProperty("name")
		errorCheck(err)
		name := strings.Split(property.(string), "\\")[0]
		begin := beginning.MatchString(name)
		if begin {
			widget.SetVisible(true)
			widget.Show()
			return
		}

		inner := contains.MatchString(name)
		if inner {
			widget.SetVisible(true)
			widget.Show()
			return
		}

		widget.SetVisible(false)
		widget.Hide()
	})
}

func onListItemActivated() {
	prop, err := gList.GetSelectedRow().GetProperty("name")
	errorCheck(err)
	split := strings.Split(prop.(string), "\\")
	index, _ := strconv.ParseInt(split[len(split)-1], 10, 32)
	results[index].ExecFunc()
}

func onMainWindowDestroy() {
	log.Println("onMainWindowDestroy")
}