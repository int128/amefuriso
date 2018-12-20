package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type Subscription struct {
	ID          SubscriptionID
	Location    Location
	Publication Publication
}

func (subscription Subscription) String() string {
	return fmt.Sprintf("Subscription(%s)", subscription.ID)
}

type SubscriptionID string

func NewSubscription(location Location) Subscription {
	return Subscription{
		ID:       SubscriptionID(uuid.New().String()),
		Location: location,
	}
}

type ErrNoSuchSubscription struct {
	UserID         UserID
	SubscriptionID SubscriptionID
}

func (e ErrNoSuchSubscription) Error() string {
	return fmt.Sprintf("No such subscription UserID=%s, SubscriptionID=%s", e.UserID, e.SubscriptionID)
}

func IsErrNoSuchSubscription(err error) bool {
	_, ok := err.(ErrNoSuchSubscription)
	return ok
}
