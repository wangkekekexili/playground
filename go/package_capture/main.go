package main

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func main() {
	fmt.Printf("libpcap version: %v\n", pcap.Version())

	handle, err := pcap.OpenLive("en0", 65535, false, -1*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	handle.SetBPFFilter("dns")
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		layer := packet.Layer(layers.LayerTypeDNS)
		if layer == nil {
			continue
		}
		dns := layer.(*layers.DNS)
		for _, ans := range dns.Answers {
			if ans.Type != 1 {
				continue
			}
			fmt.Println(string(ans.Name), ans.IP.String())
		}

	}
}
