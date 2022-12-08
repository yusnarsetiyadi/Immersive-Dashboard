package service

import (
	"api-alta-dashboard/features/class"
	"api-alta-dashboard/utils/helper"

	"github.com/go-playground/validator/v10"
)

type classService struct {
	classRepository class.RepositoryInterface
	validate        *validator.Validate
}

func New(repo class.RepositoryInterface) class.ServiceInterface {
	return &classService{
		classRepository: repo,
		validate:        validator.New(),
	}
}

func (service *classService) CreateClass(input class.Core) (err error) {

	//inputCheck := class.ToValidator(input)

	// validasi input
	if errValidate := service.validate.Struct(input); errValidate != nil {
		helper.LogDebug(errValidate)
		return helper.ServiceErrorMsg(errValidate)
	}

	errCreate := service.classRepository.CreateClass(input)
	if errCreate != nil {
		helper.LogDebug(errCreate)
		return helper.ServiceErrorMsg(errCreate)
	}

	return nil
}

func (service *classService) GetAllClass(query string) (data []class.Core, err error) {
	if query == "" {
		data, err = service.classRepository.GetAllClass()
	} else {
		data, err = service.classRepository.GetAllWithSearchClass(query)
	}

	if err != nil {
		return nil, helper.ServiceErrorMsg(err)
	}

	return data, err
}

func (service *classService) GetByIdClass(id int) (data class.Core, err error) {
	data, err = service.classRepository.GetByIdClass(id)
	if err != nil {
		helper.LogDebug(err)
		return class.Core{}, helper.ServiceErrorMsg(err)
	}
	return data, err

}

func (service *classService) UpdateClass(input class.Core, id int) error {
	errUpdate := service.classRepository.UpdateClass(input, id)
	if errUpdate != nil {
		helper.LogDebug(errUpdate)
		return helper.ServiceErrorMsg(errUpdate)
	}

	return nil
}

func (service *classService) DeleteClass(id int) error {
	errDelete := service.classRepository.DeleteClass(id)
	if errDelete != nil {
		helper.LogDebug(errDelete)
		return helper.ServiceErrorMsg(errDelete)
	}
	return nil
}
