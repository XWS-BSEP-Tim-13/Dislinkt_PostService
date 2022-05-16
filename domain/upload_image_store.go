package domain

import "github.com/aws/aws-sdk-go/service/s3"

type UploadImageStore interface {
	UploadObject(image []byte) (string, error)
	GetObject(filename string) []byte
	CreateBucket() (resp *s3.CreateBucketOutput)
}
