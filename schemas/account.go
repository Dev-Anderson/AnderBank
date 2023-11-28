package schemas

import "time"

type Account struct {
	ID            int       `json:"id"`
	NumberAccount string    `json:"numberaccount"`
	Balance       float64   `json:"balance"`
	DateCreate    time.Time `json:"datecreate"`
	DateUpdate    time.Time `json:"dateupdate"`
	DateDelete    time.Time `json:"datedelete"`
	Debit         bool      `json:"debit"`
	Credit        bool      `json:"credit"`
	Active        bool      `json:"active"`
}
