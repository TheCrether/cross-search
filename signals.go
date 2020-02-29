package main

import "log"

func onChanged() {
	text, err := gEntry.GetText()
	errorCheck(err)
	log.Println(text)
}

func onListItemClick() {
	log.Println("asdsad")
}
