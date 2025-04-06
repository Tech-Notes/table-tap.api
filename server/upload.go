package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func uploadToS3(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form (file upload)
	r.ParseMultipartForm(10 << 20) // 10 MB limit

	// Retrieve the file from the form
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error uploading file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	bucket := os.Getenv("AWS_S3_BUCKET")

	sess, err := session.NewSession(&aws.Config{
        Region:      aws.String("thailand"),
        Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
    })
    if err != nil {
        http.Error(w, "Error creating AWS session", http.StatusInternalServerError)
        return
    }

	// Create an S3 client
	s3Client := s3.New(sess)

	// Upload the file to the S3 bucket
	result, err := s3Client.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String("uploaded-image.jpg"), // Specify a unique name for the file
		Body:        file,                             // The file body
		ContentType: aws.String("image/jpeg"),         // Adjust based on the file type
		ACL:         aws.String("public-read"),        // Set ACL for public read access (optional)
	})
	if err != nil {
		http.Error(w, "Error uploading file to S3", http.StatusInternalServerError)
		return
	}

	// Respond with the S3 URL of the uploaded file
	fmt.Fprintf(w, "File uploaded successfully: %s", result.String())
}
