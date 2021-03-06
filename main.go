package main

import (
	"github.com/TheCrether/cross-search/desktop"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"log"
	"os"
	"regexp"
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
	uiBox = packr.NewBox("./ui")
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

	mainGlade, err := uiBox.FindString("main.glade")
	errorCheck(err)

	// Get the GtkBuilder UI definition in the glade file.
	//builder, err := gtk.BuilderNewFromFile("ui/main.glade")
	builder, err := gtk.BuilderNew()
	errorCheck(err)
	err = builder.AddFromString(mainGlade)
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

	css, err := uiBox.FindString("main.css")
	errorCheck(err)
	err = provider.LoadFromData(css)
	errorCheck(err)

	gtk.AddProviderForScreen(screen, provider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

	application.AddWindow(win)

	searchObj, err := builder.GetObject("search")
	errorCheck(err)
	gEntry, err = isEntry(searchObj)
	errorCheck(err)

	gList, err := gtk.ListBoxNew()
	errorCheck(err)
	_, err = gList.Connect("row-activated", onListItemActivated)
	errorCheck(err)
	gList.SetSelectionMode(gtk.SELECTION_SINGLE)

	scrolledObj, err := builder.GetObject("scrolled")
	errorCheck(err)
	scrolled, err := isScrolledWindow(scrolledObj)
	errorCheck(err)

	view, err := gtk.ViewportNew(nil, nil)
	errorCheck(err)

	view.Add(gList)
	scrolled.Add(view)

	//listObj, err := builder.GetObject("list")
	//errorCheck(err)
	//gList, err = isListBox(listObj)
	//errorCheck(err)

	win.ShowAll()

	for i := range results {
		gList.Add(createRow(i))
	}
	errorCheck(err)
	//win.ShowAll()
}

// creates a Row from a desktop.Result object
func createRow(index int) *gtk.Box {
	result := results[index]
	//builder, err := gtk.BuilderNewFromFile("ui/row.glade")
	//errorCheck(err)

	//row, err := gtk.ListBoxRowNew()
	//errorCheck(err)

	box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 2)
	errorCheck(err)

	//var icon *gtk.Image
	//
	//if iconRegex.Match([]byte(result.Icon)) {
	//	file, _ := gdk.PixbufNewFromFileAtSize(result.Icon, 32, 32)
	//	icon, err = gtk.ImageNewFromPixbuf(file)
	//	errorCheck(err)
	//} else {
	//	icon, err = gtk.ImageNewFromIconName(result.Icon, gtk.ICON_SIZE_DND)
	//	errorCheck(err)
	//}
	//icon.SetPixelSize(32)
	//icon.SetSizeRequest(resultHeight, resultHeight)
	//box.Add(icon)

	label, err := gtk.LabelNew(result.Name)
	errorCheck(err)
	box.Add(label)

	//boxObj, err := builder.GetObject("list_box")
	//errorCheck(err)
	//box, err := isBox(boxObj)
	//errorCheck(err)
	//
	//nameObj, err := builder.GetObject("name")
	//errorCheck(err)
	//name, err := isLabel(nameObj)
	//errorCheck(err)
	//name.SetText(result.Name)
	//
	//iconObj, err := builder.GetObject("icon")
	//errorCheck(err)
	//icon, err := isImage(iconObj)
	//errorCheck(err)
	//if iconRegex.Match([]byte(result.Icon)) {
	//	file, _ := gdk.PixbufNewFromFileAtSize(result.Icon, 32, 32)
	//	icon.SetFromPixbuf(file)
	//} else {
	//	icon.SetFromIconName(result.Icon, gtk.ICON_SIZE_DND)
	//}
	//icon.SetPixelSize(32)
	//icon.SetSizeRequest(resultHeight, resultHeight)

	//row.Add(box)
	//w, _ := gWin.GetDefaultSize()
	//row.SetSizeRequest(w, resultHeight)
	//propName := result.Name + "\\" + strconv.FormatInt(int64(index), 10)
	//err = row.SetProperty("name", propName)
	//errorCheck(err)

	return box
}
