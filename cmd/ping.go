/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"time"

	model "github.com/Juminiy/myping/model"
	myping_utils "github.com/Juminiy/myping/utils"

	"github.com/go-ping/ping"
	"github.com/spf13/cobra"
)

const (
	PING_DEFAULT_URL    = "baidu.com"
	PING_DEFAULT_COUNT  = 1
	PING_DEFAULT_RECORD = true
)

var (
	PingUrl     string = PING_DEFAULT_URL
	PingCount   int    = PING_DEFAULT_COUNT
	PingRecourd bool   = PING_DEFAULT_RECORD
)

// please finishes the args parameters and the URL format
// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "myping pings some domain or some ip address ",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		valid, pingUrl := myping_utils.URLFormatAssign(PingUrl)
		if !valid {
			fmt.Println("Your URL", pingUrl, " is invalid, please check and type again!")
			return
		}
		fmt.Println("🚀 We are going to ping ", pingUrl, " for ", PingCount, " time")
		fmt.Println(" Please wait for a few of time patiently, the time is flying.")
		pinger, err := ping.NewPinger(pingUrl)
		if err != nil {
			panic(err)
		}
		pinger.Count = PingCount

		// start ping
		start_ping := time.Now()
		// IPV6 function
		if err = pinger.Run(); err != nil {
			panic(err)
		} 
		totalTime := time.Since(start_ping).Milliseconds()
		// end ping
		// after ping, we start the spinner,but not show in command line
		//go utils.Spinner(time.Duration(utils.DelayTime), utils.TypeSpinner)
		statis := pinger.Statistics()
		// receive for the json string and stores the json string into redis
		viewStatis := &model.StatisticsRecord{
			PacketRecv: statis.PacketsRecv,
			PacketSent: statis.PacketsSent,
			PacketLoss: statis.PacketLoss,
			IPAddr:     statis.IPAddr,
			MinRtt:     statis.MinRtt,
			MaxRtt:     statis.MaxRtt,
			AvgRtt:     statis.AvgRtt,
			StdDevRtt:  statis.StdDevRtt,
			DestURL:    pingUrl,
			TotalTime:  totalTime,
		}
		model.RecordObserver(PingRecourd, *viewStatis)
		// format output
		viewStatis.FormatOutput(pingUrl)
	},
}
