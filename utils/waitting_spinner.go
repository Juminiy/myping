package utils

import (
	"fmt"
	"time"
)

const (
	CommSpinner        = `-\|/-`
	MazeSpinner        = `🕒️🕕️🕘️🕕️`
	DEFAULT_DELAY_TIME = 100
)

var (
	DelayTime   int64  = DEFAULT_DELAY_TIME
	TypeSpinner string = MazeSpinner
)

func Spinner(delay time.Duration, typeSpinner string) {
	for {
		for _, r := range typeSpinner {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
