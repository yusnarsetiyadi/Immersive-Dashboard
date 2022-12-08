package delivery

import (
	"api-alta-dashboard/features/class"
)

type ClassResponse struct {
	ID           uint   `json:"id"`
	ClassName    string `json:"class_name"`
	FullNameUser string `json:"full_name_user"`
	UserId       uint   `json:"user_id"`
}

func fromCore(dataCore class.Core) ClassResponse {
	return ClassResponse{
		ID:           dataCore.ID,
		ClassName:    dataCore.ClassName,
		FullNameUser: dataCore.User.FullName,
		UserId:       dataCore.UserID,
	}
}

func fromCoreList(dataCore []class.Core) []ClassResponse {
	var dataResponse []ClassResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
