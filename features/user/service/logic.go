package service

import (
	"api-alta-dashboard/features/user"
	"api-alta-dashboard/utils/helper"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepository user.RepositoryInterface
	validate       *validator.Validate
}

func New(repo user.RepositoryInterface) user.ServiceInterface {
	return &userService{
		userRepository: repo,
		validate:       validator.New(),
	}
}

// Create implements user.ServiceInterface
func (service *userService) Create(input user.Core) (err error) {
	//validate
	// if input.Name == "" || input.Email == "" || input.Password == "" {
	// 	return errors.New("Name, email, password harus diisi")
	// }

	// Default Role
	input.Role = "User"

	// validasi input
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	// validasi email harus unik
	_, rowAffected, errFindEmail := service.userRepository.FindUser(input.Email)

	if rowAffected >= 0 {
		// log.Error(errFindEmail.Error())
		return errors.New("Failed. Email " + input.Email + " already exist. Please pick another email.")
	}
	if errFindEmail != nil {
		return helper.ServiceErrorMsg(errFindEmail)
	}

	bytePass, errEncrypt := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if errEncrypt != nil {
		log.Error(errEncrypt.Error())
		return helper.ServiceErrorMsg(err)
	}

	input.Password = string(bytePass)

	_, errCreate := service.userRepository.Create(input)
	if errCreate != nil {
		log.Error(errCreate.Error())
		return helper.ServiceErrorMsg(err)
	}

	return nil
}

// GetAll implements user.ServiceInterface
func (service *userService) GetAll() (data []user.Core, err error) {
	data, err = service.userRepository.GetAll()
	return data, err
}

func (service *userService) GetById(id int) (data user.Core, err error) {
	data, err = service.userRepository.GetById(id)
	if err != nil {
		log.Error(err.Error())
		return user.Core{}, helper.ServiceErrorMsg(err)
	}
	return data, err

}

func (service *userService) Update(input user.Core, id int) error {
	if input.Password != "" {
		generate, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
		input.Password = string(generate)
	}

	_, err := service.userRepository.Update(input, id)
	if err != nil {
		log.Error(err.Error())
		return helper.ServiceErrorMsg(err)
	}

	return nil
}

func (service *userService) Delete(id int) error {
	_, err := service.userRepository.Delete(id)
	if err != nil {
		log.Error(err.Error())
		return helper.ServiceErrorMsg(err)
	}
	return nil
}
