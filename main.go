package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"time"

	probing "github.com/prometheus-community/pro-bing"
)

var usage = `
Usage:

    ncat [-p tcp/icmp] [host,port]

Examples:

    # ping google continuously
    ncat www.google.com

    # tcp check 
    ncat -p tcp github.com 22
    ncat -p tcp github.com 443 80 22 5555

`

func init() {
	flag.BoolVar(&appVersion, "v", false, "Display build and version messages")
	flag.StringVar(&protocol, "p", "icmp", "[tcp,udp,icmp], default icmp")
	flag.IntVar(&timeout, "T", 1000, "")
	flag.BoolVar(&privileged, "privileged", false, "")

	flag.Usage = func() {
		fmt.Print(usage)
	}
	flag.Parse()

}

var (
	// 版本信息
	appVersion bool // 控制是否显示版本
	APPVersion = "v0.0.2"
	BuildTime  = "2006-01-02 15:04:05"
	GitCommit  = "xxxxxxxxxxx"

	protocol   string
	timeout    int
	privileged bool
	wg         sync.WaitGroup
)

func main() {
	if appVersion {
		fmt.Printf("APPVersion: %v  BuildTime: %v  GitCommit: %v\n",
			APPVersion,
			BuildTime,
			GitCommit)
		return
	}
	if flag.NArg() == 0 {
		flag.Usage()
		return
	}

	switch protocol {
	case "icmp":
		host := flag.Arg(0)
		Ping(host)
	case "tcp", "udp":
		host := flag.Args()[0]
		port := flag.Args()[1:]
		var checkList []string
		for _, v := range port {
			member := fmt.Sprintf("%s:%s", host, v)
			checkList = append(checkList, member)
		}
		for _, m := range checkList {
			go func(_m string) {
				wg.Add(1)
				TCPOrUDPAlive(protocol, _m)
			}(m)
		}
		fmt.Println(checkList)

	}

	wg.Wait()
}

func Ping(host string) {
	// count := -1
	// size := 24
	// ttl := 64
	pinger, err := probing.NewPinger(host)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	// listen for ctrl-C signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			pinger.Stop()
		}
	}()

	pinger.OnRecv = func(pkt *probing.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d ttl=%v time=%v\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.TTL, pkt.Rtt)
	}
	pinger.OnDuplicateRecv = func(pkt *probing.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d ttl=%v time=%v (DUP!)\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.TTL, pkt.Rtt)
	}
	pinger.OnFinish = func(stats *probing.Statistics) {
		fmt.Printf("\n--- %s ping statistics ---\n", stats.Addr)
		fmt.Printf("%d packets transmitted, %d packets received, %d duplicates, %v%% packet loss\n",
			stats.PacketsSent, stats.PacketsRecv, stats.PacketsRecvDuplicates, stats.PacketLoss)
		fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
			stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
	}

	pinger.Count = -1
	pinger.Size = 56

	pinger.SetPrivileged(privileged)

	fmt.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
	err = pinger.Run()
	if err != nil {
		fmt.Println("Failed to ping target host:", err)
	}
}

func TCPOrUDPAlive(protocol string, host string) {
	tcpTimeOut := time.Duration(3000)
	conn, err := net.DialTimeout(protocol, host, time.Duration(tcpTimeOut*time.Millisecond))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		wg.Done()
		return
	} else {
		fmt.Printf("Open %s: %s\n", protocol, host)
	}
	conn.Close()
	wg.Done()
}
