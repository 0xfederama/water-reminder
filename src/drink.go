package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/shurcooL/trayhost"
)

func notify(config, icon string) {
	//Send first notification
	message := "Start drinking now"
	beeep.Alert("Drink!", message, icon)

	delay := readDelay(config)

	//While true send notifications sleeping every delay minutes
	for {
		time.Sleep(time.Duration(delay) * time.Minute)
		message := "You haven't been drinking for " + strconv.Itoa(delay) + " minutes"
		beeep.Alert("Drink!", message, icon)
	}
}

func main() {

	//Search in config path if there is the directory water-reminder
	OS := runtime.GOOS
	var configPath string
	home, _ := os.LookupEnv("HOME")
	if OS == "darwin" {
		configPath = filepath.Join(home, "Library/Application Support")
	} else {
		configPath = filepath.Join(home, ".config")
	}

	configDirPath := filepath.Join(configPath, "water-reminder")
	configFilePath := filepath.Join(configDirPath, "config.txt")
	configIconPath := filepath.Join(configDirPath, "water-glass.png")

	if !findConfig(configPath) {
		//Create config directory
		os.Mkdir(configDirPath, 0700)

		//Download icon and default config file in the new directory
		downloadFile("https://raw.githubusercontent.com/0xfederama/water-reminder/master/resources/config.txt", configFilePath)
		downloadFile("https://raw.githubusercontent.com/0xfederama/water-reminder/master/resources/water-glass.png", configIconPath)
	}

	menuItems := []trayhost.MenuItem{
		{
			Title: "Set delay 15min (reload to apply)",
			Handler: func() {
				writeDelay(configFilePath, "15")
			},
		},
		{
			Title: "Set delay 30min (reload to apply)",
			Handler: func() {
				writeDelay(configFilePath, "30")
			},
		},
		{
			Title: "Set delay 45min (reload to apply)",
			Handler: func() {
				writeDelay(configFilePath, "45")
			},
		},
		{
			Title: "Set delay 60min (reload to apply)",
			Handler: func() {
				writeDelay(configFilePath, "60")
			},
		},
		trayhost.SeparatorMenuItem(),
		{
			Title:   "Quit",
			Handler: trayhost.Exit,
		},
	}

	go notify(configFilePath, configIconPath)

	// Load tray icon
	iconData, _ := ioutil.ReadFile(configIconPath)
	trayhost.Initialize("Water Reminder", iconData, menuItems)
	trayhost.EnterLoop()

}
