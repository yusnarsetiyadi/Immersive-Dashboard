package user

import "time"

type Core struct {
	ID        uint
	FullName  string
	Email     string `validate:"required,email"`
	Password  string `valudate:"required"`
	Team      string
	Role      string `valudate:"required"`
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ServiceInterface interface {
	GetAll(query string) (data []Core, err error)
	Create(input Core) error
	GetById(id int) (data Core, err error)
	Update(input Core, id int) error
	Delete(id int) error
}

type RepositoryInterface interface {
	GetAll() (data []Core, err error)
	GetAllWithSearch(query string) (data []Core, err error)
	Create(input Core) (row int, err error)
	GetById(id int) (data Core, err error)
	Update(input Core, id int) (row int, err error)
	Delete(id int) (row int, err error)
	FindUser(email string) (data Core, row int, err error)
}
