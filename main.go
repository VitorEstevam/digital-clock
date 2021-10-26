package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

type TimeArray []digitalNumber

func (t TimeArray) print() {
	for i := 0; i < 5; i++ {
		for _, j := range t {
			fmt.Print(" " + j[i])
		}
		fmt.Println()
	}
}

func createTimeArray(hours, minutes, seconds int) TimeArray {
	var separator string
	if seconds%2 == 0 {
		separator = ":"
	} else {
		separator = " "
	}

	var n = TimeArray{
		numbers[hours/10],
		numbers[hours%10],
		simbols[separator],
		numbers[minutes/10],
		numbers[minutes%10],
		simbols[separator],
		numbers[seconds/10],
		numbers[seconds%10],
	}

	return n
}

func printTime(t time.Time) {
	timeArr := createTimeArray(t.Clock())
	timeArr.print()
}

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	for {
		clearConsole()
		dt := time.Now()
		printTime(dt)
		time.Sleep(1 * time.Second)
	}
}
