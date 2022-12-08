package class

import (
	user "api-alta-dashboard/features/user"
	"time"
)

type Core struct {
	ID        uint
	ClassName string
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
	User      user.Core
}

type Validator struct {
	ID        uint
	ClassName string `validate:"required,unique"`
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

// mengubah struct model gorm ke struct core
func ToValidator(input Core) Validator {
	return Validator{
		ID:        input.ID,
		ClassName: input.ClassName,
		UserID:    input.UserID,
	}
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
