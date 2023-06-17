## ncat
check tcp/icmp

```shell
ncat -h

Usage:

    ncat [-p tcp/icmp] [host,port]

Examples:

    # ping google continuously
    ncat www.google.com

    # tcp check 
    ncat -p tcp 8.8.8.8 22
    ncat -p tcp github.com 443 80 22 5555
```