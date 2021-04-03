package searching

import (
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/yofr4nk/lambda-face-collection/domain"
)

const confidencePercent float64 = 95

func SearchFaceByImage(searchParam domain.SearchParam) (domain.SearchResponse, error) {
	searchResponse := domain.SearchResponse{}

	input := rekognition.SearchFacesByImageInput{
		CollectionId: searchParam.CollectionId,
		Image: &rekognition.Image{
			S3Object: &rekognition.S3Object{
				Bucket: &searchParam.Bucket,
				Name:   &searchParam.Image,
			},
		},
	}

	searchByImage, err := searchParam.Rkc.SearchFacesByImage(&input)
	if err != nil {
		return searchResponse, err
	}

	for _, face := range searchByImage.FaceMatches {
		faceMatch := domain.FaceMatchInfo{
			Similarity:      *face.Similarity,
			Confidence:      *face.Face.Confidence,
			ExternalImageId: *face.Face.ExternalImageId,
			FaceId:          *face.Face.FaceId,
		}

		if faceMatch.Similarity >= confidencePercent {
			searchResponse.FaceMatches = append(searchResponse.FaceMatches, faceMatch)
		}
	}

	return searchResponse, nil

}
