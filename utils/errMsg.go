package utils

import (
	"database/sql"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func CustomErrMsg(ctx *fiber.Ctx, err error) error {
	// check error is fiber error
	if e, ok := err.(*fiber.Error); ok {
		// Return the Fiber error as JSON
		return ctx.Status(e.Code).JSON(fiber.Map{
			"error": e.Message,
		})
	}

	// For other errors, return a generic JSON error
	return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
		"error": "Unprocessable",
	})
}

func ErrMsg(err error) *fiber.Error {
	if err == sql.ErrNoRows {
		return fiber.NewError(fiber.StatusNotFound, "олдсонгүй")
	}

	return fiber.NewError(fiber.StatusUnprocessableEntity, "алдаа гарлаа")
}

func ValidationErrMsg(err error) error {
	s := ""
	for _, err := range err.(validator.ValidationErrors) {
		if err.Tag() == "required" {
			s = s + fmt.Sprint(err.Field()+" хоосон байна") + ","
		}
	}

	return fiber.NewError(fiber.StatusBadRequest, s)
}
