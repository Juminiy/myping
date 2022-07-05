/*
Copyright © 2022 myping <hln0x29a@gmail.com>

*/
package main

import (
	"time"

	"github.com/Juminiy/myping/cmd"
	"github.com/Juminiy/myping/utils"
)

func main() {

	// // Retry
	// go ping_option.PingRetry()
	// Spinner
	go utils.Spinner(time.Duration(utils.DelayTime), utils.TypeSpinner)
	// CLI
	cmd.Execute()

}
