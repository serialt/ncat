package main

import (
	"flag"
	"fmt"

	"github.com/serialt/sugar/v3"
	"golang.org/x/exp/slog"
)

var usage = `
Usage:

    ping [-c count] [-i interval] [-t timeout] [--privileged] host

Examples:

    # ping google continuously
    ping www.google.com

    # ping google 5 times
    ping -c 5 www.google.com

    # ping google 5 times at 500ms intervals
    ping -c 5 -i 500ms www.google.com

    # ping google for 10 seconds
    ping -t 10s www.google.com

    # Send a privileged raw ICMP ping
    sudo ping --privileged www.google.com

    # Send ICMP messages with a 100-byte payload
    ping -s 100 1.1.1.1
`

func init() {
	flag.BoolVar(&appVersion, "v", false, "Display build and version messages")
	flag.StringVar(&protocol, "p", "icmp", "[tcp,udp,icmp]")
	flag.BoolVar(&privileged, "privileged", false, "")
	flag.Usage = func() {
		fmt.Print(usage)
	}
	flag.Parse()

	slog.SetDefault(sugar.New())

}

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

	host := flag.Arg(0)
	Ping(host)
	// service()
}
