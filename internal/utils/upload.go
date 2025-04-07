package internal

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)


func UploadToS3(ctx context.Context, businessID int64, file []byte, fileName, prefix string) (string, error) {

	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	bucket := os.Getenv("AWS_S3_BUCKET")

	sess, err := session.NewSession(&aws.Config{
        Region:      aws.String("thailand"),
        Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
    })
    if err != nil {
        return "", err
    }

	// Create an S3 client
	s3Client := s3.New(sess)

	key := fmt.Sprintf("%d/%s/%s", businessID, prefix, fileName)

	_ , err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		Body:        bytes.NewReader(file),
		ContentType: aws.String("image/png"),
		ACL:         aws.String(string(types.ObjectCannedACLPublicRead)), // make it publicly accessible
	})

	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucket, key)
	return url, nil
}
