// build +integration

package email

import (
	"crypto/md5" // nolint: gosec
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	mailchimp "github.com/beeker1121/mailchimp-go"
	"github.com/beeker1121/mailchimp-go/lists/members"
)

// NOTE(PN): Mainly use this Mailchimp library to add addresses to mailing lists for.
// Use the Emailer via Sendgrid for delivery of emails.

// Implements MailingListMemberManager

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
	// ErrTitleResourceNotFound is the error response from Mailchimp
	ErrTitleResourceNotFound = "resource not found"
	// ErrTitleMemberExists is the error response from Mailchimp
	ErrTitleMemberExists = "member exists"
)

const (
	// ServiceNameMailchimp is the service name for Mailchimp
	ServiceNameMailchimp = "mailchimp"
)

// MailchimpAPI is a wrapper around the MailChimp API.  Uses the gochimp lib.
type MailchimpAPI struct {
	apiKey string
}

// ServiceName returns the underlying mailing list service
func (m *MailchimpAPI) ServiceName() string {
	return ServiceNameMailchimp
}

// GetListMember returns a Mailchimp mailing list member
func (m *MailchimpAPI) GetListMember(listID string, email string) (*ListMember, error) {
	err := mailchimp.SetKey(m.apiKey)
	if err != nil {
		return nil, err
	}

	mcMember, err := m.getMember(listID, m.mailchimpUserHash(email), nil)
	if err != nil {
		return nil, err
	}

	var status Status

	if mcMember.Status == members.StatusSubscribed {
		status = StatusSubscribed
	} else if mcMember.Status == members.StatusUnsubscribed {
		status = StatusUnsubscribed
	} else if mcMember.Status == members.StatusCleaned {
		status = StatusInvalid
	} else {
		status = StatusPending
	}

	var tags []Tag
	if len(mcMember.Tags) > 0 {
		tags = make([]Tag, len(mcMember.Tags))
		for ind, t := range mcMember.Tags {
			tags[ind] = Tag(t.Name)
		}
	}

	return &ListMember{
		ServiceID:       mcMember.ID,
		EmailAddress:    mcMember.EmailAddress,
		Status:          status,
		SignupDate:      mcMember.TimestampSignup,
		LastUpdatedDate: mcMember.LastChanged,
		Tags:            tags,
	}, nil
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
func (m *MailchimpAPI) SubscribeToList(listID string, email string, params *SubscriptionParams) error {
	err := mailchimp.SetKey(m.apiKey)
	if err != nil {
		return err
	}

	now := time.Now().UTC()
	newParams := &NewMemberParams{
		EmailAddress:    strings.ToLower(email),
		Status:          members.StatusSubscribed,
		TimestampSignup: now,
		TimestampOpt:    now,
	}

	var mcTags []MailchimpTag

	if params != nil {
		if len(params.Tags) > 0 {
			mcTags = make([]MailchimpTag, len(params.Tags))
			for ind, val := range params.Tags {
				mcTags[ind] = MailchimpTag(val)
			}

			newParams.Tags = mcTags
		}
	}

	member, err := m.newMember(listID, newParams)
	if err != nil {
		// If member is already on list, update user to subscribed
		if !m.isMemberAlreadyOnListError(err) {
			return err
		}

		updateParams := &members.UpdateParams{
			Status: members.StatusSubscribed,
		}
		member, err := members.Update(listID, m.mailchimpUserHash(email), updateParams)
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

// MarshalJSON handles custom JSON marshalling for the NewMembersParams object.
// Added here to correct the invalid timestamp format issue.
func (np *NewMemberParams) MarshalJSON() ([]byte, error) {
	var timestampSignup string
	var timestampOpt string

	if !np.TimestampSignup.IsZero() {
		timestampSignup = np.TimestampSignup.Format(time.RFC3339)
	}
	if !np.TimestampOpt.IsZero() {
		timestampOpt = np.TimestampOpt.Format(time.RFC3339)
	}

	type alias NewMemberParams
	return json.Marshal(&struct {
		*alias
		TimestampSignup string `json:"timestamp_signup,omitempty"`
		TimestampOpt    string `json:"timestamp_opt,omitempty"`
	}{
		alias:           (*alias)(np),
		TimestampSignup: timestampSignup,
		TimestampOpt:    timestampOpt,
	})
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
