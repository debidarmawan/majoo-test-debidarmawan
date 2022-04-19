package repositories

import (
	"fmt"
	"majoo-test-debidarmawan/config"
	"majoo-test-debidarmawan/models"
	"net/http"
)

type UserRepo interface {
	Login(request models.Login) chan models.Result
}

type userRepo struct {
	dbConn *config.DbConnection
}

func NewUserRepo(dbConn *config.DbConnection) UserRepo {
	return &userRepo{
		dbConn: dbConn,
	}
}

func (ur *userRepo) Login(request models.Login) chan models.Result {
	output := make(chan models.Result)
	go func() {
		var user models.User
		err := ur.dbConn.MajooDB.Where("user_name = ?", request.UserName).First(&user).Error
		if err != nil {
			fmt.Println(err)
			output <- models.Result{StatusCode: http.StatusInternalServerError, Error: err, Data: nil, Message: "Internal Server Error"}
			return
		}
		output <- models.Result{StatusCode: http.StatusOK, Data: user, Error: nil, Message: "Success"}
	}()
	return output
}
