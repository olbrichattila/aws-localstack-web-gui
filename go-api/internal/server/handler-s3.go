package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
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
		fmt.Println(err)
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
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *server) listS3Buckets() ([]types.Bucket, error) {
	ctx := context.Background()
	s3Client, err := s.getS3Client(ctx)

	output, err := s3Client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		return nil, err
	}

	return output.Buckets, nil
}

func (s *server) listS3BucketObjects(bucketName string) ([]types.Object, error) {
	ctx := context.Background()
	s3Client, err := s.getS3Client(ctx)

	resp, err := s3Client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return nil, err
	}

	return resp.Contents, nil
}

func (s *server) createBucket(name string) error {
	ctx := context.Background()
	s3Client, err := s.getS3Client(ctx)

	_, err = s3Client.CreateBucket(ctx, &s3.CreateBucketInput{
		Bucket: aws.String(name),
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *server) deleteBucket(name string) error {
	ctx := context.Background()
	s3Client, err := s.getS3Client(ctx)

	_, err = s3Client.DeleteBucket(ctx, &s3.DeleteBucketInput{
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

	ctx := context.Background()
	s3Client, err := s.getS3Client(ctx)
	file, err := os.Open("./uploads/" + req.FileName)
	if err != nil {
		fmt.Println("failed to open file: %v", err)
	}
	defer file.Close()

	_, err = s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(req.BucketName),
		Key:    aws.String(req.FileName),
		Body:   file,
	})
	if err != nil {
		fmt.Println("failed to upload object: %v", err)
	}

	s.s3BucketObjectList(w, r, req.BucketName)
}

func (s *server) uploadToServerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Println("not post")
		return
	}

	fmt.Println("Content-Type:", r.Header.Get("Content-Type"))

	err := r.ParseMultipartForm(10 << 20) // 10MB
	if err != nil {
		fmt.Println("File size", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the file from form
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("File ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Create a destination file
	dst, err := os.Create("./uploads/" + handler.Filename)
	if err != nil {
		fmt.Println("crete ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy the uploaded file to the destination
	_, err = io.Copy(dst, file)
	if err != nil {
		fmt.Println("copy ", err)
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

	ctx := context.Background()
	s3Client, err := s.getS3Client(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Delete object
	_, err = s3Client.DeleteObject(ctx, &s3.DeleteObjectInput{
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

	// Get a regular form value
	bucketName := r.FormValue("bucketName")
	fmt.Printf("Name: %s\n", bucketName)

	fileName := r.FormValue("fileName")
	fmt.Printf("Name: %s\n", fileName)

	ctx := context.Background()
	s3Client, err := s.getS3Client(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the object
	output, err := s3Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
	})
	if err != nil {
		fmt.Printf("failed to get object: %v", err)
	}
	defer output.Body.Close()

	// Create a local file
	filePath := "./downloads/" + fileName
	outFile, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("failed to create file: %v", err)
	}
	defer outFile.Close()

	// Copy S3 object to file
	_, err = io.Copy(outFile, output.Body)
	if err != nil {
		fmt.Printf("failed to write object to file: %v", err)
	}

	fmt.Println("Object downloaded successfully")

	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "File not found.", http.StatusNotFound)
		return
	}
	defer file.Close()

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

func (s *server) getS3Client(ctx context.Context) (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(customRegion),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			customAccessKey, customSecretKey, "",
		)),
	)
	if err != nil {
		return nil, err
	}

	return s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(customEndpoint)
		o.UsePathStyle = true // Required for custom endpoints like MinIO
	}), nil
}
