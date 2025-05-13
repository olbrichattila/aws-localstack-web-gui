package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
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

// Custom AWS configuration
var (
	customAccessKey = "your-access-key-id"
	customSecretKey = "your-secret-access-key"
	customRegion    = "us-east-1"             // or your desired region
	customEndpoint  = "http://localhost:4566" // e.g., AWS: https://s3.amazonaws.com or MinIO URL
)

func (s *server) getS3BucketListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		s.newS3BucketHandler(w, r)
	}

	if r.Method == http.MethodDelete {
		s.s3DeleteHandler(w, r)
	}

	buckets, err := s.listS3Buckets()
	if err != nil {
		// TODO log error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(buckets)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *server) newS3BucketHandler(w http.ResponseWriter, r *http.Request) {
	var req s3BucketRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := s.createBucket(req.Name); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *server) s3DeleteHandler(w http.ResponseWriter, r *http.Request) {
	var req s3BucketRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := s.deleteBucket(req.Name); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *server) listS3Buckets() ([]types.Bucket, error) {
	s3Client, ctx, err := s.awsShared.GetS3Client()
	if err != nil {
		return nil, err
	}

	output, err := s3Client.ListBuckets(*ctx, &s3.ListBucketsInput{})
	if err != nil {
		return nil, err
	}

	return output.Buckets, nil
}

func (s *server) listS3BucketObjects(bucketName string) ([]types.Object, error) {
	s3Client, ctx, err := s.awsShared.GetS3Client()
	if err != nil {
		return nil, err
	}

	resp, err := s3Client.ListObjectsV2(*ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return nil, err
	}

	return resp.Contents, nil
}

func (s *server) createBucket(name string) error {
	s3Client, ctx, err := s.awsShared.GetS3Client()
	if err != nil {
		return err
	}

	_, err = s3Client.CreateBucket(*ctx, &s3.CreateBucketInput{
		Bucket: aws.String(name),
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *server) deleteBucket(name string) error {
	s3Client, ctx, err := s.awsShared.GetS3Client()
	if err != nil {
		return err
	}

	_, err = s3Client.DeleteBucket(*ctx, &s3.DeleteBucketInput{
		Bucket: aws.String(name),
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *server) getS3BucketContentListHandler(w http.ResponseWriter, r *http.Request) {
	bucketName := path.Base(r.URL.Path)
	s.s3BucketObjectList(w, r, bucketName)
}

func (s *server) s3BucketObjectList(w http.ResponseWriter, r *http.Request, bucketName string) {
	buckets, err := s.listS3BucketObjects(bucketName)
	if err != nil {
		// TODO log error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := []byte("[]")

	if buckets != nil {
		data, err = json.Marshal(buckets)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if _, err := w.Write(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *server) uploadToS3Handler(w http.ResponseWriter, r *http.Request) {
	// Upload to S3
	var req s3UploadRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s3Client, ctx, err := s.awsShared.GetS3Client()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	file, err := os.Open(uploadDirPath + req.FileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s.s3BucketObjectList(w, r, req.BucketName)
}

func (s *server) uploadToServerHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the file from form
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	err = os.MkdirAll(uploadDirPath, 0755)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a destination file
	dst, err := os.Create(uploadDirPath + handler.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy the uploaded file to the destination
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *server) deleteS3ObjectHandler(w http.ResponseWriter, r *http.Request) {
	var req s3UploadRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s3Client, ctx, err := s.awsShared.GetS3Client()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Delete object
	_, err = s3Client.DeleteObject(*ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(req.BucketName),
		Key:    aws.String(req.FileName),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	s.s3BucketObjectList(w, r, req.BucketName)
}

func (s *server) loads3ObjectHandler(w http.ResponseWriter, r *http.Request) {
	// Limit size to avoid large file uploads (optional)
	// r.ParseMultipartForm(10 << 20) // 10MB

	bucketName := r.FormValue("bucketName")
	fileName := r.FormValue("fileName")

	s3Client, ctx, err := s.awsShared.GetS3Client()
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
