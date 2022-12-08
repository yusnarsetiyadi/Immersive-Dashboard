package delivery

import (
	"api-alta-dashboard/features/mentee"
	"time"
)

type MenteeResponse struct {
	ID                uint          `json:"id"`
	Name              string        `json:"name"`
	Status            string        `json:"status"`
	Gender            string        `json:"gender"`
	Nickname          string        `json:"nickname"`
	Address           string        `json:"address"`
	HomeAddress       string        `json:"home_address"`
	Email             string        `json:"email"`
	Telegram          string        `json:"telegram"`
	Discord           string        `json:"discord"`
	Phone             string        `json:"phone"`
	EmergencyName     string        `json:"emergency_name"`
	EmergencyPhone    string        `json:"emergency_phone"`
	EmergencyStatus   string        `json:"emergency_status"`
	EducationType     string        `json:"education_type"`
	EducationMajor    string        `json:"education_major"`
	EducationGraduate string        `json:"education_graduate"`
	ClassName         string        `json:"class_name"`
	ClassID           uint          `json:"class_id"`
	Logs              []LogResponse `json:"logs"`
}

type LogResponse struct {
	ID           uint   `json:"id"`
	Title        string `json:"title"`
	FullNameUser string `json:"fullname_user"`
	CreatedDate  string `json:"created_date"`
	Feedback     string `json:"feedback"`
	Status       string `json:"status"`
	MenteeID     uint   `json:"mentee_id"`
	UserID       uint   `json:"user_id"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
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

	var ArrLogs []LogResponse

	for _, val := range dataCore.Logs {
		ArrLogs = append(ArrLogs, LogResponse{
			ID:           val.ID,
			Title:        val.Title,
			FullNameUser: val.User.FullName,
			CreatedDate:  val.CreatedAt.Format(time.RFC822),
			Feedback:     val.Feedback,
			Status:       val.Status,
			MenteeID:     dataCore.ID,
			UserID:       val.User.ID,
		})
	}

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
		Logs:              ArrLogs,
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
