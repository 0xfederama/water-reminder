# Water Reminder

[![Build Status](https://travis-ci.com/0xfederama/water-reminder.svg?branch=master)](https://travis-ci.com/0xfederama/water-reminder)

:droplet: :droplet: :droplet:

Remember to drink every 20 minutes with this simple script.

Since the app is very small, you can simply use the binaries without recompiling it.

For now the app displays the icon only for linux

### Installation

##### Linux

- Clone the repository wherever you want
- Open "Startup Applications", press "Add" and in the command section type `path/to/water-reminder/bin/drink_linux`. Give the app the name you want.
- If you don't want the app to open at startup, just go in the terminal when you want to execute it and type `/path/to/water-reminder/bin/drink_linux`

##### MacOS

- Clone the repository wherever you want
- Only for the first time, you have to cd into `path/to/water-reminder/bin/drink_mac` and type `./drink_mac`. Then go into preferences, security and allow the app to be executed, since it is not trusted by Apple.
- Open preferences, users, [your user], login items and add the executable drink_mac
- If you don't want the app to open at startup, just go in the terminal when you want to execute it and type `/path/to/water-reminder/bin/drink_mac`
