package generator

import (
	"fmt"
	"ipv6generator/model"
	"net"

	"gopkg.in/netaddr.v1"
)

func Increment(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func GenerateIpv6Address(ipaddress string, sub string) {

	ip, ipnet, err := net.ParseCIDR(ipaddress + sub)

	if ip.IsPrivate() {
		panic("Ip Address is private!") //Only public intr.
	}

	if err != nil {
		panic(err)
	}

	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); Increment(ip) {
		fmt.Println(ip)
	}
}

func Informations(ipaddress string, sub string) {

	ip, ipnet, err := net.ParseCIDR(ipaddress + sub)
	if err != nil {
		panic(err)
	}

	fmt.Printf("IP Address: %s \n", ipaddress+sub)
	fmt.Printf("First Ip Address: %s \n", ip)
	fmt.Printf("Last Ip Address: %s \n", netaddr.BroadcastAddr(ipnet))
	fmt.Printf("Size: %s \n", (netaddr.NetSize(ipnet)))
}

func IpAddressInformations(ipaddress string, sub string) model.Response {

	ip, ipnet, err := net.ParseCIDR(ipaddress + sub)
	if err != nil {
		panic(err)
	}

	response := model.Response{
		IpAddress:      ipaddress,
		Subnet:         sub,
		FirstIpAddress: ip,
		LastIpAddress:  netaddr.BroadcastAddr(ipnet),
		Size:           netaddr.NetSize(ipnet),
	}

	return response
}
