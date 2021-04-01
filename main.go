package lambda_face_collection

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/yofr4nk/lambda-face-collection/collections"
	"github.com/yofr4nk/lambda-face-collection/sessions"
)

var collectionId = "faces-container"

func HandleRequest(ctx context.Context, event events.S3Event) (string, error) {
	var hasBeenCreated bool

	fmt.Printf("This is your event: " + event.Records[0].EventName + "\n")

	awsSession := sessions.CreateAwsSession()

	// Create a Rekognition client
	rkc := rekognition.New(awsSession)

	quantity := collections.GetFaceQuantity(&collectionId, rkc)

	if quantity == 0 {
		hasBeenCreated = collections.NewCollection(rkc, &collectionId)
	}

	if hasBeenCreated == false {
		fmt.Print("could not create a new collection")

		return "", errors.New("could not create a new collection")
	}

	return fmt.Sprintf("Collection created!"), nil
}

func main() {
	lambda.Start(HandleRequest)
}
