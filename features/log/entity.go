package log

import (
	mentee "api-alta-dashboard/features/mentee"
	user "api-alta-dashboard/features/user"
	"time"
)

type Core struct {
	ID        uint
	Title     string
	Feedback  string
	Status    string
	UserID    uint
	MenteeID  uint
	CreatedAt time.Time
	UpdatedAt time.Time
	User      user.Core
	Mentee    mentee.Core
}

type ServiceInterface interface {
	GetAllLog(query string) (data []Core, err error)
	CreateLog(input Core) error
	UpdateLog(input Core, id int) error
	DeleteLog(id int) error
}

type RepositoryInterface interface {
	GetAllLog() (data []Core, err error)
	GetAllWithSearchLog(query string) (data []Core, err error)
	CreateLog(input Core) error
	UpdateLog(input Core, id int) error
	DeleteLog(id int) error
}
