package clock

import (
	"fmt"
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

func ShowClock(t time.Time) {
	timeArr := createTimeArray(t.Clock())
	timeArr.print()
}
