package ping_option

import (
	"fmt"
	"time"

	
)

const (
	PING_RETRY_MIN_TIME     = 5
	PING_RETRY_DEFAULT_TIME = 60000
)

var (
	WorkStart int64 
)

func PingRetry() {

	for {
		if (time.Now().Unix() - WorkStart) > PING_RETRY_DEFAULT_TIME {
			fmt.Println("We are pinging again and again, please wait ... ... ")
		}
	}

}
