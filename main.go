package main

import (
	"ipv6generator/services"

	"github.com/gofiber/fiber/v2"
)

//console

//const GENERATE_IP_ADDRESSES bool = true

//func main() {

//	ipaddress := "2a02:ff0:2f9:b707::"
//	sub := "/124"

//	generator.Informations(ipaddress, sub)

//	util.CheckIPAddress(ipaddress)
//	if GENERATE_IP_ADDRESSES {
//		generator.GenerateIpv6Address(ipaddress, sub)
//	}
//}

//api
func status(c *fiber.Ctx) error {
	return c.SendString("Server is running! Send your request")
}

func setupRoutes(app *fiber.App) {

	app.Get("/", status)

	app.Post("/api/GetIpv6Calculator", services.Ipv6Calculator)
}

func main() {
	app := fiber.New()

	setupRoutes(app)
	app.Listen(":3000")
}
