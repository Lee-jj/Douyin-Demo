package dao

import (
	"DOUYIN-DEMO/config"
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Minio struct {
	MinioClient *minio.Client
	endpoint    string
	videoBucket string
	imageBucket string
}

var client Minio

func GetMinio() Minio {
	return client
}

func InitMinio() {
	// using viper import yaml
	conf := config.GetConfig()
	endpoint := conf.Minio.Endpoint
	accessKeyID := conf.Minio.AccessKsyID
	secretAccessKey := conf.Minio.SecretAccessKey
	videoBucket := conf.Minio.VideoBucket
	imageBucket := conf.Minio.ImageBucket
	useSSL := false

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatal(err)
	}

	createBucket(minioClient, videoBucket)
	createBucket(minioClient, imageBucket)

	client = Minio{
		MinioClient: minioClient,
		endpoint:    endpoint,
		videoBucket: videoBucket,
		imageBucket: imageBucket,
	}
}

func createBucket(m *minio.Client, bucketName string) {
	found, err := m.BucketExists(context.Background(), bucketName)
	if err != nil {
		fmt.Printf("check %v bucketExists err: %s\n", bucketName, err.Error())
	}

	if !found {
		m.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
	}

	//设置桶策略
	policy := `{"Version": "2012-10-17",
				"Statement": 
					[{
						"Action":["s3:GetObject"],
						"Effect": "Allow",
						"Principal": {"AWS": ["*"]},
						"Resource": ["arn:aws:s3:::` + bucketName + `/*"],
						"Sid": ""
					}]
				}`
	err = m.SetBucketPolicy(context.Background(), bucketName, policy)
	if err != nil {
		fmt.Printf("SetBucketPolicy %s  err:%s\n", bucketName, err.Error())
	}
}

func (m *Minio) UpLoadFile(fileType, filePath, userID string) (string, error) {
	var fileName strings.Builder
	var contentType, suffix, bucketName string

	if fileType == "video" {
		contentType = "video/mp4"
		suffix = ".mp4"
		bucketName = m.videoBucket
	} else {
		contentType = "image/jpeg"
		suffix = ".jpg"
		bucketName = m.imageBucket
	}

	fileName.WriteString(userID)
	fileName.WriteString("_")
	timeNow := time.Now().Unix()
	fileName.WriteString(strconv.FormatInt(timeNow, 10))
	fileName.WriteString(suffix)

	_, err := m.MinioClient.FPutObject(context.Background(), bucketName, fileName.String(), filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return "", nil
	}

	url := "http://" + m.endpoint + "/" + bucketName + "/" + fileName.String()
	return url, nil
}
