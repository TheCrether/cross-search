#!/usr/bin/env bash

# start the inspector when opening app (GTK_DEBUG=interactive)
# https://wiki.gnome.org/Projects/GTK/Inspector
GTK_DEBUG=interactive go run .
