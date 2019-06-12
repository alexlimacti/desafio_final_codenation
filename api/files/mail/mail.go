//Before use it see sendgrid docs

package mail

import (
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
)

func mail() int {
	from := mail.NewEmail("Example User", "marcos.pedro.nunes@gmail.com")
	subject := "Sending with SendGrid is Fun"
	to := mail.NewEmail("Example User", "marcos.pedro.nunes@gmail.com")
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		return response.StatusCode
	}
}
