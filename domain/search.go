package domain

import "github.com/aws/aws-sdk-go/service/rekognition"

type SearchParam struct {
	CollectionId *string
	Rkc          *rekognition.Rekognition
	Bucket       string
	Image        string
}

type FaceMatchInfo struct {
	Similarity      float64 `json:"similarity"`
	Confidence      float64 `json:"confidence"`
	ExternalImageId string  `json:"externalImageId"`
	FaceId          string  `json:"faceId"`
}

type SearchResponse struct {
	FaceMatches []FaceMatchInfo `json:"faceMatches"`
}
