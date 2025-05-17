package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"webuiApi/app/repositories/awsshared"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/olbrichattila/gofra/pkg/app/gofraerror"
	"github.com/olbrichattila/gofra/pkg/app/request"
)

type S3Controller struct {
	s3Client *s3.Client
	ctx      *context.Context
}

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

func (s *S3Controller) Before(awsShared awsshared.AWSShared) {
	// TODO ADD error handling
	s.s3Client, s.ctx, _ = awsShared.GetS3Client()

}

// ListBucketsAction function can take any parameters defined in the Di config
func (s *S3Controller) ListBucketsAction() (string, error) {
	output, err := s.s3Client.ListBuckets(*s.ctx, &s3.ListBucketsInput{})
	if err != nil {
		return "", gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
	}

	res, err := json.Marshal(output.Buckets)
	if err != nil {
		return "", gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
	}

	return string(res), nil
}

// CreateBucketAction function can take any parameters defined in the Di config
func (s *S3Controller) CreateBucketAction(req s3BucketRequest) (string, error) {
	if _, err := s.s3Client.CreateBucket(*s.ctx, &s3.CreateBucketInput{
		Bucket: aws.String(req.Name),
	}); err != nil {
		return "", gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
	}

	return s.ListBucketsAction()
}

// DeleteBucketAction function can take any parameters defined in the Di config
func (s *S3Controller) DeleteBucketAction(req s3BucketRequest) (string, error) {
	if _, err := s.s3Client.DeleteBucket(*s.ctx, &s3.DeleteBucketInput{
		Bucket: aws.String(req.Name),
	}); err != nil {
		return "", gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
	}

	return s.ListBucketsAction()
}

// GetBucketContent function can take any parameters defined in the Di config
func (s *S3Controller) GetBucketContent(r request.Requester) (string, error) {
	bucketName := r.GetOne("bucketName", "")

	return s.getBucketContent(bucketName)
}

func (s *S3Controller) getBucketContent(bucketName string) (string, error) {
	output, err := s.s3Client.ListObjects(*s.ctx, &s3.ListObjectsInput{
		Bucket: aws.String(bucketName),
	})

	if output.Contents == nil {
		return "[]", nil
	}

	res, err := json.Marshal(output.Contents)
	if err != nil {
		return "", gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
	}

	return string(res), nil
}

// FileUpload function can take any parameters defined in the Di config
func (s *S3Controller) FileUpload(req request.Requester) (string, error) {
	r := req.GetRequest()
	err := r.ParseMultipartForm(10 << 20) // 10MB
	if err != nil {
		return "", gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
	}

	// Get the file from form
	file, handler, err := r.FormFile("file")
	if err != nil {
		return "", gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
	}

	defer file.Close()

	err = os.MkdirAll(uploadDirPath, 0755)
	if err != nil {
		return "", gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
	}

	// Create a destination file
	dst, err := os.Create(uploadDirPath + handler.Filename)
	if err != nil {
		return "", gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
	}

	defer dst.Close()

	// Copy the uploaded file to the destination
	_, err = io.Copy(dst, file)
	if err != nil {
		return "", gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
	}

	return "[]", nil
}

// FileUploadToS3 function can take any parameters defined in the Di config
func (s *S3Controller) FileUploadToS3(req s3UploadRequest) (string, error) {
	file, err := os.Open(uploadDirPath + req.FileName)
	if err != nil {
		return "", gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
	}

	defer func() {
		file.Close()
		os.Remove(uploadDirPath + req.FileName)
	}()

	_, err = s.s3Client.PutObject(*s.ctx, &s3.PutObjectInput{
		Bucket: aws.String(req.BucketName),
		Key:    aws.String(req.FileName),
		Body:   file,
	})
	if err != nil {
		return "", gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
	}

	return s.getBucketContent(req.BucketName)
}

// ViewFile function can take any parameters defined in the Di config
func (s *S3Controller) ViewFile(w http.ResponseWriter, req request.Requester) {
	r := req.GetRequest()
	bucketName := r.FormValue("bucketName")
	fileName := r.FormValue("fileName")

	// Get the object
	output, err := s.s3Client.GetObject(*s.ctx, &s3.GetObjectInput{
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
func (s *S3Controller) DeleteFile(req s3UploadRequest, w http.ResponseWriter) (string, error) {

	// Delete object
	_, err := s.s3Client.DeleteObject(*s.ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(req.BucketName),
		Key:    aws.String(req.FileName),
	})
	if err != nil {
		return "", gofraerror.NewJSON(err.Error(), http.StatusInternalServerError)
	}

	return s.getBucketContent(req.BucketName)
}
