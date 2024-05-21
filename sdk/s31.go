package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	// Create a new AWS session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create an S3 client
	svc := s3.New(sess)

	// Specify the bucket you want to list objects from
	bucketName := "your-bucket-name"

	// Call the ListObjectsV2 method to list objects in the bucket
	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		fmt.Println("Error listing objects:", err)
		return
	}

	// Iterate through the objects and print their keys
	fmt.Println("Objects in bucket", bucketName+":")
	for _, obj := range resp.Contents {
		fmt.Println(*obj.Key)
	}
}
