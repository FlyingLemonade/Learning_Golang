package entities

import(
	"time"
	
)

type Category struct{
	Id uint64
	Name string
	CreatedAt time.Time
	UpdatedAt time.Time
}