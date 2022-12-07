package delivery

import "api-alta-dashboard/features/log"

type LogResponse struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Feedback string `json:"feedback"`
	Status   string `json:"status"`
	UserID   uint   `json:"user_id"`
	MenteeID uint   `json:"mentee_id"`
}

func fromCore(dataCore log.Core) LogResponse {
	return LogResponse{
		ID:       dataCore.ID,
		Title:    dataCore.Title,
		Feedback: dataCore.Feedback,
		Status:   dataCore.Status,
		UserID:   dataCore.UserID,
		MenteeID: dataCore.MenteeID,
	}
}

func fromCoreList(dataCore []log.Core) []LogResponse {
	var dataResponse []LogResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
