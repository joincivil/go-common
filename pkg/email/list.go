package email

import "time"

// Status represents a status of a member in the mailing list service
type Status string

const (
	// StatusSubscribed means a member is subscribed
	StatusSubscribed Status = "subscribed"

	// StatusUnsubscribed means a member is unsubscribed
	StatusUnsubscribed = "unsubscribed"

	// StatusInvalid means a member's status is invalid
	StatusInvalid = "invalid"

	// StatusPending means a member is pending or waiting of opt in
	StatusPending = "pending"
)

// Tag is a tag on a mailing list member
type Tag string

// ListMember represents some basic data about a member of a list.
type ListMember struct {
	ServiceID       string
	EmailAddress    string
	Status          Status
	SignupDate      time.Time
	LastUpdatedDate time.Time
	Tags            []Tag
}

// SubscriptionParams is any data can be passed along for a subscription
type SubscriptionParams struct {
	Tags []Tag
}

// ListMemberManager is an interface that represents a service that manages email lists
// consisting of members. This
type ListMemberManager interface {
	// GetListMember returns this member of the given list
	GetListMember(listID string, email string) (*ListMember, error)
	// IsSubscribedToList checks if email is subscribed to the given list
	IsSubscribedToList(listID string, email string) (bool, error)
	// SubscribeToList subscribes the email to a given list using the params
	SubscribeToList(listID string, email string, params *SubscriptionParams) error
	// UnsubscribeFromList unsubscribes the email from the given list. If delete
	// flag is true, will permanently delete from list.
	UnsubscribeFromList(listID string, email string, delete bool) error
}
