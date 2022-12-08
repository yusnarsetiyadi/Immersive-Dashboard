package repository

import (
	_log "api-alta-dashboard/features/log/repository"
	_mentee "api-alta-dashboard/features/mentee"

	"gorm.io/gorm"
)

// struct gorm model
type Mentee struct {
	gorm.Model
	ID                uint
	Name              string
	Status            string
	Gender            string
	Nickname          string
	Address           string `validate:"required"`
	HomeAddress       string
	Email             string `validate:"required,email"`
	Telegram          string `validate:"required"`
	Discord           string
	Phone             string `validate:"required"`
	EmergencyName     string `validate:"required"`
	EmergencyPhone    string `validate:"required"`
	EmergencyStatus   string `validate:"required"`
	EducationType     string `validate:"required"`
	EducationMajor    string
	EducationGraduate string
	ClassID           uint `validate:"required"`
	Logs              []_log.Log
}

// DTO
// mapping

// mengubah struct core ke struct model gorm
func fromCore(dataCore _mentee.Core) Mentee {
	menteeGorm := Mentee{
		Name:              dataCore.Name,
		Status:            dataCore.Status,
		Gender:            dataCore.Gender,
		Nickname:          dataCore.Nickname,
		Address:           dataCore.Address,
		HomeAddress:       dataCore.HomeAddress,
		Email:             dataCore.Email,
		Telegram:          dataCore.Telegram,
		Discord:           dataCore.Discord,
		Phone:             dataCore.Phone,
		EmergencyName:     dataCore.EmergencyName,
		EmergencyPhone:    dataCore.EmergencyPhone,
		EmergencyStatus:   dataCore.EmergencyStatus,
		EducationType:     dataCore.EducationType,
		EducationMajor:    dataCore.EducationMajor,
		EducationGraduate: dataCore.EducationGraduate,
		ClassID:           dataCore.ClassID,
	}
	return menteeGorm
}

// mengubah struct model gorm ke struct core
func (dataModel *Mentee) toCore() _mentee.Core {
	return _mentee.Core{
		ID:                dataModel.ID,
		Name:              dataModel.Name,
		Status:            dataModel.Status,
		Gender:            dataModel.Gender,
		Nickname:          dataModel.Nickname,
		Address:           dataModel.Address,
		HomeAddress:       dataModel.HomeAddress,
		Email:             dataModel.Email,
		Telegram:          dataModel.Telegram,
		Discord:           dataModel.Discord,
		Phone:             dataModel.Phone,
		EmergencyName:     dataModel.EmergencyName,
		EmergencyPhone:    dataModel.EmergencyPhone,
		EmergencyStatus:   dataModel.EmergencyStatus,
		EducationType:     dataModel.EducationType,
		EducationMajor:    dataModel.EducationMajor,
		EducationGraduate: dataModel.EducationGraduate,
		ClassID:           dataModel.ClassID,
		CreatedAt:         dataModel.CreatedAt,
		UpdatedAt:         dataModel.UpdatedAt,
	}
}

// mengubah slice struct model gorm ke slice struct core
func toCoreList(dataModel []Mentee) []_mentee.Core {
	var dataCore []_mentee.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
