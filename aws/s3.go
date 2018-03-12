package aws

import (
	"os"
	"bytes"
	"net/http"
	"log"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws/session"
)

const (
	S3_REGION = "us-east-1"
	S3_BUCKET = "apps-local"
)

func Upload() {
	//creds := credentials.NewEnvCredentials()
	creds := credentials.NewSharedCredentials("", "")
	_, err := creds.Get()
	if err != nil {
		log.Fatal("Bad credential ")
		log.Fatal(err)
	}

	cfg := aws.NewConfig().WithRegion(S3_REGION).WithCredentials(creds)
	s, err := session.NewSession(cfg)

	if err != nil {
		log.Fatal(err)
	}

	// Upload
	err = AddFileToS3(s, "/data/tmp.txt")
	if err != nil {
		log.Fatal(err)
	}
}

// AddFileToS3 will upload a single file to S3, it will require a pre-built aws session
// and will set file info like content type and encryption on the uploaded file.
func AddFileToS3(s *session.Session, fileDir string) error {

	// Open the file for use
	file, err := os.Open(fileDir)
	if err != nil {
		return err
	}
	defer file.Close()

	// Get file size and read the file content into a buffer
	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	// Config settings: this is where you choose the bucket, filename, content-type etc.
	// of the file you're uploading.
	_, err = s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(S3_BUCKET),
		Key:                  aws.String(fileDir),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(size),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})
	return err
}