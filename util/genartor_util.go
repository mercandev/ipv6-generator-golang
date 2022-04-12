package util

import (
	"net"
)

func CheckIPAddress(ip string) {
	if net.ParseIP(ip) == nil {
		panic("Invalid Ip Address!")
	}
}
