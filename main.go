package main

import (
	"github.com/TheCrether/cross-search/desktop"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"log"
	"os"
	"regexp"
	"strconv"
)

var (
	gApplication *gtk.Application
	gEntry       *gtk.Entry
	gList        *gtk.ListBox
	gBuilder     *gtk.Builder
	gWin         *gtk.Window
	iconRegex, _ = regexp.Compile(".*.(svg|png|xpm|gif|ico)$")
	results      []desktop.Result
	resultHeight = 50
	searchHeight = 60
)

const appID = "at.thecrether.cross-search"

func main() {
	// Create a new application.
	application, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	errorCheck(err)
	gApplication = application

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

	results = desktop.GetResults()

	// Get the GtkBuilder UI definition in the glade file.
	builder, err := gtk.BuilderNewFromFile("ui/main.glade")
	errorCheck(err)
	gBuilder = builder

	// Map the handlers to callback functions, and connect the signals
	// to the Builder.
	signals := map[string]interface{}{
		"on_main_window_destroy": onMainWindowDestroy,
		"on_search_changed":      onChanged,
		"list_item_activated":    onListItemActivated,
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
	gWin = win

	screen := win.GetScreen()

	provider, err := gtk.CssProviderNew()
	errorCheck(err)

	provider.LoadFromPath("./ui/main.css")

	gtk.AddProviderForScreen(screen, provider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

	application.AddWindow(win)

	searchObj, err := builder.GetObject("search")
	errorCheck(err)
	gEntry, err = isEntry(searchObj)
	errorCheck(err)

	listObj, err := builder.GetObject("list")
	errorCheck(err)
	gList, err = isListBox(listObj)
	errorCheck(err)

	win.ShowAll()


	for i := range results {
		gList.Add(createRow(i))
	}
}

// creates a Row from a desktop.Result object
func createRow(index int) *gtk.ListBoxRow {
	result := results[index]
	builder, err := gtk.BuilderNewFromFile("ui/row.xml")
	errorCheck(err)

	row, err := gtk.ListBoxRowNew()
	errorCheck(err)

	boxObj, err := builder.GetObject("list_box")
	errorCheck(err)
	box, err := isBox(boxObj)
	errorCheck(err)

	nameObj, err := builder.GetObject("name")
	errorCheck(err)
	name, err := isLabel(nameObj)
	errorCheck(err)
	name.SetText(result.Name)

	iconObj, err := builder.GetObject("icon")
	errorCheck(err)
	icon, err := isImage(iconObj)
	errorCheck(err)
	if iconRegex.Match([]byte(result.Icon)) {
		file, _ := gdk.PixbufNewFromFileAtSize(result.Icon, 32, 32)
		icon.SetFromPixbuf(file)
	} else {
		icon.SetFromIconName(result.Icon, gtk.ICON_SIZE_DND)
	}
	icon.SetPixelSize(32)
	icon.SetSizeRequest(resultHeight, resultHeight)

	row.Add(box)
	w, _ := gWin.GetDefaultSize()
	row.SetSizeRequest(w, resultHeight)
	propName := result.Name+"\\"+strconv.FormatInt(int64(index), 10)
	err = row.SetProperty("name", propName)
	errorCheck(err)

	return row
}
