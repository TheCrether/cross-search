package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func onChanged() {
	text := gSearch.Text()

	escaped := regexp.QuoteMeta(text)

	contains, err := regexp.Compile("(?i).*" + escaped + ".*")
	errorCheck(err)
	listModel := gList.ObserveChildren()
	n := listModel.NItems()
	matched := 0
	for i := uint(0); i < n; i++ {
		item := listModel.Item(i)
		widget := item.Cast().(*gtk.ListBoxRow)
		property := widget.ObjectProperty("name")
		name := strings.Split(property.(string), "\\")[0]
		// begin := beginning.MatchString(name)
		// if begin {
		// 	widget.SetVisible(true)
		// 	widget.Show()
		// 	continue
		// }

		inner := contains.MatchString(name)
		if inner {
			widget.SetVisible(true)
			widget.Show()
			matched++
			continue
		}

		widget.SetVisible(false)
		widget.Hide()
	}

	// TODO add style/class change to entry when no results are found

	calculateNewSize(matched)
}

func onListItemActivated() {
	prop := gList.SelectedRow().ObjectProperty("name")
	split := strings.Split(prop.(string), "\\")
	index, _ := strconv.ParseInt(split[len(split)-1], 10, 32)
	results[index].ExecFunc()
}

func onMainWindowDestroy() {
	log.Println("onMainWindowDestroy")
}
