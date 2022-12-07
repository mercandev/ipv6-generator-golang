package services

import (
	"ipv6generator/generator"
	"ipv6generator/model"

	"github.com/gofiber/fiber/v2"
)

func Ipv6Calculator(c *fiber.Ctx) error {

	request := new(model.Request)
	err := c.BodyParser(request)

	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
		return err
	}

	result := generator.IpAddressInformations(request.IpAddress, request.Subnet)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
		return err
	}

	c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "",
		"data":    result,
	})
	return nil
}
