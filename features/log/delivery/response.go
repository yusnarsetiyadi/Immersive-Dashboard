package delivery

import (
	"api-alta-dashboard/features/log"
	"time"
)

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

func fromCore(dataCore log.Core) LogResponse {
	return LogResponse{
		ID:           dataCore.ID,
		Title:        dataCore.Title,
		FullNameUser: dataCore.User.FullName,
		CreatedDate:  dataCore.CreatedAt.Format(time.RFC822),
		Feedback:     dataCore.Feedback,
		Status:       dataCore.Status,
		UserID:       dataCore.UserID,
		MenteeID:     dataCore.MenteeID,
	}
}

func fromCoreList(dataCore []log.Core) []LogResponse {
	var dataResponse []LogResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
