package main

import (
	"bufio"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"

	"github.com/gen2brain/beeep"
)

func findConfig(config string) bool {
	if _, err := os.Stat(filepath.Join(config, "water-reminder")); os.IsNotExist(err) {
		return false
	}
	return true
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

func notify(OS, configIconPath, message string) {
	if OS == "darwin" {
		osa, err := exec.LookPath("osascript")
		if err != nil {
			panic(err)
		}
		cmd := exec.Command(osa, "-e", `tell application "System Events" to display notification "`+message+`" with title "Drink!" sound name "default"`)
		cmd.Run()
	} else {
		beeep.Notify("Drink!", message, configIconPath)
		beeep.Beep(440.0, 200)
	}
}

func downloadFile(URL, fileName string) error {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
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
