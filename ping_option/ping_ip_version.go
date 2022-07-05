package ping_option

const (
	PING_DEFAULT_IP_VERSION = PING_IP_VERSION_4
	PING_CMD_IP_VERSION_4 = "4"
	PING_CMD_IP_VERSION_6 = "6" 
	PING_IP_VERSION_4 = "ip4"
	PING_IP_VERSION_6 = "ip6"
)

var (
	PingIpVersion string = PING_CMD_IP_VERSION_4 
)
