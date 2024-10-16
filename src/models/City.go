package models

import (
	"github.com/guregu/null"
)

// Represent the city
type City struct {
	ID        string     `json:"id,omitempty"`
	Name      string     `json:"name,omitempty" validate:"required"`
	UF        string     `json:"uf,omitempty" validate:"required,oneof=AC PB AL PA AP PE AM PI BA RJ CE RN ES RO GO RR MA RS MT SC MS SE MG SP TO PR DF"`
	CreatedAt *null.Time `json:"createdAt,omitempty"`
}
