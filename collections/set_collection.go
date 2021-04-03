package collections

import (
	"fmt"
	"github.com/yofr4nk/lambda-face-detector/domain"
)

func SetCollection(searchParam domain.SearchParam) bool {
	var hasBeenCreated bool

	collectionARN, err := GetCollection(searchParam.CollectionId, searchParam.Rkc)

	if err != nil || len(collectionARN) == 0 {
		hasBeenCreated = NewCollection(searchParam.Rkc, searchParam.CollectionId)
	}

	if hasBeenCreated == false && len(collectionARN) == 0 {
		fmt.Print("could not create a new collection")

		return false
	}

	return true
}
