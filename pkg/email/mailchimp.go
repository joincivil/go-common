// build +integration

package email

import (
	"crypto/md5" // nolint: gosec
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	mailchimp "github.com/beeker1121/mailchimp-go"
	"github.com/beeker1121/mailchimp-go/lists/members"
)

// NOTE(PN): Mainly use this Mailchimp library to add addresses to mailing lists for.
// Use the Emailer via Sendgrid for delivery of emails.
// TODO(PN): Abstract away the services a bit more, this could just be an
// implementation of a mailing list interface

// NewMailchimpAPI is a convenience function to instantiate a new MailChimpAPI
// struct
func NewMailchimpAPI(apiKey string) *MailchimpAPI {
	return &MailchimpAPI{
		apiKey: apiKey,
	}
}

// MailchimpTag is a Mailchimp tag
type MailchimpTag string

const (
	// MailchimpTagNewsroomSignup is the Mailchimp tag to indicate the user is signed up
	// from the newsroom signup module
	MailchimpTagNewsroomSignup MailchimpTag = "Newsroom Signup"

	// MailchimpTagTokenStorefront is the Mailchimp tag to indicate the user is signed up
	// from the token storefront module
	MailchimpTagTokenStorefront = "Token Storefront"
)

const (
	// ErrTitleResourceNotFound is the error response from Mailchimp
	ErrTitleResourceNotFound = "resource not found"
	// ErrTitleMemberExists is the error response from Mailchimp
	ErrTitleMemberExists = "member exists"
)

// MailchimpAPI is a wrapper around the MailChimp API.  Uses the gochimp lib.
type MailchimpAPI struct {
	apiKey string
}

// GetListMember returns a Mailchimp mailing list member
func (m *MailchimpAPI) GetListMember(listID string, email string) (*Member, error) {
	err := mailchimp.SetKey(m.apiKey)
	if err != nil {
		return nil, err
	}

	return m.getMember(listID, m.mailchimpUserHash(email), nil)
}

// IsSubscribedToList returns true if an email is subscribed on a specified list
func (m *MailchimpAPI) IsSubscribedToList(listID string, email string) (bool, error) {
	err := mailchimp.SetKey(m.apiKey)
	if err != nil {
		return false, err
	}

	member, err := m.getMember(listID, m.mailchimpUserHash(email), nil)
	if err != nil {
		if m.isMemberNotFoundError(err) {
			return false, nil
		}
		return false, err
	}

	if member == nil {
		return false, nil
	}

	if member.Status != members.StatusSubscribed {
		return false, nil
	}

	return true, nil
}

// SubscribeToList adds an email address to a specified list. Only adds tags on
// creation, not re-subscribes
func (m *MailchimpAPI) SubscribeToList(listID string, email string, tags []MailchimpTag) error {
	err := mailchimp.SetKey(m.apiKey)
	if err != nil {
		return err
	}

	params := &NewMemberParams{
		EmailAddress: strings.ToLower(email),
		Status:       members.StatusSubscribed,
	}
	if len(tags) > 0 {
		params.Tags = tags
	}

	member, err := m.newMember(listID, params)
	if err != nil {
		// If member is already on list, update user to subscribed
		if !m.isMemberAlreadyOnListError(err) {
			return err
		}

		params := &members.UpdateParams{
			Status: members.StatusSubscribed,
		}
		member, err := members.Update(listID, m.mailchimpUserHash(email), params)
		if err != nil {
			return err
		}

		if member == nil {
			return fmt.Errorf("No member returns on update")
		}
		return nil
	}

	if member == nil {
		return fmt.Errorf("No member returns on subscribe")
	}

	return nil
}

// UnsubscribeFromList unsubscribes an email address from a specified list.  The
// delete flag will completely delete it from a list.
func (m *MailchimpAPI) UnsubscribeFromList(listID string, email string, delete bool) error {
	err := mailchimp.SetKey(m.apiKey)
	if err != nil {
		return err
	}

	if delete {
		return members.Delete(listID, m.mailchimpUserHash(email))
	}

	params := &members.UpdateParams{
		Status: members.StatusUnsubscribed,
	}

	_, err = members.Update(listID, m.mailchimpUserHash(email), params)
	return err
}

