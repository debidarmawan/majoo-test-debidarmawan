package handlers

import (
	"majoo-test-debidarmawan/models"
	"majoo-test-debidarmawan/usecases"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userUseCase usecases.UserUseCase
}

func NewUserHandler(userUseCase usecases.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

func (uh *UserHandler) Routes(group fiber.Router) {
	group.Post("/login", uh.Login)
}

// @Title Majoo Assessment User Login
// @Summary Majoo Assessment User Login
// @Tags Users
// @Description Majoo Assessment User Login
// @param	body	body	models.Login	true	"User Login"
// @Success 200 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /login [post]
func (uh *UserHandler) Login(c *fiber.Ctx) error {
	var loginModel models.Login
	if err := c.BodyParser(&loginModel); err != nil {
		result := models.Result{
			StatusCode: fiber.StatusBadRequest,
			Error:      err,
			Data:       nil,
			Message:    "please make sure your payload data",
		}
		return c.Status(result.StatusCode).JSON(result.ToResponseError(result.StatusCode))
	}
	result := <-uh.userUseCase.Login(loginModel)
	if result.Error != nil {
		return c.Status(result.StatusCode).JSON(result.ToResponseError(result.StatusCode))
	}
	return c.Status(result.StatusCode).JSON(result.ToResponse())
}
