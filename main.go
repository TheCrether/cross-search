package main

import (
	"log"
	"math"
	"os"
	"regexp"
	"strconv"

	"github.com/TheCrether/cross-search/desktop"
	"github.com/TheCrether/cross-search/ui"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

var (
	gApplication *gtk.Application
	gSearch      *gtk.Entry
	gList        *gtk.ListBox
	gBuilder     *gtk.Builder
	gWin         *gtk.Window
	imgRegex, _  = regexp.Compile(`.*\.(svg|png|xpm|gif|ico)$`)
	results      []desktop.Result
	resultHeight = 50
	searchHeight = 60
)

const appID = "at.thecrether.cross-search"

func main() {
	// Create a new application.
	application := gtk.NewApplication(appID, gio.ApplicationFlagsNone)
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

	builder := gtk.NewBuilder()
	gBuilder = builder

	// Map the handlers to callback functions, and connect the signals
	// to the Builder.
	signals := map[string]map[string]interface{}{
		"search": {
			"changed": onChanged,
		},
		"main_window": {
			"destroy": onMainWindowDestroy,
		},
		"list": {
			"row-activated": onListItemActivated,
		},
	}

	builder.AddFromString(ui.MainGlade, len(ui.MainGlade))
	for name, f := range signals {
		obj := builder.GetObject(name)
		for signal, f := range f {
			obj.Connect(signal, f)
		}
	}

	gWin = builder.GetObject("main_window").Cast().(*gtk.Window)

	provider := gtk.NewCSSProvider()
	provider.LoadFromData(ui.MainCss)
	gtk.StyleContextAddProviderForDisplay(gdk.DisplayGetDefault(), provider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

	application.AddWindow(gWin)

	gSearch = builder.GetObject("search").Cast().(*gtk.Entry)
	gSearch.SetSizeRequest(-1, searchHeight-4)

	gList = builder.GetObject("list").Cast().(*gtk.ListBox)
	for i := range results {
		row := createRow(i)
		gList.Append(row)
		if i > 4 {
			row.SetVisible(false)
			row.Hide()
		}
	}

	calculateNewSize(len(results))
	gWin.Show()
}

// creates a Row from a desktop.Result object
func createRow(index int) *gtk.ListBoxRow {
	result := results[index]

	row := gtk.NewListBoxRow()
	box := gtk.NewBox(gtk.OrientationHorizontal, 2)

	var icon *gtk.Image

	if imgRegex.MatchString(result.Icon) {
		pixbuf, err := gdkpixbuf.NewPixbufFromFileAtSize(result.Icon, 32, 32)
		if err != nil {
			icon = gtk.NewImage()
		} else {
			icon = gtk.NewImageFromPixbuf(pixbuf)
		}
	} else {
		theme := gtk.IconThemeGetForDisplay(gdk.DisplayGetDefault())
		if theme.HasIcon(result.Icon) {
			icon = gtk.NewImageFromIconName(result.Icon)
		} else {
			icon = gtk.NewImage()
		}
	}
	icon.SetIconSize(gtk.IconSizeLarge)
	icon.SetSizeRequest(resultHeight, resultHeight)
	box.Append(icon)

	label := gtk.NewLabel(result.Name)
	label.SetHExpand(true)
	label.SetJustify(gtk.JustifyLeft)
	box.Append(label)

	w, _ := gWin.DefaultSize()
	row.SetSizeRequest(w, resultHeight)
	row.QueueResize()
	row.SetObjectProperty("name", result.Name+"\\"+strconv.FormatInt(int64(index), 10))
	row.SetChild(box)

	return row
}

func calculateNewSize(resultSize int) {
	listHeight := resultHeight * int(math.Min(4, float64(resultSize)))
	height := searchHeight + listHeight
	gList.SetSizeRequest(500, listHeight)
	gList.QueueResize()
	gWin.SetDefaultSize(500, height)
}
