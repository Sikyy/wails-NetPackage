package define

import (
	"sync"
	"time"

	"github.com/google/gopacket/pcap"
)

//定义常量

var SessionMap sync.Map

type Network struct {
	Name        string
	Description string
	Addresses   []pcap.InterfaceAddress
}

type SessionInfo struct {
	ID                 int64         // 会话ID
	Name               string        // 会话名称
	Status             string        // 会话状态
	Bytes              float64       // 数据包的长度
	Time               time.Time     // 数据包的时间
	Time_s             string        // 数据包的时间
	SrcMAC             string        // 源MAC地址
	SrcIP              string        // 源IP地址
	SrcPort            int           // 源端口号
	DstMAC             string        // 目的MAC地址
	DstIP              string        // 目的IP地址
	DstPort            int           // 目的端口号
	Protocol           string        // 协议号
	TCPStatus          string        // TCP状态
	UDPStatus          string        // UDP状态
	SessionUpTraffic   float64       // 会话上行流量信息
	SessionDownTraffic float64       // 会话下行流量信息
	StartTime          time.Time     //会话开始时间
	EndTime            time.Time     //会话结束时间
	Length_of_time     time.Duration //会话持续时间
	Host               string        //会话的主机
	Method             string        //会话的请求方法
}

// 前端显示信息
type SessionInfoFront struct {
	ID                 int64  // 会话ID
	ClientName         string // 会话名称
	Time_s             string // 数据包的时间
	Status             string // 会话状态
	SessionUpTraffic   string // 会话上行流量信息
	SessionDownTraffic string // 会话下行流量信息
	Length_of_time     int64  //会话持续时间
	Method             string //会话的请求方法
	Host               string //会话的主机
}
