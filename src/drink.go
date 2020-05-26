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
		beeep.Beep(440.0, 200)
	}
}

func main() {

	time.Sleep(5 * time.Second)

	//Prendo il path dell'eseguibile per trovare il path dell'icona
	exPath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exDir := filepath.Dir(exPath)
	iconPath := "../resources/water-glass.png"
	icon := filepath.Join(exDir, iconPath)

	//Executing first command
	Alert("Drink!!!", "Start drinking now", icon)

	//Eseguo comando ogni 20 minuti
	for {
		time.Sleep(30 * time.Minute)
		Alert("Drink!!!", "You haven't been drinking for 30 minutes", icon)
	}

}
