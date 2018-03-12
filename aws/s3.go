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
	"github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds"
)

const (
	S3_REGION = "us-east-1"
	S3_BUCKET = "apps-local"
)

func Upload() {
	creds, err := GetCredential()
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
	err = AddFileToS3(s, "/data/tmp3.txt")
	if err != nil {
		log.Fatal(err)
	}
}

func GetCredential() (*credentials.Credentials, error) {
	provider := []credentials.Provider{&ec2rolecreds.EC2RoleProvider{},
		&credentials.SharedCredentialsProvider{Filename: "", Profile:  ""},
		&credentials.EnvProvider{},
		&credentials.StaticProvider{Value: credentials.Value{
			AccessKeyID:     "",
			SecretAccessKey: "",
			SessionToken:    "",
		}}}

	for _, p := range provider {
		_, err := getCredentialOfProvider(p)
		if err == nil {
			return credentials.NewCredentials(p), nil
		}
	}

	return nil, credentials.ErrNoValidProvidersFoundInChain
}

func getCredentialOfProvider(provider credentials.Provider) (v credentials.Value, err error){
	defer func(){
		if x := recover(); x != nil {
			log.Printf("run time panic: %v", x)
			err = credentials.ErrNoValidProvidersFoundInChain
		}
	}()
	return provider.Retrieve();
}

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