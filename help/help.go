package help

import (
	"changeme/internal/define"
	"sync"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

// 使用 sync.Map 安全地加载 SessionInfo
func LoadSessionInfo(sessionMap *sync.Map, key string) (define.SessionInfo, bool) {
	value, ok := sessionMap.Load(key)
	if !ok {
		return define.SessionInfo{}, false
	}
	return value.(define.SessionInfo), true
}

// 使用 sync.Map 安全地加载 SessionInfo
func LoadSessionInfoReturnPointer(sessionMap *sync.Map, key string) (*define.SessionInfo, bool) {
	value, ok := sessionMap.Load(key)
	if !ok {
		return nil, false
	}
	return value.(*define.SessionInfo), true
}

// 使用 sync.Map 安全地存储 SessionInfo
func StoreSessionInfo(sessionMap *sync.Map, key string, value define.SessionInfo) {
	sessionMap.Store(key, value)
}

// 辅助函数，根据协议获取协议名称
func GetProtocolName(transportLayer gopacket.TransportLayer) string {
	switch transportLayer.(type) {
	case *layers.TCP:
		return "TCP"
	case *layers.UDP:
		return "UDP"
	default:
		return "Unknown"
	}
}

// 辅助函数，获取源端口
func GetSourcePort(transportLayer gopacket.TransportLayer) int {
	switch transportLayer := transportLayer.(type) {
	case *layers.TCP:
		return int(transportLayer.SrcPort)
	case *layers.UDP:
		return int(transportLayer.SrcPort)
	default:
		return 0
	}
}

// 辅助函数，获取目标端口
func GetDestinationPort(transportLayer gopacket.TransportLayer) int {
	switch transportLayer := transportLayer.(type) {
	case *layers.TCP:
		return int(transportLayer.DstPort)
	case *layers.UDP:
		return int(transportLayer.DstPort)
	default:
		return 0
	}
}

// HandleClientIP 获取客户端的IP地址
func GetClientIP() string {
	return "192.168.50.7"
}
