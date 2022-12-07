package model

import (
	"math/big"
	"net"
)

type Response struct {
	IpAddress           string   `json:"ipAddress"`
	Subnet              string   `json:"subnet"`
	FirstIpAddress      net.IP   `json:"firstIpAddress"`
	LastIpAddress       net.IP   `json:"lastIpAddress"`
	Size                *big.Int `json:"size"`
	GeneratedIpAdresses string   `json:"generatedIpAdresses"`
}
