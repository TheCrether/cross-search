package main

import (
	"github.com/TheCrether/cross-search/desktop"
	"github.com/gotk3/gotk3/gdk"
	"log"
	"os"
	"regexp"
	"strings"
	"unsafe"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

var (
	gApplication *gtk.Application
	gEntry       *gtk.Entry
	gList        *gtk.ListBox
	gBuilder     *gtk.Builder
	gWin         *gtk.Window
	iconRegex, _ = regexp.Compile(".*.(svg|png|xpm|gif|ico)$")
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

	// Get the GtkBuilder UI definition in the glade file.
	builder, err := gtk.BuilderNewFromFile("ui/main.glade")
	errorCheck(err)
	gBuilder = builder

	// Map the handlers to callback functions, and connect the signals
	// to the Builder.
	signals := map[string]interface{}{
		"on_main_window_destroy": onMainWindowDestroy,
		"on_search_changed":      onChanged,
		"list_item_clicked":      onListItemClick,
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


	results := desktop.GetResults()

	for _, result := range results {
		gList.Add(makeListItem(result))
	}

	win.ShowAll()

	gList.SetSizeRequest(496, gList.GetAllocatedHeight())
}

func makeListItem(result desktop.Result) *gtk.ListBoxRow {
	builder, err := gtk.BuilderNewFromFile("ui/main.glade")
	errorCheck(err)

	row, err := gtk.ListBoxRowNew()
	errorCheck(err)
	alloc := gtk.Allocation{}
	alloc.SetWidth(480)
	alloc.SetHeight(50)
	row.Connect("show", func() {
		row.SetAllocation(&alloc)
	})
	row.SetSizeRequest(496, 48)
	//row.SetAllocation(&alloc)

	boxObj, err := builder.GetObject("list_box")
	errorCheck(err)
	box, err := isBox(boxObj)
	errorCheck(err)

	// get the types of a Label and of a Image
	label, err := gtk.LabelNew("s")
	errorCheck(err)
	labelType := label.TypeFromInstance()
	img, err := gtk.ImageNew()
	errorCheck(err)
	imgType := img.TypeFromInstance()

	// get all children and move through each
	box.GetChildren().Foreach(func(item interface{}) {
		//check if it is a label or a image
		if item.(*gtk.Widget).IsA(labelType) {
			label := (*gtk.Label)(unsafe.Pointer(item.(*gtk.Widget)))
			label.SetText(result.Name)
		} else if item.(*gtk.Widget).IsA(imgType) {
			image := (*gtk.Image)(unsafe.Pointer(item.(*gtk.Widget)))

			if strings.Contains(result.Icon, "import") {
				log.Println(result.Icon)
			}

			if iconRegex.Match([]byte(result.Icon)) {
				//image.SetFromFile(result.Icon)
				file, _ := gdk.PixbufNewFromFileAtSize(result.Icon, 36, 36)
				image.SetFromPixbuf(file)
			} else {
				image.SetFromIconName(result.Icon, gtk.ICON_SIZE_SMALL_TOOLBAR)
			}
			image.SetSizeRequest(48,48)
		}
	})

	row.Add(box)

	return row
}

// onMainWindowDestory is the callback that is linked to the
// on_main_window_destroy handler. It is not required to map this,
// and is here to simply demo how to hook-up custom callbacks.
func onMainWindowDestroy() {
	log.Println("onMainWindowDestroy")
}
