package models

import "time"

// Represent the city
type City struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	UF        string    `json:"uf,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}
