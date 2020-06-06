package main

import (
	"bufio"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/deckarep/gosx-notifier"
	"github.com/gen2brain/beeep"
	"github.com/shurcooL/trayhost"
)

func findConfig(config string) bool {
	if _, err := os.Stat(filepath.Join(config, "water-reminder")); os.IsNotExist(err) {
		return false
	}
	return true
}

func downloadFile(URL, fileName string) error {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	//Create a empty file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the fiel
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	return nil
}

func writeDelay(file, delay string) {
	config, _ := os.Create(file)
	defer config.Close()
	config.WriteString(delay)
}

func readDelay(configFilePath string) int {
	file, err := os.Open(configFilePath)
	if err != nil {
		return 30
	}
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	delay := scanner.Text()
	//Convert the string read from config to int
	minutes, err := strconv.Atoi(string(delay))
	if err != nil {
		return 30
	}
	return minutes
}

func createTray(configFilePath string) []trayhost.MenuItem {
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
	return menuItems
}

func notify(config, icon, os string) {
	//Send first notification
	message := "Start drinking now"
	if os == "linux" {
		beeep.Alert("Drink!", message, icon)
	} else {
		note := gosxnotifier.NewNotification(message)
		note.Title = "Drink!"
		note.AppIcon = icon
		note.Push()
	}

	delay := readDelay(config)

	//While true send notifications sleeping every delay minutes
	for {
		time.Sleep(time.Duration(delay) * time.Minute)
		message := "You haven't been drinking for " + strconv.Itoa(delay) + " minutes"
		if os == "linux" {
			beeep.Alert("Drink!", message, icon)
		} else {
			note := gosxnotifier.NewNotification(message)
			note.Title = "Drink!"
			note.AppIcon = icon
			note.Push()
		}
	}
}
