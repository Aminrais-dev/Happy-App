package helper

import (
	"capstone/happyApp/config"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadFileToS3(DirName string, FileName string, Type string, fileData multipart.File) (string, error) {
	sess := config.GetSession()
	uploader := s3manager.NewUploader(sess)

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(os.Getenv("AWS_BUCKET")),
		Key:         aws.String("/" + DirName + "/" + FileName),
		Body:        fileData,
		ContentType: aws.String(Type),
	})

	if err != nil {
		log.Print(err.Error())
		return "", fmt.Errorf("Failed to upload file")
	}

	return result.Location, nil
}
func CheckFileType(filename string) (string, error) {
	extension := strings.ToLower(filename[strings.LastIndex(filename, ".")+1:])
	return extension, nil
}

func CheckFileExtension(filename string, contentType string) (string, error) {
	extension := strings.ToLower(filename[strings.LastIndex(filename, ".")+1:])

	if contentType == config.FileImageType {
		if extension != "jpg" && extension != "jpeg" && extension != "png" {
			return "", fmt.Errorf("Hanya menerima jpg, png atau jpeg")
		}
	}

	return extension, nil
}

func CheckFileSize(size int64, contentType string) error {
	if size == 0 {
		return fmt.Errorf("illegal file size")
	}
	if contentType == config.FileImageType {
		if size > 1097152 {
			return fmt.Errorf("file size too big")
		}
	}

	return nil
}
