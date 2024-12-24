package models

import "time"

type Rates struct {
	ID        int       `json:"id"`
	Usd       float32   `json:"usd"`
	Eur       float32   `json:"eur"`
	CreatedAt time.Time `json:"created_at"`
}
