package main

import (
	"bufio"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"

	"github.com/deckarep/gosx-notifier"
	"github.com/tcnksm/go-latest"
	"github.com/gen2brain/beeep"
	"github.com/shurcooL/trayhost"
)

func checkVersion(version, icon string) {
	githubTag := &latest.GithubTag{
		Owner:      "0xfederama",
		Repository: "water-reminder",
		FixVersionStrFunc: latest.DeleteFrontV(),
	}

	res, _ := latest.Check(githubTag, version)
	if res.Outdated {
		beeep.Alert("Water Reminder", "You should update to version "+ res.Current +". Visit github.com/0xfederama/water-reminder", icon)
	}
}

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

func createTray(configFilePath, icon string) []trayhost.MenuItem {
	menuItems := []trayhost.MenuItem{
		{
			Title: "Set delay - 15min",
			Handler: func() {
				writeDelay(configFilePath, "15")
				sendNotif("Water Reminder", "Delay set to 15 min. Reload the app to apply changes", icon)
			},
		},
		{
			Title: "Set delay - 30min",
			Handler: func() {
				writeDelay(configFilePath, "30")
				sendNotif("Water Reminder", "Delay set to 30 min. Reload the app to apply changes", icon)
			},
		},
		{
			Title: "Set delay - 45min",
			Handler: func() {
				writeDelay(configFilePath, "45")
				sendNotif("Water Reminder", "Delay set to 45 min. Reload the app to apply changes", icon)
			},
		},
		{
			Title: "Set delay - 60min",
			Handler: func() {
				writeDelay(configFilePath, "60")
				sendNotif("Water Reminder", "Delay set to 60 min. Reload the app to apply changes", icon)
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
		note.Sound = "'default'"
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
			note.Sound = "'default'"
			note.AppIcon = icon
			note.Push()
		}
	}
}

func sendNotif(title, message, icon string) {
	if runtime.GOOS == "linux" {
		beeep.Alert(title, message, icon)
	} else {
		note := gosxnotifier.NewNotification(message)
		note.Title = title
		note.Sound = "'default'"
		note.AppIcon = icon
		note.Push()
	}
}
