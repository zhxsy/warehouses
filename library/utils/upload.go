package utils

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/cfx/warehouses/app"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func FileUploadToAws(file io.Reader, param map[string]string) (string, error) {
	tmpFile, err := ioutil.TempFile("", "awsTmp")
	if err != nil {
		app.Log().Error("create tmpFile error", err)
		return "", err
	}
	data, err := io.ReadAll(file)
	if err != nil {
		app.Log().Error("read file failed")
		return "", err
	}
	_, err = tmpFile.Write(data)
	if err != nil {
		app.Log().Error("tmp write file failed")
		return "", err
	}
	upload, err := AwsUpload(tmpFile.Name(), param)
	if err != nil {
		app.Log().Error("aws upload failed:", err)
		return "", err
	}
	err = tmpFile.Close()
	if err != nil {
		app.Log().Info("close tmpFile error")
		return "", err
	}
	defer os.Remove(tmpFile.Name())
	return upload, err
}

func AwsUploadAndRemove(path string, param map[string]string) (string, error) {
	// create an AWS session which can be
	s, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1"), // 账户的region
		Credentials: credentials.NewStaticCredentials(
			param["secretId"],  // secret-id
			param["secretKey"], // secret-key
			""),                // token can be left blank for now
	})
	if err != nil {
		return "", err
	}

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	url, err := UploadFileToS3(s, file, param)

	if err != nil {
		return "", err
	}
	file.Close()
	defer os.Remove(path)
	return url, nil
}

func AwsUpload(path string, param map[string]string) (string, error) {

	// create an AWS session which can be
	s, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1"), // 账户的region
		Credentials: credentials.NewStaticCredentials(
			param["secretId"],  // secret-id
			param["secretKey"], // secret-key
			""),                // token can be left blank for now
	})
	if err != nil {
		return "", err
	}

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	url, err := UploadFileToS3(s, file, param)

	if err != nil {
		return "", err
	}

	return url, nil
}

func UploadFileToS3(s *session.Session, file *os.File, param map[string]string) (string, error) {
	stat, _ := file.Stat()
	size := stat.Size()
	buffer := make([]byte, size)
	_, err := file.Read(buffer)

	// 此处文件名称即为相对文件名
	tempFileName := param["path"] + stat.Name() + filepath.Ext(stat.Name())
	_, err = s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(param["bucket"]), // bucket名称，把自己创建的bucket名称替换到此处即可
		Key:                  aws.String(tempFileName),
		ACL:                  aws.String(param["acl"]), // 权限枚举 https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl-overview.html#canned-acl
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(int64(size)),
		ContentDisposition:   aws.String(param["contentDisposition"]), // 预览
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ServerSideEncryption: aws.String("AES256"),
		StorageClass:         aws.String("INTELLIGENT_TIERING"),
	})
	if err != nil {
		return "", err
	}

	fileUrl := param["region"] + tempFileName

	return fileUrl, err
}
