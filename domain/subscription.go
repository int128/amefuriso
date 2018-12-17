package domain

import (
	"fmt"
	"github.com/google/uuid"
)

type Subscription struct {
	ID           SubscriptionID
	Location     Location
	Notification Notification
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
