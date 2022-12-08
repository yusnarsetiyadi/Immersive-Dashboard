package delivery

import (
	"api-alta-dashboard/features/class"
	"api-alta-dashboard/features/user"
)

type ClassResponse struct {
	ID        uint   `json:"id"`
	ClassName string `json:"class_name"`
	User      user.Core
}

func fromCore(dataCore class.Core) ClassResponse {
	return ClassResponse{
		ID:        dataCore.ID,
		ClassName: dataCore.ClassName,
		User:      dataCore.User,
	}
}

func fromCoreList(dataCore []class.Core) []ClassResponse {
	var dataResponse []ClassResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
