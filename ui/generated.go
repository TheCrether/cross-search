package ui

import (
	_ "embed"
)

//go:embed main.css
var MainCss string

//go:embed main.ui
var MainGlade string

//go:embed row.glade
var RowGlade string
