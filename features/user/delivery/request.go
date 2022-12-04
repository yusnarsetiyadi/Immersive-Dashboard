package delivery

import "api-alta-dashboard/features/user"

type InsertRequest struct {
	FullName string `json:"full_name" form:"full_name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Team     string `json:"team" form:"team"`
	Role     string `json:"role" form:"role"`
	Status   string `json:"status" form:"status"`
}

type UpdateRequest struct {
	ID       uint   `json:"id" form:"id"`
	FullName string `json:"full_name" form:"full_name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Team     string `json:"team" form:"team"`
	Role     string `json:"role" form:"role"`
	Status   string `json:"status" form:"status"`
}

func toCore(i interface{}) user.Core {
	switch i.(type) {
	case InsertRequest:
		cnv := i.(InsertRequest)
		return user.Core{
			FullName: cnv.FullName,
			Email:    cnv.Email,
			Password: cnv.Password,
			Team:     cnv.Team,
			Role:     cnv.Role,
			Status:   cnv.Status,
		}

	case UpdateRequest:
		cnv := i.(UpdateRequest)
		return user.Core{
			ID:       cnv.ID,
			FullName: cnv.FullName,
			Email:    cnv.Email,
			Password: cnv.Password,
			Team:     cnv.Team,
			Role:     cnv.Role,
			Status:   cnv.Status,
		}
	}

	return user.Core{}
}
