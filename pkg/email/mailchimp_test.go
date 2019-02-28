package email_test

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/joincivil/go-common/pkg/email"
)

const (
	mailchimpApiKeyEnvVar = "MAILCHIMP_TEST_KEY"

	testListId = "a02d21e80f" // Test User List in Mailchimp
)

func interfaceTest(test email.ListMemberManager) {
}

func getMailchimpKeyFromEnvVar() string {
	return os.Getenv(mailchimpApiKeyEnvVar)
}

func TestInterface(t *testing.T) {
	apiKey := getMailchimpKeyFromEnvVar()
	if apiKey == "" {
		t.Log("No MAILCHIMP_TEST_KEY set, skipping mailchimp test")
		return
	}

	// Ensure we are properly honoring the ListMemberManager interface
	mcAPI := email.NewMailchimpAPI(apiKey)
	interfaceTest(mcAPI)
}

func TestMailchimpAddExistsRemove(t *testing.T) {
	apiKey := getMailchimpKeyFromEnvVar()
	if apiKey == "" {
		t.Log("No MAILCHIMP_TEST_KEY set, skipping mailchimp test")
		return
	}

	mcAPI := email.NewMailchimpAPI(apiKey)
	rand.Seed(time.Now().UnixNano())
	testEmail := fmt.Sprintf("testuser%d@civil.co", rand.Intn(500))

	// Ensure it is unsubscribed on the list
	_ = mcAPI.UnsubscribeFromList(testListId, testEmail, true)

	// Should not have existed at first
	subscribed, err := mcAPI.IsSubscribedToList(testListId, testEmail)
	if err != nil {
		t.Errorf("Should not have gotten error for subscribed on list: err: %v", err)
	}

	if subscribed {
		t.Errorf("Email should not have existed")
	}

	// Add it to the list
	err = mcAPI.SubscribeToList(testListId, testEmail, nil)
	if err != nil {
		t.Errorf("Should not have gotten error for add to list: err: %v", err)
	}

	// Check to see that it in fact subscribed
	subscribed, err = mcAPI.IsSubscribedToList(testListId, testEmail)
	if err != nil {
		t.Errorf("Should not have gotten error for exists on list: err: %v", err)
	}

	if !subscribed {
		t.Errorf("Email should have been on the list")
	}

	// Try to subscribe this user again
	// Will just update to subscribed for this user.
	err = mcAPI.SubscribeToList(testListId, testEmail, nil)
	if err != nil {
		t.Errorf("Should not have gotten error for add duplicate to list: err: %v", err)
	}

	// Remove it from the list
	err = mcAPI.UnsubscribeFromList(testListId, testEmail, false)
	if err != nil {
		t.Errorf("Should not have gotten error for unsubscribe from list: err: %v", err)
	}

	// Ensure it has been removed properly
	subscribed, err = mcAPI.IsSubscribedToList(testListId, testEmail)
	if err != nil {
		t.Errorf("Should not have gotten error for subscribed on list: err: %v", err)
	}

	if subscribed {
		t.Errorf("Email should not have existed")
	}

	err = mcAPI.UnsubscribeFromList(testListId, testEmail, true)
	if err != nil {
		t.Errorf("Should not have gotten error for permanent remove from list: err: %v", err)
	}
}

func TestMailchimpSubscriberWithTags(t *testing.T) {
	apiKey := getMailchimpKeyFromEnvVar()
	if apiKey == "" {
		t.Log("No MAILCHIMP_TEST_KEY set, skipping mailchimp test")
		return
	}

	mcAPI := email.NewMailchimpAPI(apiKey)
	rand.Seed(time.Now().UnixNano())
	testEmail := fmt.Sprintf("testuser%d@civil.co", rand.Intn(500))

	// Ensure it is unsubscribed on the list
	_ = mcAPI.UnsubscribeFromList(testListId, testEmail, true)

	testTag := email.Tag("Test Tag")
	tags := []email.Tag{testTag}
	subParams := &email.SubscriptionParams{Tags: tags}

	// Add it to the list
	err := mcAPI.SubscribeToList(testListId, testEmail, subParams)
	if err != nil {
		t.Errorf("Should not have gotten error for add to list: err: %v", err)
	}

	var subscribed bool

	// Check to see that it in fact subscribed
	subscribed, err = mcAPI.IsSubscribedToList(testListId, testEmail)
	if err != nil {
		t.Errorf("Should not have gotten error for exists on list: err: %v", err)
	}

	if !subscribed {
		t.Errorf("Should not have still been subscribed")
	}

	err = mcAPI.UnsubscribeFromList(testListId, testEmail, false)
	if err != nil {
		t.Errorf("Should not have gotten error for permanent remove from list: err: %v", err)
	}

	// Check to see that it in fact unsubscribed
	subscribed, err = mcAPI.IsSubscribedToList(testListId, testEmail)
	if err != nil {
		t.Errorf("Should not have gotten error for exists on list: err: %v", err)
	}

	if subscribed {
		t.Errorf("Should not have still been subscribed")
	}

	// Add back it to the list
	err = mcAPI.SubscribeToList(testListId, testEmail, subParams)
	if err != nil {
		t.Errorf("Should not have gotten error for add to list: err: %v", err)
	}

	// Check to see that it in fact subscribed
	subscribed, err = mcAPI.IsSubscribedToList(testListId, testEmail)
	if err != nil {
		t.Errorf("Should not have gotten error for exists on list: err: %v", err)
	}

	if !subscribed {
		t.Errorf("Should have been subscribed")
	}

	member, err := mcAPI.GetListMember(testListId, testEmail)
	if err != nil {
		t.Errorf("Should not have gotten error for get subscriber: err: %v", err)
	}

	if member.Tags == nil || len(member.Tags) == 0 {
		t.Errorf("Should have gotten some tags")
	}

	if member.Tags[0] != testTag {
		t.Errorf("Should have gotten the newsroom signup tag")
	}

	err = mcAPI.UnsubscribeFromList(testListId, testEmail, true)
	if err != nil {
		t.Errorf("Should not have gotten error for permanent remove from list: err: %v", err)
	}
}
