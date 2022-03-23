package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Response struct {
	Data string `json:"body"`
}

func Handler(request events.APIGatewayProxyRequest) (Response, error) {
	switch request.HTTPMethod {
	case "GET":
		return GetObjectData()
	default:
		panic("Unsupported method")
	}
}

func GetObjectData() (Response, error) {
	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}
	svc := s3.New(sess)
	bucket := "gidai-pastprob"
	resp, err := svc.ListObjects(&s3.ListObjectsInput{Bucket: aws.String(bucket)})
	if err != nil {
		panic(err)
	}
	data, err := json.Marshal(resp)
	if err != nil {
		panic(nil)
	}
	return Response {
		Data: string(data),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
