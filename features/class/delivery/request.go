package delivery

import (
	"api-alta-dashboard/features/class"
)

type CreateClassRequest struct {
	ClassName string `json:"class_name" form:"class_name"`
	UserID    uint   `json:"user_id" form:"user_id"`
}

type UpdateClassRequest struct {
	ID        uint   `json:"id" form:"id"`
	ClassName string `json:"class_name" form:"class_name"`
	UserID    uint   `json:"user_id" form:"user_id"`
}

func toCore(i interface{}) class.Core {
	switch i.(type) {
	case CreateClassRequest:
		cnv := i.(CreateClassRequest)
		return class.Core{
			ClassName: cnv.ClassName,
			UserID:    cnv.UserID,
		}

	case UpdateClassRequest:
		cnv := i.(UpdateClassRequest)
		return class.Core{
			ID:        cnv.ID,
			ClassName: cnv.ClassName,
			UserID:    cnv.UserID,
		}
	}

	return class.Core{}
}
