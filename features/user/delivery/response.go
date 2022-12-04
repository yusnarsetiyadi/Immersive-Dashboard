package delivery

import "api-alta-dashboard/features/user"

type UserResponse struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Team     string `json:"team"`
	Role     string `json:"role"`
	Status   string `json:"status"`
}

func fromCore(dataCore user.Core) UserResponse {
	return UserResponse{
		ID:       dataCore.ID,
		FullName: dataCore.FullName,
		Email:    dataCore.Email,
		Team:     dataCore.Team,
		Role:     dataCore.Role,
		Status:   dataCore.Status,
	}
}

func fromCoreList(dataCore []user.Core) []UserResponse {
	var dataResponse []UserResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
