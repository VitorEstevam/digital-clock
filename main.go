package main

import (
	"os"
	"os/exec"
	"time"

	"github.com/VitorEstevam/digital-clock/clock"
)

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	for {
		clearConsole()
		dt := time.Now()
		clock.ShowClock(dt)
		time.Sleep(1 * time.Second)
	}
}
