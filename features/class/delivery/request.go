package delivery

import (
	"api-alta-dashboard/features/class"
)

type CreateClassRequest struct {
	Name   string `json:"name" form:"name"`
	UserID uint   `json:"user_id" form:"user_id"`
}

type UpdateClassRequest struct {
	ID     uint   `json:"id" form:"id"`
	Name   string `json:"name" form:"name"`
	UserID uint   `json:"user_id" form:"user_id"`
}

func toCore(i interface{}) class.Core {
	switch i.(type) {
	case CreateClassRequest:
		cnv := i.(CreateClassRequest)
		return class.Core{
			Name:   cnv.Name,
			UserID: cnv.UserID,
		}

	case UpdateClassRequest:
		cnv := i.(UpdateClassRequest)
		return class.Core{
			ID:     cnv.ID,
			Name:   cnv.Name,
			UserID: cnv.UserID,
		}
	}

	return class.Core{}
}
