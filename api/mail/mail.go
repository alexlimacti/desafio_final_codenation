//Before use it see sendgrid docs
package main

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

//Send EMAIL
func message(p contact) (int, error) {
	from := mail.NewEmail("UATI", "uati@uatibank.com")
	subject := p.subject
	to := mail.NewEmail(p.user, p.email)
	plainTextContent := p.plain
	htmlContent := p.html
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		return 0, err
	}
	return response.StatusCode, nil
}

type contact struct {
	user    string
	email   string
	subject string
	plain   string
	html    string
}

func main() {
	p := contact{"Marcos", "marcos.pedro.nunes@gmail.com", "teste", "isso é só um teste", "<h1>isso é só um teste</h1>"}
	lg, err := message(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(lg)
}
