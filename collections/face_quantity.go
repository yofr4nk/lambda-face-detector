package collections

import (
	"github.com/aws/aws-sdk-go/service/rekognition"
	"log"
)

func GetFaceQuantity(collectionId *string, rkc *rekognition.Rekognition) int64 {
	collectionInput := rekognition.DescribeCollectionInput{
		CollectionId: collectionId,
	}

	collection, err := rkc.DescribeCollection(&collectionInput)

	if err != nil {
		log.Print("something went wrong creating collection " + err.Error())

		return 0
	}

	return *collection.FaceCount

}
