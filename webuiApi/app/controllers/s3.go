package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"webuiApi/app/repositories/awsshared"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/olbrichattila/gofra/pkg/app/request"
)

const (
	uploadDirPath   = "./uploads/"
	downloadDirPath = "./downloads/"
)

type s3BucketRequest struct {
	Name string `json:"bucketName"`
}

type s3UploadRequest struct {
	BucketName string `json:"bucketName"`
	FileName   string `json:"fileName"`
}

// ListBucketsAction function can take any parameters defined in the Di config
func ListBucketsAction(awsShared awsshared.AWSShared) (string, error) {
	s3Client, ctx, err := awsShared.GetS3Client()
	if err != nil {
		return "", err
	}

	output, err := s3Client.ListBuckets(*ctx, &s3.ListBucketsInput{})
	if err != nil {
		return "", err
	}

	res, err := json.Marshal(output.Buckets)
	return string(res), err
}

// CreateBucketAction function can take any parameters defined in the Di config
func CreateBucketAction(awsShared awsshared.AWSShared, r request.Requester) (string, error) {
	var req s3BucketRequest
	if err := json.Unmarshal([]byte(r.Body()), &req); err != nil {
		return "", err
	}

	s3Client, ctx, err := awsShared.GetS3Client()
	if err != nil {
		return "", err
	}

	if _, err := s3Client.CreateBucket(*ctx, &s3.CreateBucketInput{
		Bucket: aws.String(req.Name),
	}); err != nil {
		return "", err
	}

	return ListBucketsAction(awsShared)
}

// DeleteBucketAction function can take any parameters defined in the Di config
func DeleteBucketAction(awsShared awsshared.AWSShared, r request.Requester) (string, error) {
	var req s3BucketRequest
	if err := json.Unmarshal([]byte(r.Body()), &req); err != nil {
		return "", err
	}

	s3Client, ctx, err := awsShared.GetS3Client()
	if err != nil {
		return "", err
	}

	if _, err := s3Client.DeleteBucket(*ctx, &s3.DeleteBucketInput{
		Bucket: aws.String(req.Name),
	}); err != nil {
		return "", err
	}

	return ListBucketsAction(awsShared)
}

// GetBucketContent function can take any parameters defined in the Di config
func GetBucketContent(r request.Requester, awsShared awsshared.AWSShared) (string, error) {
	bucketName := r.GetOne("bucketName", "")

	return getBucketContent(awsShared, bucketName)
}

func getBucketContent(awsShared awsshared.AWSShared, bucketName string) (string, error) {
	s3Client, ctx, err := awsShared.GetS3Client()
	if err != nil {
		return "", err
	}

	output, err := s3Client.ListObjects(*ctx, &s3.ListObjectsInput{
		Bucket: aws.String(bucketName),
	})

	if output.Contents == nil {
		return "[]", nil
	}

	res, err := json.Marshal(output.Contents)
	return string(res), err
}

// FileUpload function can take any parameters defined in the Di config
func FileUpload(req request.Requester, awsShared awsshared.AWSShared) (string, error) {
	r := req.GetRequest()
	err := r.ParseMultipartForm(10 << 20) // 10MB
	if err != nil {
		return "", err
	}

	// Get the file from form
	file, handler, err := r.FormFile("file")
	if err != nil {
		return "", err
	}

	defer file.Close()

	err = os.MkdirAll(uploadDirPath, 0755)
	if err != nil {
		return "", err
	}

	// Create a destination file
	dst, err := os.Create(uploadDirPath + handler.Filename)
	if err != nil {
		return "", err
	}

	defer dst.Close()

	// Copy the uploaded file to the destination
	_, err = io.Copy(dst, file)
	if err != nil {
		return "", err
	}

	return "[]", nil
}

// FileUploadToS3 function can take any parameters defined in the Di config
func FileUploadToS3(r request.Requester, awsShared awsshared.AWSShared) (string, error) {
	var req s3UploadRequest
	if err := json.Unmarshal([]byte(r.Body()), &req); err != nil {
		return "", err
	}

	s3Client, ctx, err := awsShared.GetS3Client()
	if err != nil {
		return "", err
	}

	file, err := os.Open(uploadDirPath + req.FileName)
	if err != nil {
		return "", err
	}

	defer func() {
		file.Close()
		os.Remove(uploadDirPath + req.FileName)
	}()

	_, err = s3Client.PutObject(*ctx, &s3.PutObjectInput{
		Bucket: aws.String(req.BucketName),
		Key:    aws.String(req.FileName),
		Body:   file,
	})
	if err != nil {
		return "", err
	}

	return getBucketContent(awsShared, req.BucketName)
}

// ViewFile function can take any parameters defined in the Di config
func ViewFile(w http.ResponseWriter, req request.Requester, awsShared awsshared.AWSShared) {
	r := req.GetRequest()
	bucketName := r.FormValue("bucketName")
	fileName := r.FormValue("fileName")

	s3Client, ctx, err := awsShared.GetS3Client()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get the object
	output, err := s3Client.GetObject(*ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer output.Body.Close()

	err = os.MkdirAll(downloadDirPath, 0755)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a local file
	filePath := downloadDirPath + fileName
	outFile, err := os.Create(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer outFile.Close()

	// Copy S3 object to file
	_, err = io.Copy(outFile, output.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "File not found.", http.StatusNotFound)
		return
	}

	defer func() {
		file.Close()
		os.Remove(filePath)
	}()

	// Get the filename from the path
	downloadFileName := filepath.Base(filePath)

	// Set headers to force download
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename=\""+downloadFileName+"\"")
	w.Header().Set("Content-Transfer-Encoding", "binary")

	// Optionally, set content length
	stat, _ := file.Stat()
	w.Header().Set("Content-Length", fmt.Sprintf("%d", stat.Size()))

	// Stream the file to the response
	io.Copy(w, file)
}

// DeleteFile function can take any parameters defined in the Di config
func DeleteFile(w http.ResponseWriter, r request.Requester, awsShared awsshared.AWSShared) (string, error) {
	var req s3UploadRequest
	if err := json.Unmarshal([]byte(r.Body()), &req); err != nil {
		return "", err
	}

	s3Client, ctx, err := awsShared.GetS3Client()
	if err != nil {
		return "", err
	}

	// Delete object
	_, err = s3Client.DeleteObject(*ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(req.BucketName),
		Key:    aws.String(req.FileName),
	})
	if err != nil {
		return "", err
	}

	return getBucketContent(awsShared, req.BucketName)
}
