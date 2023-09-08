package handlers

import (
	"go-backend/db"
	"go-backend/models"
	"go-backend/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slog"
)

func (hd *Handlers) GetTestById(ctx *fiber.Ctx) error {
	queries := db.New(hd.pgsql)
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	testInfo, err := queries.GetTestById(ctx.Context(), int32(id))
	if err != nil {
		slog.Error("unable to get test by id", slog.Any("err", err))
		return utils.ErrMsg(err)
	}

	return ctx.JSON(testInfo)
}

func (hd *Handlers) CreateTest(ctx *fiber.Ctx) error {
	queries := db.New(hd.pgsql)

	body := &models.CreateTest{}

	err := ctx.BodyParser(body)
	if err != nil {
		return utils.ErrMsg(err)
	}

	err = validate.Struct(body)
	if err != nil {
		return utils.ValidationErrMsg(err)
	}

	testInfo, err := queries.CreateTest(ctx.Context(), db.CreateTestParams{
		Name:          body.Name,
		Desc:          body.Desc,
		Img:           body.Img,
		Minute:        body.Minute,
		AgeCls:        body.AgeCls,
		BeforeDesc:    body.BeforeDesc,
		AfterDesc:     body.AfterDesc,
		ExampleReport: body.ExampleReport,
		IsActive:      body.IsActive,
	})
	if err != nil {
		slog.Error("unable to create test", slog.Any("err", err))
		return utils.ErrMsg(err)
	}

	return ctx.JSON(testInfo)
}
