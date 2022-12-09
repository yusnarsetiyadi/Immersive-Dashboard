package service

import (
	"api-alta-dashboard/features/mentee"
	"api-alta-dashboard/utils/helper"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
)

type menteeService struct {
	menteeRepository mentee.RepositoryInterface
	validate         *validator.Validate
}

func New(repo mentee.RepositoryInterface) mentee.ServiceInterface {
	return &menteeService{
		menteeRepository: repo,
		validate:         validator.New(),
	}
}

// Create implements user.ServiceInterface
func (service *menteeService) Create(input mentee.Core) (err error) {

	// validasi input
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	// validasi email harus unik
	_, errFindEmail := service.menteeRepository.FindUser(input.Email)

	// helper.LogDebug("\n\n\n find email res  ", res)
	// helper.LogDebug("\n\n\n find email rowAffected  ", rowAffected)

	// if rowAffected > 0 {
	// 	return errors.New("Failed. Email " + input.Email + " already exist. Please pick another email.")
	// }

	if errFindEmail != nil && !strings.Contains(errFindEmail.Error(), "found") {
		return helper.ServiceErrorMsg(errFindEmail)
	}

	errCreate := service.menteeRepository.Create(input)
	if errCreate != nil {
		log.Error(errCreate.Error())
		return helper.ServiceErrorMsg(err)
	}

	return nil
}

// GetAll implements user.ServiceInterface
func (service *menteeService) GetAll(queryName, queryStatus, queryIdClass, queryEdType string) (data []mentee.Core, err error) {

	// if queryName == "" && queryStatus == "" && queryIdClass == "" && queryEdType == "" {
	// 	data, err = service.menteeRepository.GetAll()
	// }
	// if queryName == "not nil" || queryStatus == "not nil" || queryIdClass == "not nil" || queryEdType == "not nil" {
	// 	data, err = service.menteeRepository.GetAllWithSearch(queryName, queryStatus, queryIdClass, queryEdType)
	// }
	data, err = service.menteeRepository.GetAllWithSearch(queryName, queryStatus, queryIdClass, queryEdType)

	// if len(queryName) == 0 {
	// 	data, err = service.menteeRepository.GetAll(queryStatus, queryIdClass, queryEdType)
	// } else {
	// }

	if err != nil {
		helper.LogDebug(err)
		return nil, helper.ServiceErrorMsg(err)
	}

	if len(data) == 0 {
		helper.LogDebug("Get data success. No data.")
		return nil, errors.New("Get data success. No data.")
	}

	return data, err
}

func (service *menteeService) GetById(id int) (data mentee.Core, err error) {
	data, err = service.menteeRepository.GetById(id)
	if err != nil {
		log.Error(err.Error())
		return mentee.Core{}, helper.ServiceErrorMsg(err)
	}
	return data, err

}

func (service *menteeService) Update(input mentee.Core, id int) error {
	// validasi user dgn id path param, apakah ada datanya di database
	_, errFindId := service.menteeRepository.GetById(id)
	if errFindId != nil {
		log.Error(errFindId.Error())
		return helper.ServiceErrorMsg(errFindId)
	}

	// proses
	err := service.menteeRepository.Update(input, id)
	if err != nil {
		log.Error(err.Error())
		return helper.ServiceErrorMsg(err)
	}

	return nil
}

func (service *menteeService) Delete(id int) error {
	// validasi user dgn id path param, apakah ada datanya di database
	_, errFindId := service.menteeRepository.GetById(id)
	if errFindId != nil {
		log.Error(errFindId.Error())
		return helper.ServiceErrorMsg(errFindId)
	}

	// proses
	err := service.menteeRepository.Delete(id)
	if err != nil {
		log.Error(err.Error())
		return helper.ServiceErrorMsg(err)
	}
	return nil
}
