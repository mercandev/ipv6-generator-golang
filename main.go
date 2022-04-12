package main

import (
	"ipv6generator/generator"
	"ipv6generator/util"
)

const GENERATE_IP_ADDRESSES bool = false

func main() {

	ipaddress := "2a02:ff0:2f9:b707::"
	sub := "/48"

	generator.Informations(ipaddress, sub)

	util.CheckIPAddress(ipaddress)
	if GENERATE_IP_ADDRESSES {
		generator.GenerateIpv6Address(ipaddress, sub)
	}
}