func (m *MailchimpAPI) mailchimpUserHash(emailAddress string) string {
	// md5 hash of lowercase version of the email address
	h := md5.New()                                 // nolint: gosec
	h.Write([]byte(strings.ToLower(emailAddress))) // nolint: gosec
	return hex.EncodeToString(h.Sum(nil))
}

func (m *MailchimpAPI) isMemberNotFoundError(err error) bool {
	e, ok := err.(*mailchimp.APIError)
	if ok {
		if e.Status == 404 && strings.ToLower(e.Title) == ErrTitleResourceNotFound {
			return true
		}
	}
	return false
}

func (m *MailchimpAPI) isMemberAlreadyOnListError(err error) bool {
	e, ok := err.(*mailchimp.APIError)
	if ok {
		if e.Status == 400 && strings.ToLower(e.Title) == ErrTitleMemberExists {
			return true
		}
	}
	return false
}

// NOTE(PN): The items below extend the mailchimp-go library to include tags
// Will try and submit a PR into that repo to include these changes when time permits

// NewMemberParams is a version of NewParams that adds a tags field
type NewMemberParams struct {
	EmailType       members.EmailType      `json:"email_type,omitempty"`
	Status          members.Status         `json:"status"`
	MergeFields     map[string]interface{} `json:"merge_fields,omitempty"`
	Interests       map[string]bool        `json:"interests,omitempty"`
	Language        string                 `json:"language,omitempty"`
	VIP             bool                   `json:"vip,omitempty"`
	Location        *members.Location      `json:"location,omitempty"`
	IPSignup        string                 `json:"ip_signup,omitempty"`
	TimestampSignup time.Time              `json:"timestamp_signup,omitempty"`
	IPOpt           string                 `json:"ip_opt,omitempty"`
	TimestampOpt    time.Time              `json:"timestamp_opt,omitempty"`
	EmailAddress    string                 `json:"email_address"`
	Tags            []MailchimpTag         `json:"tags,omitempty"`
}

// MemberTag represents a tag on a Member struct
type MemberTag struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// Member defines a single member within a list that adds a field for tags
type Member struct {
	ID              string                 `json:"id"`
	EmailAddress    string                 `json:"email_address"`
	UniqueEmailID   string                 `json:"unique_email_id"`
	EmailType       members.EmailType      `json:"email_type,omitempty"`
	Status          members.Status         `json:"status"`
	MergeFields     map[string]interface{} `json:"merge_fields,omitempty"`
	Interests       map[string]bool        `json:"interests,omitempty"`
	Stats           *members.Stats         `json:"stats,omitempty"`
	IPSignup        string                 `json:"ip_signup,omitempty"`
	TimestampSignup time.Time              `json:"timestamp_signup,omitempty"`
	IPOpt           string                 `json:"ip_opt,omitempty"`
	TimestampOpt    time.Time              `json:"timestamp_opt,omitempty"`
	MemberRating    uint8                  `json:"member_rating,omitempty"`
	LastChanged     time.Time              `json:"last_changed,omitempty"`
	Language        string                 `json:"language,omitempty"`
	VIP             bool                   `json:"vip,omitempty"`
	EmailClient     string                 `json:"email_client,omitempty"`
	Location        *members.Location      `json:"location,omitempty"`
	LastNote        *members.Note          `json:"last_note,omitempty"`
	ListID          string                 `json:"list_id"`
	Tags            []*MemberTag           `json:"tags,omitempty"`
}

// newMember adds a new list member with tags
func (m *MailchimpAPI) newMember(listID string, params *NewMemberParams) (*Member, error) {
	res := &Member{}
	path := fmt.Sprintf("lists/%s/members", listID)

	if params == nil {
		if err := mailchimp.Call("POST", path, nil, nil, res); err != nil {
			return nil, err
		}
		return res, nil
	}

	if err := mailchimp.Call("POST", path, nil, params, res); err != nil {
		return nil, err
	}
	return res, nil
}

// getMember retrieves information about a specific member within a list with tags
func (m *MailchimpAPI) getMember(listID, hash string, params *members.GetMemberParams) (*Member, error) {
	res := &Member{}
	path := fmt.Sprintf("lists/%s/members/%s", listID, hash)

	if params == nil {
		if err := mailchimp.Call("GET", path, nil, nil, res); err != nil {
			return nil, err
		}
		return res, nil
	}

	if err := mailchimp.Call("GET", path, params, nil, res); err != nil {
		return nil, err
	}
	return res, nil
}
