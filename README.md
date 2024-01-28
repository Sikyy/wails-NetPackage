# Siky-Net


基于Wails构建的请求查看工具 Vue3/Golang

应用界面
![db28dd9f2804445905fea4640ecf2db8.png](https://i.miji.bid/2024/01/28/db28dd9f2804445905fea4640ecf2db8.png)
目前暂时只支持mac
（目前只测试了M2）intel芯片没有条件测试


点击开启请求即会开启捕捉数据包

左边是遍历出的本机网卡，暂时还没有功能，后续打算加上每张网卡走的上下行流量

闪退问题原因：
由于调用了gopacket包，他会需要BPF的权限，sudo是能解决的 但是解决的太暴力，太不优雅了，而且每次走命令行也很丑

闪退问题解决
```
siky at Siky_Macmini in ~
$ cd /dev

siky at Siky_Macmini in /dev
$ who am i
siky             ttys004      Jan 28 18:46

siky at Siky_Macmini in /dev
$ sudo chown siky:admin bp*
Password:

siky at Siky_Macmini in /dev
$ ls -la | grep bp
crw-------   1 siky   admin     0x17000000 Jan 26 20:30 bpf0
crw-------   1 siky   admin     0x17000001 Jan 26 20:30 bpf1
crw-------   1 siky   admin     0x1700000a Jan 23 20:02 bpf10
crw-------   1 siky   admin     0x1700000b Jan 23 20:02 bpf11
crw-------   1 siky   admin     0x1700000c Jan 23 20:02 bpf12
crw-------   1 siky   admin     0x1700000d Jan 23 20:02 bpf13
crw-------   1 siky   admin     0x1700000e Jan 23 20:02 bpf14
crw-------   1 siky   admin     0x1700000f Jan 23 20:02 bpf15
crw-------   1 siky   admin     0x17000010 Jan 23 20:02 bpf16
crw-------   1 siky   admin     0x17000011 Jan 23 20:02 bpf17
crw-------   1 siky   admin     0x17000012 Jan 23 20:02 bpf18
crw-------   1 siky   admin     0x17000013 Jan 23 20:02 bpf19
crw-------   1 siky   admin     0x17000002 Jan 23 21:42 bpf2
crw-------   1 siky   admin     0x17000014 Jan 23 20:02 bpf20
crw-------   1 siky   admin     0x17000015 Jan 23 20:02 bpf21
crw-------   1 siky   admin     0x17000003 Jan 22 23:34 bpf3
crw-------   1 siky   admin     0x17000004 Jan 23 20:02 bpf4
crw-------   1 siky   admin     0x17000005 Jan 23 20:02 bpf5
crw-------   1 siky   admin     0x17000006 Jan 23 22:16 bpf6
crw-------   1 siky   admin     0x17000007 Jan 23 22:16 bpf7
crw-------   1 siky   admin     0x17000008 Jan 23 20:02 bpf8
crw-------   1 siky   admin     0x17000009 Jan 23 20:02 bpf9
```
改成这样就行了

或者直接
```
sudo chmod 777 /dev/bpf*
```


待解决问题：
无法获取客户端名称：我一开始是想着先获取本地的进程pid，然后来拿到进程名字，然后再是端口关联来确认流量包是谁的，然后显示名字，但是感觉要维护数据库，而且这样感觉有点奇怪

无法解析TLS包的host：这个有两种解决办法：
1.写一个代理服务器，让所有的流量先走代理，再出去，那我就应该能把上面的那个问题一起解决了。
2.通过gopacket包来解析第一次拿到的HTTPS包的ClientHello，这应该可以，但是我之前看了半天都知道该怎么解析他的TLS层，先放进待做清单吧

