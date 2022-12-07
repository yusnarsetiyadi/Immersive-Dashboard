package service

import (
	"api-alta-dashboard/features/log"
	"api-alta-dashboard/utils/helper"

	"github.com/go-playground/validator/v10"
)

type logService struct {
	logRepository log.RepositoryInterface
	validate      *validator.Validate
}

func New(repo log.RepositoryInterface) log.ServiceInterface {
	return &logService{
		logRepository: repo,
		validate:      validator.New(),
	}
}

func (service *logService) CreateLog(input log.Core) (err error) {
	errCreate := service.logRepository.CreateLog(input)
	if errCreate != nil {
		return errCreate
	}

	return nil
}

func (service *logService) GetAllLog(query string) (data []log.Core, err error) {
	if query == "" {
		data, err = service.logRepository.GetAllLog()
	} else {
		data, err = service.logRepository.GetAllWithSearchLog(query)
	}
	if err != nil {
		helper.LogDebug(err)
		return nil, helper.ServiceErrorMsg(err)
	}
	return data, err
}

func (service *logService) UpdateLog(input log.Core, id int) error {
	errUpdate := service.logRepository.UpdateLog(input, id)
	if errUpdate != nil {
		return errUpdate
	}

	return nil
}

func (service *logService) DeleteLog(id int) error {
	errDelete := service.logRepository.DeleteLog(id)
	if errDelete != nil {
		return errDelete
	}
	return nil
}
