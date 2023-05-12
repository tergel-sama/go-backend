package handlers

import (
	"mind-demo-backend/db"
	"mind-demo-backend/models"
	"mind-demo-backend/utils"

	"github.com/gofiber/fiber/v2"
)

// @Summary get test by id
// @Tags Test
// @Param id path int true "Item ID"
// @Router /api/test/{id} [get]
func (hd *Handlers) GetTestById(ctx *fiber.Ctx) error {
	queries := db.New(hd.pgsql)
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	testInfo, err := queries.GetTestById(ctx.Context(), int32(id))
	if err != nil {
		hd.lerr.Printf("unable to get test by id: %v", err)
		return utils.ErrMsg(err)
	}

	return ctx.JSON(testInfo)
}

// @Summary create test
// @Tags Test
// @Param requestBody body models.CreateTest true "Test information"
// @Router /api/test/ [post]
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
		hd.lerr.Printf("unable to create test: %v", err)
		return utils.ErrMsg(err)
	}

	return ctx.JSON(testInfo)
}
