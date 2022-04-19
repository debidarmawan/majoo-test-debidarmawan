package usecases

import (
	"encoding/json"
	"errors"
	"majoo-test-debidarmawan/libs"
	"majoo-test-debidarmawan/libs/helpers"
	"majoo-test-debidarmawan/models"
	"majoo-test-debidarmawan/repositories"

	"github.com/gofiber/fiber/v2"
)

type UserUseCase interface {
	Login(request models.Login) chan models.Result
}

type userUsecase struct {
	userRepo repositories.UserRepo
}

func NewUserUseCase(userRepo repositories.UserRepo) UserUseCase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (uu *userUsecase) Login(request models.Login) chan models.Result {
	output := make(chan models.Result)
	go func() {
		result := <-uu.userRepo.Login(request)
		if result.StatusCode == fiber.StatusOK {
			var userModel models.User
			dataByte, _ := json.Marshal(result.Data)
			_ = json.Unmarshal(dataByte, &userModel)
			md5Password := helpers.GetMD5Hash(request.Password)
			if md5Password != userModel.Password {
				output <- models.Result{
					Data:       nil,
					StatusCode: fiber.StatusUnauthorized,
					Error:      errors.New("unauthorized"),
					Message:    "Login Failed!",
				}
				return
			}
			token := libs.GenerateToken(userModel.ID)
			resp := map[string]interface{}{
				"user_data": userModel,
				"token":     token,
			}
			result.Data = resp
		}
		output <- result
	}()
	return output
}
