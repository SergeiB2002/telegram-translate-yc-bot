package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"strconv"
	"strings"
	// "mime/multipart"
)

const (
	vkCloudHotboxEndpoint = "https://hb.vkcs.cloud"
	defaultRegion         = "us-east-1"
)

func deleteFileFromBucket(filename string) {
	sess, err := session.NewSession()
	if err != nil {
		log.Fatalf("Unable to create session, %v", err)
	}
	svc := s3.New(sess, &aws.Config{Credentials: credentials.NewStaticCredentials("n", "dBq", ""), Region: aws.String(defaultRegion), Endpoint: aws.String(vkCloudHotboxEndpoint)})

	bucket := "bot_bucket"
	key := filename

	if _, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}); err != nil {
		log.Fatalf("Unable to delete object %q from bucket %q, %v\n", key, bucket, err)
		sendMessage("Ошибка удаления файла.")
	} else {
		log.Printf("Object %q deleted from bucket %q\n", key, bucket)
		sendMessage("Файл успешно удалён.")
	}

}

func uploadToBucket(update *TelegramMessageResponse, fileText string) {

	sess, _ := session.NewSession()

	svc := s3.New(sess, &aws.Config{Credentials: credentials.NewStaticCredentials("n", "dBq", ""), Region: aws.String(defaultRegion), Endpoint: aws.String(vkCloudHotboxEndpoint)})

	bucket := "bucket_raw_data"
	idStr := strconv.Itoa(user.Id)
	dateStr := strconv.Itoa(update.Message.Date)
	key := "default_name"
	switch flag {
	case 0:
		key = "detect_error_" + idStr + "_" + dateStr + ".txt"
	case 1:
		key = "detect_not_support_" + idStr + "_" + dateStr + ".txt"
	case 2:
		key = "detect_" + idStr + "_" + dateStr + ".txt"
	case 3:
		key = "translate_error_" + idStr + "_" + dateStr + ".txt"
	case 4:
		key = "translate_not_support_" + idStr + "_" + dateStr + ".txt"
	case 5:
		key = "translate_" + idStr + "_" + dateStr + ".txt"
	default:
		key = "error_" + idStr + "_" + dateStr + ".txt"
	}

	fmt.Println("СТАРТ ЗАГРУЗКИ В БАКЕТ")
	if _, err := svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   strings.NewReader(fileText),
	}); err != nil {
		log.Fatalf("Unable to upload %q to %q, %v\n", key, bucket, err)
	} else {
		fmt.Printf("File %q uploaded to bucket %q\n", key, bucket)
	}

}

func copyFileBetweenBuckets(filename string) {
	sess, _ := session.NewSession()

	svc := s3.New(sess, &aws.Config{Credentials: credentials.NewStaticCredentials("n", "dBq", ""), Region: aws.String(defaultRegion), Endpoint: aws.String(vkCloudHotboxEndpoint)})

	sourceBucket := "bucket_raw_data"
	sourceKey := filename
	destBucket := "bucket_processed_data"
	destKey := filename

	if _, err := svc.CopyObject(&s3.CopyObjectInput{
		Bucket:     aws.String(destBucket),
		Key:        aws.String(destKey),
		CopySource: aws.String(fmt.Sprintf("%s/%s", sourceBucket, sourceKey)),
	}); err != nil {
		fmt.Printf("Unable to copy object from %q to %q\n", sourceBucket, destBucket)
		sendMsg("Ошибка копирования файла.")
	} else {
		fmt.Printf("Object copied from %q to %q\n", sourceBucket, destBucket)
		sendMsg("Файл успешно скопирован.")
	}
}

func moveFileBetweenBuckets(filename string) {
	copyFileBetweenBuckets(filename)
	deleteFileFromBucket(filename)
}
