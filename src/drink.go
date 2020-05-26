package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"

	"github.com/gen2brain/beeep"
)

//Alert sends desktop notifications based on the os is being used
func Alert(title, message, appIcon string) {
	exOS := runtime.GOOS
	if exOS == "darwin" {
		osa, err := exec.LookPath("osascript")
		if err != nil {
			panic(err)
		}
		cmd := exec.Command(osa, "-e", `tell application "System Events" to display notification "`+message+`" with title "`+title+`" sound name "default"`)
		cmd.Run()
	} else if exOS == "linux" {
		beeep.Notify(title, message, appIcon)
		beeep.Beep(440.0, 200) //Needs special permissions in Linux, doesn't work (for now)
	}
}

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
