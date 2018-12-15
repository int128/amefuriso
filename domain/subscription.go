package domain

type Subscription struct {
	ID           SubscriptionID
	User         UserID
	Location     Location
	Notification Notification
}

type SubscriptionID string
