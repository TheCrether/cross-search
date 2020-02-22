#!/usr/bin/env python3

from gi.repository import Gtk
import sys

if len(sys.argv) != 2:
    exit()

icon_name = sys.argv[1]
if icon_name:
    theme = Gtk.IconTheme.get_default()
    found_icons = set()
    for res in range(0, 512, 2):
        icon = theme.lookup_icon(icon_name, res, 0)
        if icon:
            found_icons.add(icon.get_filename())

    if found_icons:
        print("\n".join(found_icons))
    else:
        print(icon_name, "was not found")
