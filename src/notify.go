package main

import (
	"os/exec"
	"runtime"

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
