package signal

import (
	"fmt"
	"testing"

	"github.com/go-ping/ping"
)

func ICMPTest(t *testing.T) {
	var domain string
	var count int
	fmt.Println("Please input Domain:")
	//fmt.Scan(&domain)
	domain = "baidu.com"
	fmt.Println("Please input Count:")
	//fmt.Scan(&count)
	count = 1
	pinger, err := ping.NewPinger(domain)
	if err != nil {
		panic(err)
	}
	pinger.Count = count
	if err = pinger.Run(); err != nil {
		panic(err)
	}
	stats := pinger.Statistics()
	fmt.Println(stats)
}
