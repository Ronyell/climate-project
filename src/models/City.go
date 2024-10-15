package models

import (
	"time"
)

// Represent the city
type City struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty" validate:"required"`
	UF        string    `json:"uf,omitempty" validate:"required,oneof=AC PB AL PA AP PE AM PI BA RJ CE RN ES RO GO RR MA RS MT SC MS SE MG SP TO PR DF"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}
