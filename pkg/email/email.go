package email

import (
	log "github.com/golang/glog"

	sendgrid "github.com/sendgrid/sendgrid-go"

	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// NewEmailer is a convenience function that returns a new Emailer struct with
// the given apiKey
func NewEmailer(apiKey string) *Emailer {
	return &Emailer{
		sendGridClient: sendgrid.NewSendClient(apiKey),
	}
}

// Emailer is a struct that wraps around an Email provider and makes it a little
// easier to use.
type Emailer struct {
	sendGridClient *sendgrid.Client
}

// SendEmailRequest provides all the parameters to SendEmail to deliver an email.
type SendEmailRequest struct {
	ToName    string
	ToEmail   string
	FromName  string
	FromEmail string
	Subject   string
	Text      string
	HTML      string
}

// SendEmail sends a basic email via the Email provider
func (e *Emailer) SendEmail(req *SendEmailRequest) error {
	from := mail.NewEmail(req.FromName, req.FromEmail)
	to := mail.NewEmail(req.ToName, req.ToEmail)
	msg := mail.NewSingleEmail(from, req.Subject, to, req.Text, req.HTML)

	resp, err := e.sendGridClient.Send(msg)
	if err != nil {
		log.Errorf("Error sending email: err: %v", err)
		return err
	}
	log.Infof("sendemail: response status: %v", resp.StatusCode)
	log.Infof("sendemail: response body: %v", resp.Body)
	log.Infof("sendemail: response headers: %v", resp.Headers)
	return nil
}

// TemplateData represents the key-value data for the template
type TemplateData map[string]string

// SendTemplateEmailRequest provides all the parameters to SendTemplateEmail to
// deliver an templated email.
type SendTemplateEmailRequest struct {
	ToName       string
	ToEmail      string
	FromName     string
	FromEmail    string
	TemplateID   string
	TemplateData TemplateData
	AsmGroupID   int
}

// SendTemplateEmail sends an email based on a template in the email provider.
func (e *Emailer) SendTemplateEmail(req *SendTemplateEmailRequest) error {
	msg := mail.NewV3Mail()

	from := mail.NewEmail(req.FromName, req.FromEmail)
	to := mail.NewEmail(req.ToName, req.ToEmail)

	msg.SetFrom(from)
	msg.SetReplyTo(from)
	msg.SetTemplateID(req.TemplateID)

	p := mail.NewPersonalization()
	p.AddTos(to)
	for key, val := range req.TemplateData {
		p.SetDynamicTemplateData(key, val)
	}

	msg.AddPersonalizations(p)

	if req.AsmGroupID != 0 {
		a := mail.NewASM()
		a.SetGroupID(req.AsmGroupID)
		msg.SetASM(a)
	}

	resp, err := e.sendGridClient.Send(msg)
	if err != nil {
		log.Errorf("Error sending email: err: %v", err)
		return err
	}
	log.Infof("sendemail: response status: %v", resp.StatusCode)
	log.Infof("sendemail: response body: %v", resp.Body)
	log.Infof("sendemail: response headers: %v", resp.Headers)
	return nil
}
