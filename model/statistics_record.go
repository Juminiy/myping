package model

import (
	"fmt"
	"net"
	"time"

	"github.com/Juminiy/myping/db"

	jsoniter "github.com/json-iterator/go"
)

const (
	REDIS_KEEP_TTL = -1
)

var (
	GlobalJson jsoniter.API
)

func init() {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	time.Local = loc
	GlobalJson = jsoniter.ConfigCompatibleWithStandardLibrary
}

type StatisticsRecord struct {
	PacketRecv int           `json:"packet_recv"`
	PacketSent int           `json:"packet_sent"`
	PacketLoss float64       `json:"packet_loss"`
	DestURL    string        `json:"desk_url"`
	IPAddr     *net.IPAddr   `json:"ip_addr"`
	MinRtt     time.Duration `json:"min_rtt"`
	MaxRtt     time.Duration `json:"max_rtt"`
	AvgRtt     time.Duration `json:"avg_rtt"`
	StdDevRtt  time.Duration `json:"std_dev_rtt"`
	TotalTime  int64         `json:"total_time"`
}

func (statis *StatisticsRecord) FormatOutput(url string) {
	fmt.Println("🎉 Ping work has finished, the statistic data is following : ")
	fmt.Println("🕒️ 🕕️ ", url, "( ", statis.IPAddr, ")", " ping statistics 🕘️ 🕛️ ")
	fmt.Println("Packet Sent: ", statis.PacketSent, "| Packet Received: ", statis.PacketRecv, "| Packet Loss: ", statis.PacketLoss, "| Total Time:", statis.TotalTime, "ms")
	fmt.Println("RTT Min: ", statis.MinRtt, "| RTT Max: ", statis.MaxRtt, "| RTT Avg: ", statis.AvgRtt, "| RTT Mdev: ", statis.StdDevRtt)
}

// struct -> json
// jsonString -> interface
// rdb.set(key,interface)

// Observer Mode
// redis key is time,value is ping data
func RecordObserver(record bool, statistics StatisticsRecord) {
	if record {
		currentTime := time.Now().Local().GoString()
		statJson, err := GlobalJson.MarshalToString(statistics)
		if err != nil {
			panic(err)
		}
		db.RedisClient.Set(db.Ctx, db.PING_RECORD_DEFAULT_PREFIX+currentTime, statJson, REDIS_KEEP_TTL)
	}
}
