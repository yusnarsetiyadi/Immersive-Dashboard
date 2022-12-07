package delivery

import "api-alta-dashboard/features/log"

type CreateLogRequest struct {
	Title    string `json:"title" form:"title"`
	Feedback string `json:"feedback" form:"feedback"`
	Status   string `json:"status" form:"status"`
	MenteeID uint   `json:"mentee_id" form:"mentee_id"`
}

type UpdateLogRequest struct {
	ID       uint   `json:"id" form:"id"`
	Title    string `json:"title" form:"title"`
	Feedback string `json:"feedback" form:"feedback"`
	Status   string `json:"status" form:"status"`
	MenteeID uint   `json:"mentee_id" form:"mentee_id"`
}

func toCore(i interface{}) log.Core {
	switch i.(type) {
	case CreateLogRequest:
		cnv := i.(CreateLogRequest)
		return log.Core{
			Title:    cnv.Title,
			Feedback: cnv.Feedback,
			Status:   cnv.Status,
			MenteeID: cnv.MenteeID,
		}

	case UpdateLogRequest:
		cnv := i.(UpdateLogRequest)
		return log.Core{
			ID:       cnv.ID,
			Title:    cnv.Title,
			Feedback: cnv.Feedback,
			Status:   cnv.Status,
			MenteeID: cnv.MenteeID,
		}
	}

	return log.Core{}
}
