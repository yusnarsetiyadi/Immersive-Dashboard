package delivery

import (
	"api-alta-dashboard/features/class"
)

type ClassResponse struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	UserID uint   `json:"user_id"`
}

func fromCore(dataCore class.Core) ClassResponse {
	return ClassResponse{
		ID:     dataCore.ID,
		Name:   dataCore.Name,
		UserID: dataCore.UserID,
	}
}

func fromCoreList(dataCore []class.Core) []ClassResponse {
	var dataResponse []ClassResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
