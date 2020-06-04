# Water Reminder

[![Build Status](https://travis-ci.com/0xfederama/water-reminder.svg?branch=master)](https://travis-ci.com/0xfederama/water-reminder)

:droplet: :droplet: :droplet:

Remember to drink every 30 minutes with this simple script.

For now the app displays the icon only for linux.

# Installation

### Linux

- Download the latest [release](https://github.com/0xfederama/water-reminder/releases) for Linux (the simple `drink_linux` binary file) and place it wherever you want
- To make the app run at startup open "Startup Applications", press "Add" and in the command section type `path/to/drink_linux`. Give the app the name you want.
- If you don't want the app to open at startup, just go in the terminal when you want to execute it and type `/path/to/drink_linux`
- If you want to change the delay between two notifications, simply change the number of minutes in `$HOME/.config/water-reminder/config.txt`

### MacOS

- Download the latest [release](https://github.com/0xfederama/water-reminder/releases) for MacOS (`.app` file)
- Move the downloaded file to `/Applications`
- Launch the app, then go into preferences, security and allow the app to be executed, since it is not trusted by Apple.
- To make the app run at startup open preferences, users, [your user], login items and add the app to the list
- If you want to change the delay between two notifications, simply change the number of minutes in `$HOME/Library/Application Support/water-reminder/config.txt`
