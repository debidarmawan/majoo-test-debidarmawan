package middleware

import (
	"majoo-test-debidarmawan/models"
	"os"

	"github.com/gofiber/fiber/v2"
	jwtMiddleware "github.com/gofiber/jwt/v2"
)

func JWTProtected() func(*fiber.Ctx) error {
	config := jwtMiddleware.Config{
		SigningKey:   []byte(os.Getenv("SECRET_KEY")),
		ContextKey:   "jwt",
		ErrorHandler: jwtError,
	}
	return jwtMiddleware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		result := models.Result{
			StatusCode: fiber.StatusBadRequest,
			Error:      err,
			Data:       nil,
			Message:    err.Error(),
		}
		return c.Status(result.StatusCode).JSON(result.ToResponseError(result.StatusCode))
	}
	result := models.Result{
		StatusCode: fiber.StatusUnauthorized,
		Error:      err,
		Data:       nil,
		Message:    err.Error(),
	}
	return c.Status(result.StatusCode).JSON(result.ToResponseError(result.StatusCode))
}
