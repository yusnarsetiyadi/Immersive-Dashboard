package class

import (
	user "api-alta-dashboard/features/user"
	"time"
)

type Core struct {
	ID        uint
	Name      string
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
	User      user.Core
}

type ServiceInterface interface {
	GetAllClass(query string) (data []Core, err error)
	CreateClass(input Core) error
	GetByIdClass(id int) (data Core, err error)
	UpdateClass(input Core, id int) error
	DeleteClass(id int) error
}

type RepositoryInterface interface {
	GetAllClass() (data []Core, err error)
	GetAllWithSearchClass(query string) (data []Core, err error)
	CreateClass(input Core) error
	GetByIdClass(id int) (data Core, err error)
	UpdateClass(input Core, id int) error
	DeleteClass(id int) error
}
