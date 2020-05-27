package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func main() {

	time.Sleep(5 * time.Second)

	//Path to the image used for the notification
	exPath, err := os.Executable()
	if err != nil {
		return
	}
	exDir := filepath.Dir(exPath)
	iconPath := "../resources/water-glass.png"
	icon := filepath.Join(exDir, iconPath)
	configPath := "../resources/config.txt"
	config := filepath.Join(exDir, configPath)

	//Executing first command
	Alert("Drink!!!", "Start drinking now", icon)

	//Read from config file in ../resources how many minutes to wait between two notifications
	file, err := os.Open(config)
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	delay := scanner.Text()
	//Convert the string read from config in int
	minutes, err := strconv.Atoi(string(delay))
	if err != nil {
		return
	}

	//Wait *minutes* minutes before sending another notification
	for {
		time.Sleep(time.Duration(minutes) * time.Minute)
		Alert("Drink!!!", "You haven't been drinking for "+delay+" minutes", icon)
	}

}
