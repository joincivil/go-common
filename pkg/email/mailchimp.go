// build +integration

package email

// NOTE(PN): Mainly use this Mailchimp library to add addresses to mailing lists for.
// Mainly use the Emailer via Sendgrid for delivery of emails.

import (
	"github.com/mattbaird/gochimp"
)

// NewMailchimpAPI is a convenience function to instantiate a new MailChimpAPI
// struct
func NewMailchimpAPI(apiKey string) *MailchimpAPI {
	chimpAPI := gochimp.NewChimp(apiKey, true)
	return &MailchimpAPI{
		api:    chimpAPI,
		apiKey: apiKey,
	}
}

// MailchimpAPI is a wrapper around the MailChimp API.  Uses the gochimp lib.
type MailchimpAPI struct {
	apiKey string
	api    *gochimp.ChimpAPI
}

// IsSubscribedToList returns true if an email is subscribed on a specified list
func (m *MailchimpAPI) IsSubscribedToList(listID string, email string) (bool, error) {
	chimpEmail := gochimp.Email{Email: email}
	memberInfo := gochimp.ListsMemberInfo{
		ListId: listID,
		Emails: []gochimp.Email{chimpEmail},
	}

	res, err := m.api.MemberInfo(memberInfo)
	if err != nil {
		return false, err
	}

	if res.SuccessCount <= 0 {
		return false, nil
	}

	for _, info := range res.MemberInfoRecords {
		if info.Email == email {
			if info.Status == "subscribed" {
				return true, nil
			}
			break
		}
	}

	return false, nil
}

// SubscribeToList adds an email address to a specified list
func (m *MailchimpAPI) SubscribeToList(listID string, email string) error {
	chimpEmail := gochimp.Email{Email: email}
	subscriber := gochimp.ListsSubscribe{
		ListId:      listID,
		Email:       chimpEmail,
		DoubleOptIn: false,
	}

	_, err := m.api.ListsSubscribe(subscriber)
	return err
}

// UnsubscribeFromList removes an email address from a specified list
func (m *MailchimpAPI) UnsubscribeFromList(listID string, email string, delete bool) error {
	chimpEmail := gochimp.Email{Email: email}
	unsub := gochimp.ListsUnsubscribe{
		ListId:       listID,
		Email:        chimpEmail,
		DeleteMember: delete,
	}

	return m.api.ListsUnsubscribe(unsub)
}
