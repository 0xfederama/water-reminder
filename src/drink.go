package main

import (
	"os"
	"path/filepath"
	"time"
)

func main() {

	time.Sleep(5 * time.Second)

	//Path to the image used for the notification
	exPath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exDir := filepath.Dir(exPath)
	iconPath := "../resources/water-glass.png"
	icon := filepath.Join(exDir, iconPath)

	//Executing first command
	Alert("Drink!!!", "Start drinking now", icon)

	//Wait 30 minutes before sending another notification
	for {
		time.Sleep(30 * time.Minute)
		Alert("Drink!!!", "You haven't been drinking for 30 minutes", icon)
	}

}
