// build +integration

package email_test

// To test this, add SENDGRID_TEST_KEY=<sendgrid key> to your environment before
// running, or add that var inline to your test command.

import (
	"os"
	"testing"

	"github.com/joincivil/go-common/pkg/email"
)

const (
	sendGridKeyEnvVar = "SENDGRID_TEST_KEY"

	useSandbox = true
)

func getSendGridKeyFromEnvVar() string {
	return os.Getenv(sendGridKeyEnvVar)
}

func TestSimpleEmailSend(t *testing.T) {
	sendGridKey := getSendGridKeyFromEnvVar()
	if sendGridKey == "" {
		t.Log("No SENDGRID_TEST_KEY set, skipping sendgrid test")
		return
	}

	emailer := email.NewEmailerWithSandbox(sendGridKey, useSandbox)

	req := &email.SendEmailRequest{
		ToName:    "Peter Ng",
		ToEmail:   "peter@civil.co",
		FromName:  "The Civil Media Company",
		FromEmail: "support@civil.co",
		Subject:   "Testing Emailer",
		Text:      "This is test text.\nThis is another line of test text.\tTabbed text",
		HTML:      "<p>This is a paragraph</p><p>This is another paragraph</p><p><b>TESTTESTSETSTSETEST!</b></p>",
	}
	err := emailer.SendEmail(req)
	if err != nil {
		t.Errorf("Should have sent the email: err: %v", err)
	}
}

func TestTemplateEmailSend(t *testing.T) {
	sendGridKey := getSendGridKeyFromEnvVar()
	if sendGridKey == "" {
		t.Log("No SENDGRID_TEST_KEY set, skipping sendgrid test")
		return
	}

	emailer := email.NewEmailerWithSandbox(sendGridKey, useSandbox)

	templateData := email.TemplateData{}
	templateData["name"] = "Peter Ng"
	templateData["subject"] = "Testing Emailer with Template"
	templateData["preheader"] = "Preheader Test"

	req := &email.SendTemplateEmailRequest{
		ToName:       "Peter Ng",
		ToEmail:      "peter@civil.co",
		FromName:     "The Civil Media Company",
		FromEmail:    "support@civil.co",
		TemplateID:   "d-19c58510201a4deab1ec6634632ccd11",
		TemplateData: templateData,
	}
	err := emailer.SendTemplateEmail(req)
	if err != nil {
		t.Errorf("Should have sent the email: err: %v", err)
	}
}
