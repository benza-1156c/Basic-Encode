package controllers

import (
	"encode/modules/http/usecase"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Req struct {
	Input string `json:"input"`
}

func Encode(c *fiber.Ctx) error {
	var req Req
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	binaryInput := ""
	for _, b := range []byte(req.Input) {
		binaryInput += fmt.Sprintf("%08b ", b)
	}

	encode, key := usecase.Encode(req.Input)
	binary_encode := ""
	for _, b := range []byte(encode) {
		binary_encode += fmt.Sprintf("%08b ", b)
	}

	binary_key := ""
	for _, b := range []byte(key) {
		binary_key += fmt.Sprintf("%08b ", b)
	}

	decode, _ := usecase.Encode(encode)

	return c.JSON(fiber.Map{
		"input":         req.Input,
		"binary-input":  binaryInput,
		"key":           key,
		"binary-key":    binary_key,
		"encode":        encode,
		"binary-encode": binary_encode,
		"decode":        decode,
	})
}
