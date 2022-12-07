package delivery

import (
	"api-alta-dashboard/features/mentee"
)

type MenteeResponse struct {
	ID                uint   `json:"id"`
	Name              string `json:"name"`
	Status            string `json:"status"`
	Gender            string `json:"gender"`
	Nickname          string `json:"nickname"`
	Address           string `json:"address"`
	HomeAddress       string `json:"home_address"`
	Email             string `json:"email"`
	Telegram          string `json:"telegram"`
	Discord           string `json:"discord"`
	Phone             string `json:"phone"`
	EmergencyName     string `json:"emergency_name"`
	EmergencyPhone    string `json:"emergency_phone"`
	EmergencyStatus   string `json:"emergency_status"`
	EducationType     string `json:"education_type"`
	EducationMajor    string `json:"education_major"`
	EducationGraduate string `json:"education_graduate"`
	ClassName         string `json:"class_name"`
	ClassID           uint   `json:"class_id"`
}

type MenteeListResponse struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	ClassName     string `json:"class_name"`
	Status        string `json:"status"`
	Gender        string `json:"gender"`
	EducationType string `json:"education_type"`
	ClassID       uint   `json:"class_id"`
}

func toResponseMentee(dataCore mentee.Core) MenteeResponse {
	return MenteeResponse{
		ID:                dataCore.ID,
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
		ClassName:         "Get Class name",
		ClassID:           dataCore.ClassID,
	}
}

func toResponseMenteeList(dataCore mentee.Core) MenteeListResponse {
	return MenteeListResponse{
		ID:            dataCore.ID,
		Name:          dataCore.Name,
		ClassName:     "Get Class from ID",
		Status:        dataCore.Status,
		Gender:        dataCore.Gender,
		EducationType: dataCore.EducationType,
		ClassID:       dataCore.ClassID,
	}
}

func toResponseMenteeArray(dataCore []mentee.Core) []MenteeResponse {
	var dataResponse []MenteeResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, toResponseMentee(v))
	}
	return dataResponse
}

func toResponseMenteeListArray(dataCore []mentee.Core) []MenteeListResponse {
	var dataResponse []MenteeListResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, toResponseMenteeList(v))
	}
	return dataResponse
}
