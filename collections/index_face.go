package collections

import (
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/yofr4nk/lambda-face-detector/domain"
	"strings"
)

func IndexFaceImage(param domain.SearchParam) (domain.SearchResponse, error) {
	var qualifyFilter = rekognition.QualityFilterAuto
	var detectionAttribute = "ALL"
	searchResponse := domain.SearchResponse{}

	indexFacesInput := rekognition.IndexFacesInput{
		CollectionId:        param.CollectionId,
		ExternalImageId:     &strings.Split(param.Image, "/")[1],
		DetectionAttributes: []*string{&detectionAttribute},
		Image: &rekognition.Image{
			S3Object: &rekognition.S3Object{
				Bucket: &param.Bucket,
				Name:   &param.Image,
			},
		},
		QualityFilter: &qualifyFilter,
	}

	indexFacesOutput, err := param.Rkc.IndexFaces(&indexFacesInput)
	if err != nil {
		return searchResponse, err
	}

	for _, face := range indexFacesOutput.FaceRecords {
		faceMatch := domain.FaceMatchInfo{
			Similarity:      0,
			Confidence:      *face.Face.Confidence,
			ExternalImageId: *face.Face.ExternalImageId,
			FaceId:          *face.Face.FaceId,
		}

		searchResponse.FaceMatches = append(searchResponse.FaceMatches, faceMatch)
	}

	return searchResponse, nil
}
