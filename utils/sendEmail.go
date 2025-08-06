package utils

import "github.com/resend/resend-go/v2"

type SendEmailProps struct {
	To      []string
	Subject string
	Html    string
}

func SendEmail(props SendEmailProps) {
	apiKey := "re_huahRdwQ_3NTq8JTfCjgSyJQZa3wgbC8m"

	client := resend.NewClient(apiKey)

	params := &resend.SendEmailRequest{
		From:    "noreply@sdyson.dev",
		To:      props.To,
		Subject: props.Subject,
		Html:    props.Html,
	}

	client.Emails.Send(params)
}
