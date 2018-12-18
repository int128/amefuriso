package domain

import (
	"fmt"
	"github.com/google/uuid"
)

type User struct {
	ID            UserID
	YahooClientID YahooClientID
}

func (user User) String() string {
	return fmt.Sprintf("User(%s)", user.ID)
}

type UserID string

type YahooClientID string

func NewUser() User {
	return User{ID: UserID(uuid.New().String())}
}

type ErrNoSuchUser struct {
	ID UserID
}

func (e ErrNoSuchUser) Error() string {
	return fmt.Sprintf("No such user %s", e.ID)
}

func IsErrNoSuchUser(err error) bool {
	_, ok := err.(ErrNoSuchUser)
	return ok
}
