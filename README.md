# Water Reminder

[![Build Status](https://travis-ci.com/0xfederama/water-reminder.svg?branch=master)](https://travis-ci.com/0xfederama/water-reminder)

:droplet: :droplet: :droplet:

Remember to drink every 30 minutes with this simple app.

For now the app displays the icon in the notifications only in linux.

# Installation

### Linux

- Download the latest [release](https://github.com/0xfederama/water-reminder/releases) for Linux (the simple `drink_linux` binary file) and place it wherever you want
- To make the app runnable from the applications grid, first launch it from the terminal, then you need to create `water-reminder.desktop` file in `$HOME/.local/share/applications` and copy this, changing [your username]
```
  [Desktop Entry]
  Name=Water Reminder
  Exec=/path/to/drink_linux
  Terminal=false
  Type=Application
  Comment=Remember to drink every 30 minutes
  Icon=/home/[your username]/.config/water-reminder/water-glass.png
```
- To make the app run at startup (using `water-reminder.desktop` and if you have `gnome-tweak-tools` installed) you can open Tweaks and add Water Reminder to the startup applications. Otherwise, if you didn't create `water-reminder.desktop` or if you don't have `gnome-tweak-tools`, open "Startup Applications", press "Add" and in the command section type `path/to/drink_linux`. Give the app the name you want

### MacOS

- Download the latest [release](https://github.com/0xfederama/water-reminder/releases) for MacOS (`.dmg` file)
- Move `Water Reminder.app` file in Applications
- Launch the app, then go into preferences, security and allow the app to be executed, since it is not trusted by Apple.
- To make the app run at startup open preferences, users, [your user], login items and add the app to the list
