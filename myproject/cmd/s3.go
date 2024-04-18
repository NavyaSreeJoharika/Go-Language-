/*package s3

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/cobra"
)

// s3Cmd represents the s3 command
var s3Cmd = &cobra.Command{
	Use:   "s3",
	Short: "Interact with AWS S3",
	Run: func(cmd *cobra.Command, args []string) {
		// Load AWS configuration from the environment
		cfg, err := config.LoadDefaultConfig(context.Background())
		if err != nil {
			log.Fatalf("Failed to load AWS config: %v", err)
		}

		// Create an S3 client
		s3Client := s3.NewFromConfig(cfg)

		// List S3 buckets
		result, err := s3Client.ListBuckets(context.Background(), &s3.ListBucketsInput{})
		if err != nil {
			log.Fatalf("Failed to list S3 buckets: %v", err)
		}

		// Print bucket names
		fmt.Println("S3 Buckets:")
		for _, bucket := range result.Buckets {
			fmt.Println(*bucket.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(s3Cmd)
}
*/

package cmd
