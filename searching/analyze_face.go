package searching

import (
	"errors"
	"fmt"
	"github.com/yofr4nk/lambda-face-collection/collections"
	"github.com/yofr4nk/lambda-face-collection/domain"
)

func AnalyzeFace(searchParam domain.SearchParam) (domain.SearchResponse, error) {

	hasBeenConfigured := collections.SetCollection(searchParam)

	if hasBeenConfigured == false {
		return domain.SearchResponse{}, errors.New("could not create a new collection")
	}

	searchImageResponse, err := SearchFaceByImage(searchParam)
	if err != nil {
		fmt.Print("Error searching face" + err.Error())

		return searchImageResponse, err
	}

	if len(searchImageResponse.FaceMatches) == 0 {
		indexImageResponse, err := collections.IndexFaceImage(searchParam)
		if err != nil {
			fmt.Print("Error indexing face" + err.Error())

			return indexImageResponse, err
		}

		return indexImageResponse, nil

	}

	return searchImageResponse, nil
}
