package main

import (
	"fmt"
	"log"

	"github.com/miekg/dns"
)

func main() {
	clientConfig, err := dns.ClientConfigFromFile("/etc/resolv.conf")
	if err != nil {
		log.Panic()
	}

	servers := clientConfig.Servers
	fmt.Printf("local dns servers: %+v\n", servers)
}

