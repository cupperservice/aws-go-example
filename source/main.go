package main

import (
	"errors"
	"encoding/json"
	"fmt"
	"bytes"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/ashwanthkumar/slack-go-webhook"
)

var (
	// DefaultHTTPGetAddress Default Address
	DefaultHTTPGetAddress = "https://checkip.amazonaws.com"

	// ErrNoIP No IP found in response
	ErrNoIP = errors.New("No IP in HTTP response")

	// ErrNon200Response non 200 status code in response
	ErrNon200Response = errors.New("Non 200 Response found")
)

type Message struct {
	Name	string	`json:"name"`
	Key		string 	`json:"key"`
	Value	string	`json:"value"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	name := request.PathParameters["name"]
	key := request.QueryStringParameters["key"]
	value := request.QueryStringParameters["value"]

	message := Message{
		Name:		name,
		Key:		key,
		Value:	value,
	}

	storeToS3(message)
	notifyToSlack(message)

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Hello, %v", message.Name),
		StatusCode: 200,
	}, nil
}

func notifyToSlack(message Message) {
	field1 := slack.Field {Title: "title", Value: ":gopher:"}
	field2 := slack.Field {Title: message.Key, Value: message.Value}

	attachment := slack.Attachment{}
	attachment.AddField(field1).AddField(field2)
	color := "good"
	attachment.Color = &color
	payload := slack.Payload{
		Text: 				"Hello, " + message.Name,
		Attachments:	[]slack.Attachment{attachment},
	}

	slack.Send(os.Getenv("SLACK_WEB_HOOK"), "", payload)
}

func storeToS3(message Message) {
	sess := session.Must(session.NewSession())
	uploader := s3.New(sess)
	data, _ := json.Marshal(message)

	if _, err := uploader.PutObject(&s3.PutObjectInput{
		Bucket:		aws.String(os.Getenv("S3_BACKET")),
		Key:			aws.String(message.Name),
		Body:			bytes.NewReader(data),
	}); err != nil {
		panic("x")
	}
}

func main() {
	lambda.Start(handler)
}
