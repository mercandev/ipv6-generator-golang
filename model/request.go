package model

type Request struct {
	IpAddress         string `json:"ipAddress"`
	Subnet            string `json:"subnet"`
	GenerateIpAddress bool   `json:"generateIpAddress"`
}
