package core

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/aprialgatto/internal/utils/events"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

func (c *core) firebaseInit() {
	ctx := context.Background()
	var err error
	opt := option.WithCredentialsFile("/home/robot/data/aprialgatto-firebase-adminsdk.json")
	//opt := option.WithCredentialsFile("/Users/eliofrancesconi/Documents/Projects/MY/aprialgatto/data/aprialgatto-firebase-adminsdk.json")
	c.app, err = firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	c.client, err = c.app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	actx := context.Background()
	c.aclient, err = c.app.Auth(actx)
	if err != nil {
		log.Fatalln(err)
	}

	mctx := context.Background()
	c.mclient, err = c.app.Messaging(mctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}
	events.Sub(VerifyAuthToken, c.Authenticator)

}
func (c *core) Authenticator(event events.Event) events.Result {
	// Obtain a messaging.Client from the App.
	switch msg := event.Payload.(type) {
	case *VerifyAuthTokenEvt:
		//log.Infof("token: %s", msg.TokenID)
		_, err := c.aclient.VerifyIDToken(context.Background(), msg.TokenID)
		if err != nil {
			return events.Errorf("error verifying ID token")
		}

		return events.OK()
	}
	return events.OK()
}

func (c *core) SendMessage(topic, action string) error {
	c.Lock()
	defer c.Unlock()
	// See documentation on defining a message payload.
	message := &messaging.Message{
		Android: &messaging.AndroidConfig{
			Data: map[string]string{
				"action": action,
			},
		},
		Topic: topic,
	}

	ctx := context.Background()
	// Send a message to the device corresponding to the provided
	// registration token.
	response, err := c.mclient.Send(ctx, message)
	if err != nil {
		log.Error(err)
		return err
	}
	// Response is a message ID string.
	log.Debugf("Successfully sent message:%v", response)
	return nil
}
