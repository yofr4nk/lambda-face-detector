package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/yofr4nk/lambda-face-collection/domain"
	"github.com/yofr4nk/lambda-face-collection/searching"
	"github.com/yofr4nk/lambda-face-collection/sessions"
)

var collectionId = "face-container-list"

func HandleRequest(ctx context.Context, event events.S3Event) (domain.SearchResponse, error) {
	awsSession := sessions.CreateAwsSession()

	// Create a Rekognition client
	rkc := rekognition.New(awsSession)

	searchParam := domain.SearchParam{
		CollectionId: &collectionId,
		Rkc:          rkc,
		Bucket:       event.Records[0].S3.Bucket.Name,
		Image:        event.Records[0].S3.Object.Key,
	}

	return searching.AnalyzeFace(searchParam)
}

func main() {
	lambda.Start(HandleRequest)
}
