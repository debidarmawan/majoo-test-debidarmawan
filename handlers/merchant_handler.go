package handlers

import (
	"majoo-test-debidarmawan/libs"
	"majoo-test-debidarmawan/middleware"
	"majoo-test-debidarmawan/models"
	"majoo-test-debidarmawan/usecases"
	"time"

	"github.com/gofiber/fiber/v2"
)

type MerchantHandler struct {
	merchantUseCase usecases.MerchantUseCase
}

func NewMerchantHandler(merchantUseCase usecases.MerchantUseCase) *MerchantHandler {
	return &MerchantHandler{
		merchantUseCase: merchantUseCase,
	}
}

func (mh *MerchantHandler) Routes(group fiber.Router) {
	group.Get("/merchant/omzet", mh.MerchantOmzet, middleware.JWTProtected())
	group.Get("/merchant/outlet/omzet", mh.MerchantsOutletOmzet, middleware.JWTProtected())
}

// @Title Majoo Assessment Get Merchant Omzet
// @Summary Majoo Assessment Get Merchant Omzet
// @Tags Merchants
// @Description Majoo Assessment Get Merchant Omzet
// @Param Authorization	header	string	true	"Authorization"
// @Param period		query	string	false	"Period"
// @Param limit			query	int		false	"Limit"
// @Param page			query	int		false	"Page"
// @Success 200 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /merchant/omzet [get]
func (mh *MerchantHandler) MerchantOmzet(c *fiber.Ctx) error {
	now := time.Now().Unix()
	claims, err := libs.AuthValidator(c)
	if err != nil {
		result := models.Result{
			StatusCode: fiber.StatusInternalServerError,
			Error:      err,
			Data:       nil,
			Message:    err.Error(),
		}
		return c.Status(result.StatusCode).JSON(result.ToResponseError(result.StatusCode))
	}
	expires := claims.Expires
	if now > expires {
		if err != nil {
			result := models.Result{
				StatusCode: fiber.StatusUnauthorized,
				Error:      err,
				Data:       nil,
				Message:    "unauthorized, token is expired",
			}
			return c.Status(result.StatusCode).JSON(result.ToResponseError(result.StatusCode))
		}
	}

	var merchantOmzet models.MerchantOmzet
	merchantOmzet.UserID = claims.UserID

	if err := c.QueryParser(&merchantOmzet); err != nil {
		result := models.Result{
			StatusCode: fiber.StatusBadRequest,
			Error:      err,
			Data:       nil,
			Message:    "please make sure your payload data",
		}
		return c.Status(result.StatusCode).JSON(result.ToResponseError(result.StatusCode))
	}
	result := <-mh.merchantUseCase.MerchantOmzet(merchantOmzet)
	if result.Error != nil {
		return c.Status(result.StatusCode).JSON(result.ToResponseError(result.StatusCode))
	}
	return c.Status(result.StatusCode).JSON(result.ToResponse())
}

// @Title Majoo Assessment Get Merchant Outlet Omzet
// @Summary Majoo Assessment Get Merchant Outlet Omzet
// @Tags Merchants
// @Description Majoo Assessment Get Merchant Outlet Omzet
// @Param Authorization	header	string	true	"Authorization"
// @Param outlet_id		query	uint64	false	"Outlet ID"
// @Param period		query	string	false	"Period"
// @Param limit			query	int		false	"Limit"
// @Param page			query	int		false	"Page"
// @Success 200 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /merchant/outlet/omzet [get]
func (mh *MerchantHandler) MerchantsOutletOmzet(c *fiber.Ctx) error {
	now := time.Now().Unix()
	claims, err := libs.AuthValidator(c)
	if err != nil {
		result := models.Result{
			StatusCode: fiber.StatusInternalServerError,
			Error:      err,
			Data:       nil,
			Message:    err.Error(),
		}
		return c.Status(result.StatusCode).JSON(result.ToResponseError(result.StatusCode))
	}
	expires := claims.Expires
	if now > expires {
		if err != nil {
			result := models.Result{
				StatusCode: fiber.StatusUnauthorized,
				Error:      err,
				Data:       nil,
				Message:    "unauthorized, token is expired",
			}
			return c.Status(result.StatusCode).JSON(result.ToResponseError(result.StatusCode))
		}
	}

	var MerchantOutletOmzet models.MerchantOutletOmzet
	MerchantOutletOmzet.UserID = claims.UserID

	if err := c.QueryParser(&MerchantOutletOmzet); err != nil {
		result := models.Result{
			StatusCode: fiber.StatusBadRequest,
			Error:      err,
			Data:       nil,
			Message:    "please make sure your payload data",
		}
		return c.Status(result.StatusCode).JSON(result.ToResponseError(result.StatusCode))
	}
	result := <-mh.merchantUseCase.MerchantsOutletOmzet(MerchantOutletOmzet)
	if result.Error != nil {
		return c.Status(result.StatusCode).JSON(result.ToResponseError(result.StatusCode))
	}
	return c.Status(result.StatusCode).JSON(result.ToResponse())
}
