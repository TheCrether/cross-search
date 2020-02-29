package main

import (
	"errors"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"log"
)

func isWindow(obj glib.IObject) (*gtk.Window, error) {
	// Make type assertion (as per gtk.go).
	if win, ok := obj.(*gtk.Window); ok {
		return win, nil
	}
	return nil, errors.New("not a *gtk.Window")
}

func isEntry(obj glib.IObject) (*gtk.Entry, error) {
	// Make type assertion (as per gtk.go).
	if search, ok := obj.(*gtk.Entry); ok {
		return search, nil
	}
	return nil, errors.New("not a *gtk.Entry")
}

func isListBox(obj glib.IObject) (*gtk.ListBox, error) {
	// Make type assertion (as per gtk.go).
	if search, ok := obj.(*gtk.ListBox); ok {
		return search, nil
	}
	return nil, errors.New("not a *gtk.ListBox")
}

func isBox(obj glib.IObject) (*gtk.Box, error) {
	// Make type assertion (as per gtk.go).
	if search, ok := obj.(*gtk.Box); ok {
		return search, nil
	}
	return nil, errors.New("not a *gtk.Box")
}

func isLabel(obj glib.IObject) (*gtk.Label, error) {
	// Make type assertion (as per gtk.go).
	if search, ok := obj.(*gtk.Label); ok {
		return search, nil
	}
	return nil, errors.New("not a *gtk.Label")
}

func isImage(obj glib.IObject) (*gtk.Image, error) {
	// Make type assertion (as per gtk.go).
	if search, ok := obj.(*gtk.Image); ok {
		return search, nil
	}
	return nil, errors.New("not a *gtk.Image")
}

func isScrolledWindow(obj glib.IObject) (*gtk.ScrolledWindow, error) {
	// Make type assertion (as per gtk.go).
	if search, ok := obj.(*gtk.ScrolledWindow); ok {
		return search, nil
	}
	return nil, errors.New("not a *gtk.ScrolledWindow")
}

func errorCheck(e error) {
	if e != nil {
		// panic for any errors.
		log.Panic(e)
	}
}