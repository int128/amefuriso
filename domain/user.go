package domain

type User struct {
	ID            UserID
	YahooClientID YahooClientID
}

type UserID string

type YahooClientID string
