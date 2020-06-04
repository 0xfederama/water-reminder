package main

import (
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"

	"github.com/gen2brain/beeep"
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

	//Send first notification after 5 seconds
	time.Sleep(5 * time.Second)
	message := "Start drinking now"
	beeep.Alert("Drink!", message, configIconPath)

	delay := readDelay(configFilePath)

	//While true send notifications sleeping every delay minutes
	for {
		time.Sleep(time.Duration(delay) * time.Minute)
		message := "You haven't been drinking for " + strconv.Itoa(delay) + " minutes"
		beeep.Alert("Drink!", message, configIconPath)
	}

}
