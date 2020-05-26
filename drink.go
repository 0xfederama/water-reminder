package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"
)

func main() {

	time.Sleep(15 * time.Second)

	//Controllo il tipo di os per mandare la notifica. A seconda dell'os eseguo comandi diversi
	exOS := runtime.GOOS
	var notify string = ""
	var param string = ""
	if exOS == "linux" {
		notify = "notify-send"
		param = "-i"
	} else {
		notify = "osascript"
		param = "-e"
	}

	//Prendo il path dell'eseguibile per trovare il path dell'icona
	exPath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exDir := filepath.Dir(exPath)
	iconPath := "resources/water-glass.png"
	icon := filepath.Join(exDir, iconPath)

	//Eseguo il primo comando
	if exOS == "linux" {
		cmd := exec.Command(notify, param, icon, "Drink!!!", "Start drinking now")
		cmd.Run()
	} else {
		cmd := exec.Command(notify, param, "'display notification \"Start drinking now\" with title \"Drink\"'")
		cmd.Run()
	}

	//Eseguo comando ogni 20 minuti
	for {
		time.Sleep(20 * time.Minute)
		if exOS == "linux" {
			cmd := exec.Command(notify, param, icon, "Drink!!!", "You haven't been drinking for 20 minutes")
			cmd.Run()
		} else {
			cmd := exec.Command(notify, param, "'display notification \"You haven't been drinking for 20 minutes\" with title \"Drink\"'")
			cmd.Run()
		}
	}

}
