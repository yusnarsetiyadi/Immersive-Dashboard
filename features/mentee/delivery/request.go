package delivery

import (
	"api-alta-dashboard/features/mentee"
)

type InsertRequest struct {
	Name              string `json:"name" form:"name"`
	Status            string `json:"status" form:"status"`
	Gender            string `json:"gender" form:"gender"`
	Nickname          string `json:"nickname" form:"nickname"`
	Address           string `json:"address" form:"address"`
	HomeAddress       string `json:"home_address" form:"home_address"`
	Email             string `json:"email" form:"email"`
	Telegram          string `json:"telegram" form:"telegram"`
	Discord           string `json:"discord" form:"discord"`
	Phone             string `json:"phone" form:"phone"`
	EmergencyName     string `json:"emergency_name" form:"emergency_name"`
	EmergencyPhone    string `json:"emergency_phone" form:"emergency_phone"`
	EmergencyStatus   string `json:"emergency_status" form:"emergency_status"`
	EducationType     string `json:"education_type" form:"education_type"`
	EducationMajor    string `json:"education_major" form:"education_major"`
	EducationGraduate string `json:"education_graduate" form:"education_major"`
	IDClass           uint   `json:"id_class" form:"id_class"`
}

type UpdateRequest struct {
	ID                uint   `json:"id" form:"id"`
	Name              string `json:"name" form:"name"`
	Status            string `json:"status" form:"status"`
	Gender            string `json:"gender" form:"gender"`
	Nickname          string `json:"nickname" form:"nickname"`
	Address           string `json:"address" form:"address"`
	HomeAddress       string `json:"home_address" form:"home_address"`
	Email             string `json:"email" form:"email"`
	Telegram          string `json:"telegram" form:"telegram"`
	Discord           string `json:"discord" form:"discord"`
	Phone             string `json:"phone" form:"phone"`
	EmergencyName     string `json:"emergency_name" form:"emergency_name"`
	EmergencyPhone    string `json:"emergency_phone" form:"emergency_phone"`
	EmergencyStatus   string `json:"emergency_status" form:"emergency_status"`
	EducationType     string `json:"education_type" form:"education_type"`
	EducationMajor    string `json:"education_major" form:"education_major"`
	EducationGraduate string `json:"education_graduate" form:"education_major"`
	IDClass           uint   `json:"id_class" form:"id_class"`
}

func toCore(i interface{}) mentee.Core {
	switch i.(type) {
	case InsertRequest:
		cnv := i.(InsertRequest)
		return mentee.Core{
			Name:              cnv.Name,
			Status:            cnv.Status,
			Gender:            cnv.Gender,
			Nickname:          cnv.Nickname,
			Address:           cnv.Address,
			HomeAddress:       cnv.HomeAddress,
			Email:             cnv.Email,
			Telegram:          cnv.Telegram,
			Discord:           cnv.Discord,
			Phone:             cnv.Phone,
			EmergencyName:     cnv.EmergencyName,
			EmergencyPhone:    cnv.EmergencyPhone,
			EmergencyStatus:   cnv.EmergencyStatus,
			EducationType:     cnv.EducationType,
			EducationMajor:    cnv.EducationMajor,
			EducationGraduate: cnv.EducationGraduate,
			IDClass:           cnv.IDClass,
		}

	case UpdateRequest:
		cnv := i.(UpdateRequest)
		return mentee.Core{
			ID:                cnv.ID,
			Name:              cnv.Name,
			Status:            cnv.Status,
			Gender:            cnv.Gender,
			Nickname:          cnv.Nickname,
			Address:           cnv.Address,
			HomeAddress:       cnv.HomeAddress,
			Email:             cnv.Email,
			Telegram:          cnv.Telegram,
			Discord:           cnv.Discord,
			Phone:             cnv.Phone,
			EmergencyName:     cnv.EmergencyName,
			EmergencyPhone:    cnv.EmergencyPhone,
			EmergencyStatus:   cnv.EmergencyStatus,
			EducationType:     cnv.EducationType,
			EducationMajor:    cnv.EducationMajor,
			EducationGraduate: cnv.EducationGraduate,
			IDClass:           cnv.IDClass,
		}
	}

	return mentee.Core{}
}
