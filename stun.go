package stun

import(
	"net/smtp"
	// "fmt"
	// "os"
)

//Email struct represents the email to be sent
type Email struct{
	From string
	To string
	Subject string
	Body string
}

//NewEmail creates a new email instance
func NewEmail(to, subject, body string) *Email  {
	return &Email{
		To: to,
		Subject: subject,
		Body: body,
	}
	
}

// ** GHAT GPT **
// Send sends the email using the specified SMTP server.
// func (e *Email) Send() error {
	// Read email credentials from environment variables
	// from := os.Getenv("STUN_EMAIL_SENDER")
	// password := os.Getenv("STUN_EMAIL_PASSWORD")

	// Set up the SMTP server configuration
	// smtpHost := "smtp.gmail.com"
	// smtpPort := "587"

	// Compose the email
	// message := []byte("Subject: " + e.Subject + "\r\n\r\n" + e.Body)

	// Connect to the SMTP server
	// auth := smtp.PlainAuth("", from, password, smtpHost)
	// err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{e.To}, message)
	// if err != nil {
		// return err
	// }

	// return nil
// }

// ** GITHUB COPILOT **
// Send sends the email using the specified SMTP server.
func (e *Email) Send() error {
	// Connect to the remote SMTP server.
	client, err := smtp.Dial("smtp.gmail.com:587")
	if err != nil {
		return err
	}
	defer client.Close()

	// Set the sender and recipient.
	client.Mail(e.From)
	client.Rcpt(e.To)

	// Send the email body.
	WriteCloser, err := client.Data()
	if err != nil {
		return err
	}
	defer WriteCloser.Close()
	_, err = WriteCloser.Write([]byte(e.Body))
	if err != nil {
		return err
	}
	return nil
	
}