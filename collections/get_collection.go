package collections

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

func GetCollection(collectionId *string, rkc *rekognition.Rekognition) (string, error) {
	collectionInput := rekognition.DescribeCollectionInput{
		CollectionId: collectionId,
	}

	collection, err := rkc.DescribeCollection(&collectionInput)

	if err != nil {
		fmt.Errorf("something went wrong describing collection %w", err)

		return "", err
	}

	return *collection.CollectionARN, nil

}
