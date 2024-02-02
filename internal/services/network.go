package services

import (
	"changeme/internal/define"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var ID int64 = 0
var NetHandle *pcap.Handle
var StopCaptureCh chan struct{}
var StopSessionInfoCh chan struct{}

// Network list 枚举网卡
func Networklist() ([]string, error) {
	// 得到所有的(网络)设备
	interfaces, err := pcap.FindAllDevs()
	if err != nil {
		return nil, err
	}
	// 定义切片来存储接口名称
	var ifaceNames []string

	for _, iface := range interfaces {
		// fmt.Printf("Capturing traffic on interface %s...\n", iface.Name)
		// 将接口名称添加到切片中
		ifaceNames = append(ifaceNames, iface.Name)
	}
	return ifaceNames, nil
}

// 进行数据包捕获和处理
func CaptureTraffic(SessionInfoCh chan<- define.SessionInfoFront) {
	// 枚举网卡
	ifaces, err := Networklist()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, iface := range ifaces {
		// 定义一个管道，用于停止数据包捕获
		StopCaptureCh = make(chan struct{})
		// 定义一个互斥锁
		var mu sync.Mutex
		// 打开网络接口，返回一个网络接口的句柄
		handle, err := pcap.OpenLive(
			iface,             // 网络接口名
			65536,             // 每个数据包的最大长度
			true,              // 设置为 true，表示开启混杂模式
			pcap.BlockForever, //表示永远阻塞，不会超时返回。而使用负数时间，可以认为是一个“非常大的超时时间”，基本上达到了永远阻塞的效果
		)
		if err != nil {
			fmt.Printf("Error opening interface %s: %v\n", iface, err)
			log.Fatal(err)
		}
		// 保证程序退出时关闭网络接口
		defer handle.Close()
		// 保证程序退出时关闭管道
		defer close(StopCaptureCh)

		// 存储句柄到全局变量
		NetHandle = handle

		// 创建一个数据包源，用于接收数据包。
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

		mu.Lock()
		defer mu.Unlock()

		for {
			select {
			case packet := <-packetSource.Packets():
				// 处理数据包
				// 把信息存入map中，并生成ID，计算会话开始时间，实现会话初始化
				SessionInfo := JudgeIDAndWriteByteSessionMap(packet, &ID, &define.SessionMap)
				fmt.Println(packet)
				// 按照Surge请求查看器格式输出
				ProcessPacket(packet, &SessionInfo)
				// 实时发送会话信息到通道
				info := ReturnPacket(packet, &SessionInfo)
				select {
				case SessionInfoCh <- info:
					//接收关闭信号，关闭通道
				// case <-StopSessionInfoCh:
				// 	// 重制全局变量
				// 	SessionInfoCh = make(chan<- define.SessionInfoFront)
				// 	return
				default:
					fmt.Println("SessionInfoCh is full")
				}

			case <-StopCaptureCh:
				// 接收到停止捕获信号，退出循环
				// 重制全局变量
				StopCaptureCh = make(chan struct{})
				return
			}
		}
	}
}

// 按照Surge请求查看器格式输出
func ProcessPacket(packet gopacket.Packet, sessionInfo *define.SessionInfo) {
	// 在这里添加解析和整理数据包的逻辑
	if sessionInfo.ID == 0 {
		fmt.Println("Error: SessionInfo is empty")
		return
	}

	// 以下是一个示例，你需要根据实际情况进行修改
	fmt.Printf("会话ID: %d\n", sessionInfo.ID)
	fmt.Printf("会话名称: %s\n", sessionInfo.Name)
	fmt.Printf("日期: %s\n", time.Now().Format("15:04:05"))
	fmt.Printf("客户端 状态: %s\n", sessionInfo.TCPStatus)
	fmt.Printf("策略: %s\n", "Normal") // 这里需要替换为实际的策略
	fmt.Printf("上传: %s 下载: %s\n", FormatBytes(sessionInfo.SessionUpTraffic), FormatBytes(sessionInfo.SessionDownTraffic))
	fmt.Printf("时长: %v ms, 方法: %s\n", sessionInfo.EndTime.Sub(sessionInfo.StartTime), sessionInfo.Method)
	fmt.Printf("开始时间: %v, 结束时间: %v\n", sessionInfo.StartTime, sessionInfo.EndTime)
	fmt.Printf("地址: %s\n", sessionInfo.Host)
	fmt.Println("----")
}

// 返回数据包格式给前端进行展示
func ReturnPacket(packet gopacket.Packet, sessionInfo *define.SessionInfo) define.SessionInfoFront {
	// 在这里添加解析和整理数据包的逻辑
	return define.SessionInfoFront{
		ID:                 sessionInfo.ID,                                            //会话ID
		ClientName:         sessionInfo.Name,                                          // 会话名称
		Time_s:             time.Now().Format("15:04:05"),                             // 数据包的时间
		Status:             sessionInfo.TCPStatus,                                     // 会话状态
		SessionUpTraffic:   FormatBytes(sessionInfo.SessionUpTraffic),                 // 会话上行流量信息
		SessionDownTraffic: FormatBytes(sessionInfo.SessionDownTraffic),               // 会话下行流量信息
		Length_of_time:     GetMSTime(sessionInfo.EndTime.Sub(sessionInfo.StartTime)), //会话持续时间
		Method:             sessionInfo.Method,                                        //会话的请求方法
		Host:               sessionInfo.Host,                                          //会话的主机
	}
}

// 停止数据包捕获
func StopCapture() {
	// 互斥锁，确保在停止捕获的同时可以安全地修改全局变量
	mu.Lock()
	defer mu.Unlock()

	if NetHandle != nil {
		// 发送停止捕获信号
		StopCaptureCh <- struct{}{}
		// StopSessionInfoCh <- struct{}{}

	}
}

func GetMSTime(duration time.Duration) int64 {
	return duration.Milliseconds()
}
