package service

import (
	"api-alta-dashboard/features/auth"
	"api-alta-dashboard/middlewares"
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	authData auth.RepositoryInterface
	validate *validator.Validate
}

func New(data auth.RepositoryInterface) auth.ServiceInterface {
	return &authService{
		authData: data,
		validate: validator.New(),
	}
}

func (service *authService) Login(dataCore auth.Core) (auth.Core, string, error) {

	if errValidate := service.validate.Struct(dataCore); errValidate != nil {
		log.Error(errValidate.Error())
		return auth.Core{}, "", errors.New("Failed to Login. Error validate input. Please check your input.")
	}

	result, errLogin := service.authData.FindUser(dataCore.Email)
	if errLogin != nil {
		log.Error(errLogin.Error())
		if strings.Contains(errLogin.Error(), "table") {
			return auth.Core{}, "", errors.New("Failed to Login. Error on request. Please contact your administrator.")
		} else if strings.Contains(errLogin.Error(), "found") {
			return auth.Core{}, "", errors.New("Failed to Login. Email not found. Please check password again.")
		} else {
			return auth.Core{}, "", errors.New("Failed to Login. Other Error. Please contact your administrator.")
		}
	}

	errCheckPass := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(dataCore.Password))
	fmt.Println("Data Core = ", dataCore, "\n\n\n")
	fmt.Println("Result = ", result, "\n\n\n")
	if errCheckPass != nil {
		log.Error(errCheckPass.Error())
		return auth.Core{}, "", errors.New("Failed to Login. Password didn't match. Please check password again.")
	}

	token, errToken := middlewares.CreateToken(int(result.ID), result.Role)
	if errToken != nil {
		log.Error(errToken.Error())
		return auth.Core{}, "", errors.New("Failed to login. Error on generate token. Please check password again.")
	}

	return result, token, nil
}
