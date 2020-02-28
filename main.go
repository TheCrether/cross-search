package main

import (
	"errors"
	"log"
	"os"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

var (
	gApplication  *gtk.Application
	gEntry        *gtk.Entry
	gStyleContext *gtk.StyleContext
)

const appID = "at.thecrether.cross-search"

func main() {
	// Create a new application.
	application, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	errorCheck(err)

	// Connect function to application startup event, this is not required.
	application.Connect("startup", func() {
		log.Println("application startup")
	})

	// Connect function to application activate event
	application.Connect("activate", func() {
		onActivate(application)
	})

	// Connect function to application shutdown event, this is not required.
	application.Connect("shutdown", func() {
		log.Println("application shutdown")
	})

	// Launch the application
	application.Run(os.Args)
}

func onActivate(application *gtk.Application) {
	log.Println("application activate")

	// Get the GtkBuilder UI definition in the glade file.
	builder, err := gtk.BuilderNewFromFile("ui/main.glade")
	errorCheck(err)

	// Map the handlers to callback functions, and connect the signals
	// to the Builder.
	signals := map[string]interface{}{
		"on_main_window_destroy": onMainWindowDestroy,
		"on_search_changed":      onChanged,
	}

	builder.ConnectSignals(signals)

	// Get the object with the id of "main_window".
	obj, err := builder.GetObject("main_window")
	errorCheck(err)

	// Verify that the object is a pointer to a gtk.ApplicationWindow.
	win, err := isWindow(obj)
	errorCheck(err)

	// Show the Window and all of its components.
	win.Show()

	screen := win.GetScreen()

	provider, err := gtk.CssProviderNew()
	errorCheck(err)

	provider.LoadFromPath("./ui/main.css")

	gtk.AddProviderForScreen(screen, provider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

	application.AddWindow(win)

	searchEntryObj, err := builder.GetObject("search")
	errorCheck(err)
	gEntry, err = isEntry(searchEntryObj)
	errorCheck(err)
}

func onChanged() {
	text, err := gEntry.GetText()
	errorCheck(err)
	log.Println(text)
}

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

func errorCheck(e error) {
	if e != nil {
		// panic for any errors.
		log.Panic(e)
	}
}

// onMainWindowDestory is the callback that is linked to the
// on_main_window_destroy handler. It is not required to map this,
// and is here to simply demo how to hook-up custom callbacks.
func onMainWindowDestroy() {
	log.Println("onMainWindowDestroy")
}
