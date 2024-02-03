# Siky-Net


基于Wails构建的请求查看工具 Vue3/Golang

应用界面
![1](https://siky-telegraph-image.pages.dev/file/862d0bd5b94807ae9b7eb.png)
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

2024.02.04 更新
解决了HTTPS远端IP的显示
解决了进程名字的显示

因为UDP包拿不到进程信息我就把UDP砍了
感觉除了HTTPS和HTTP的TCP包不是很重要，而且有些会获取不到客户端信息（也有可能是我还在获取进程然后对面端口就关了的原因...）简而言之太快了
所以我把这个也砍了
还有一个是获取不到远端IP的HTTPS数据包，一般来说要么是关太快，要么是。。感觉没其他可能了，要是长时间监听是一定有的
也砍了

待解决问题：
wails dev环境下一切正常
但是在运行build打包好的应用后中文字符会出现\xe5\x93\x94\xe5\x93\xa9\xe5\x93\x94\xe5\x93\xa9 的情况
本知道是不是我本地环境的问题，因为我的终端也会这个样子，但是按道理来说我函数都没有给终端经手，我直接把它原数据拿过来了...再看看吧，我也找不到人测试

不正常情况：
![2](https://siky-telegraph-image.pages.dev/file/15b44071f7b579765092e.png)

