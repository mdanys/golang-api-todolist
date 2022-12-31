package todo

import "time"

type Core struct {
	ID              uint
	ActivityGroupID uint
	Title           string
	IsActive        bool
	Priority        string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Repository interface {
	ShowAll(query string) ([]Core, error)
	ShowOne(id uint) (Core, error)
	Insert(data Core) (Core, error)
	Edit(data Core, id uint) (Core, error)
	Remove(id uint) (Core, error)
}

type Service interface {
	GetAll(query string) ([]Core, error)
	GetOne(id uint) (Core, error)
	Create(data Core) (Core, error)
	Update(data Core, id uint) (Core, error)
	Delete(id uint) (Core, error)
}
