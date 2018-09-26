// build +integration

package utils_test

import (
	"testing"
	// "github.com/joincivil/civil-api-server/pkg/utils"
)

// const (
// 	sendGridEmailKey = ""
// )

func TestSimpleEmailSend(t *testing.T) {
	// emailer := utils.NewEmailer(sendGridEmailKey)

	// req := &utils.SendEmailRequest{
	// 	ToName:    "Peter Ng",
	// 	ToEmail:   "peter@civil.co",
	// 	FromName:  "The Civil Media Company",
	// 	FromEmail: "support@civil.co",
	// 	Subject:   "Testing Emailer",
	// 	Text:      "This is test text.\nThis is another line of test text.\tTabbed text",
	// 	HTML:      "<p>This is a paragraph</p><p>This is another paragraph</p><p><b>TESTTESTSETSTSETEST!</b></p>",
	// }
	// err := emailer.SendEmail(req)
	// if err != nil {
	// 	t.Errorf("Should have sent the email: err: %v", err)
	// }
}

func TestTemplateEmailSend(t *testing.T) {
	// emailer := utils.NewEmailer(sendGridEmailKey)

	// templateData := utils.TemplateData{}
	// templateData["name"] = "Peter Ng"
	// templateData["subject"] = "Testing Emailer with Template"
	// templateData["preheader"] = "Preheader Test"

	// req := &utils.SendTemplateEmailRequest{
	// 	ToName:       "Peter Ng",
	// 	ToEmail:      "peter@civil.co",
	// 	FromName:     "The Civil Media Company",
	// 	FromEmail:    "support@civil.co",
	// 	TemplateID:   "d-19c58510201a4deab1ec6634632ccd11",
	// 	TemplateData: templateData,
	// }
	// err := emailer.SendTemplateEmail(req)
	// if err != nil {
	// 	t.Errorf("Should have sent the email: err: %v", err)
	// }
}
