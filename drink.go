package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"github.com/shurcooL/trayhost"
)

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

	menuItems := createTray(configFilePath)

	go notify(configFilePath, configIconPath, OS)

	// Load tray icon
	iconData, err := ioutil.ReadFile(configIconPath)
	if err != nil {
		return
	}
	trayhost.Initialize("Water Reminder", iconData, menuItems)
	trayhost.EnterLoop()

}
