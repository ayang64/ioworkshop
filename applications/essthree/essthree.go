package essthree

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func Session() (*session.Session, error) {
	secret := struct {
		ID     string `json:"id"`
		Secret string `json:"secret"`
		Token  string `json:"token"`
	}{}

	readcreds := func() error {
		in, err := os.Open("/home/ayan/ioworkshop-aws-demo-secrets.json")
		if err != nil {
			return err
		}
		dc := json.NewDecoder(in)
		if err := dc.Decode(&secret); err != nil {
			return err
		}
		return nil
	}

	if err := readcreds(); err != nil {
		log.Print("this is a major error. i hope this isn't happening during a live demo, ayan!  get your act together.")
		log.Fatalf("error reading credentials file: %v", err)
		return nil, err
	}

	return session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(
			secret.ID,
			secret.Secret,
			secret.Token),
	})
}

// copy {{src}} to https://s3.amazonaws.com/ioworkshopdemo/{{dst}}
func CopyTraditional(src string, dst string) error {
	bucket := "ioworkshopdemo"
	sess, err := Session()

	in, err := os.Open(src)

	if err != nil {
		return err
	}

	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: &bucket,
		Key:    &dst,
		Body:   in,
	})

	if err != nil {
		return err
	}

	return nil
}

func Create(dst string) (io.Writer, error) {
	bucket := "ioworkshopdemo"
	sess, err := Session()
	if err != nil {
		return nil, err
	}

	pr, pw := io.Pipe()
	go func() {
		uploader := s3manager.NewUploader(sess)
		_, err := uploader.Upload(&s3manager.UploadInput{
			Bucket: &bucket,
			Key:    &dst,
			Body:   pr,
		})
		if err != nil {
			pr.CloseWithError(err)
		}
	}()
	return pw, nil
}
