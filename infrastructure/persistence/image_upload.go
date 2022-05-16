package persistence

import (
	"bytes"
	"fmt"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/domain"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"io/ioutil"
)

var (
	s3session *s3.S3
)

const (
	BUCKET_NAME = "dislinkt"
	REGION      = "eu-central-1"
)

type UploadImageStoreImpl struct {
}

func NewUploadImageStore() domain.UploadImageStore {
	return &UploadImageStoreImpl{}
}

func init() {
	s3session = s3.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String(REGION),
	})))
}

func (store *UploadImageStoreImpl) CreateBucket() (resp *s3.CreateBucketOutput) {
	resp, err := s3session.CreateBucket(&s3.CreateBucketInput{
		// ACL: aws.String(s3.BucketCannedACLPrivate),
		// ACL: aws.String(s3.BucketCannedACLPublicRead),
		Bucket: aws.String(BUCKET_NAME),
		CreateBucketConfiguration: &s3.CreateBucketConfiguration{
			LocationConstraint: aws.String(REGION),
		},
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeBucketAlreadyExists:
				fmt.Println("Bucket name already in use!")
				panic(err)
			case s3.ErrCodeBucketAlreadyOwnedByYou:
				fmt.Println("Bucket exists and is owned by you!")
			default:
				panic(err)
			}
		}
	}

	return resp
}

func (store *UploadImageStoreImpl) UploadObject(image []byte) (string, error) {
	filename := uuid.New()
	fmt.Println("Uploading:", filename)
	r := bytes.NewReader(image)
	_, err := s3session.PutObject(&s3.PutObjectInput{
		Body:   r,
		Bucket: aws.String(BUCKET_NAME),
		Key:    aws.String(filename.String()),
		ACL:    aws.String(s3.BucketCannedACLPublicRead),
	})

	if err != nil {
		return "", err
	}

	return filename.String(), nil
}

func (store *UploadImageStoreImpl) GetObject(filename string) []byte {
	fmt.Println("Downloading: ", filename)

	resp, err := s3session.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(BUCKET_NAME),
		Key:    aws.String(filename),
	})

	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	return body
}
