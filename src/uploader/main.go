package uploader

import (
	"backup_system/configs"
	"backup_system/log"
	"fmt"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var s3Client *s3.S3
var credentialsAWS configs.AWS

func init() {
	credentialsAWS = configs.ReturnAWSCredentials()
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(credentialsAWS.S3Region),
		Credentials: credentials.NewStaticCredentials(credentialsAWS.S3Key, credentialsAWS.S3Secret, ""),
	})
	if err != nil {
		panic(err)
	}
	s3Client = s3.New(sess)
}

func UploadFileToS3(filePath string, databaseName string) error {
	if !credentialsAWS.Active {
		log.Log("Upload para o S3 desativado")
		return nil
	}

	fileName := filepath.Base(filePath)

	log.Log(fmt.Sprint("Enviando arquivo para o S3:", fileName))
	// Abre o arquivo para upload
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Configura os parâmetros de upload
	params := &s3.PutObjectInput{
		Bucket: aws.String(credentialsAWS.S3Bucket),
		Key:    aws.String(databaseName + "/" + fileName),
		Body:   file,
	}

	// Faz o upload do arquivo
	_, err = s3Client.PutObject(params)
	if err != nil {
		return err
	}

	log.Log(fmt.Sprintf("Arquivo %s enviado com sucesso para o S3", fileName))
	return nil
}
