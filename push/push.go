package push

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
)

func Send(token string, text string, title string) int {
	var result int = 0
	ctx := context.Background()
	//opt := option.WithCredentialsFile("C:\\push\\credenciais.jsonstomach-8f980-firebase-adminsdk-fbsvc-7d143ac99b.json")
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatal(err)
		return result
	}
	if app == nil {
		log.Fatal(err)
		return result
	}
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatal(err)
		return result
	}
	if client == nil {
		log.Fatal(err)
		return result
	}
	message := &messaging.Message{
		Token: token,
		Notification: &messaging.Notification{
			Title: title,
			Body:  text,
		},
	}
	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatal(err)
	}
	if response == "" {
		log.Fatal(err)
	}
	result = 1
	return result
}
