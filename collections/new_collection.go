package collections

import (
	"github.com/aws/aws-sdk-go/service/rekognition"
	"log"
)

func NewCollection(rkc *rekognition.Rekognition, collectionId *string) bool {
	collectionInput := rekognition.CreateCollectionInput{
		CollectionId: collectionId,
	}

	collection, err := rkc.CreateCollection(&collectionInput)

	if err != nil {
		log.Print("something went wrong creating collection " + err.Error())

		return false
	}

	return *collection.StatusCode == 200

}
