# LocalStack Manager API Documentation

This document provides a comprehensive overview of all API endpoints used in the LocalStack Manager application. Use this as a reference when rewriting the backend from PHP/Laravel to Go.

## Base URL

All API requests are prefixed with: `http://localhost:8000/api`

## Settings API

### Get Settings
- **URL**: `/settings`
- **Method**: `GET`
- **Description**: Retrieves AWS configuration settings
- **Response**:
  ```json
  {
    "version": "latest",
    "region": "us-east-1",
    "endpoint": "http://localhost:4566",
    "use_path_style_endpoint": true,
    "credentials": {
      "key": "your-access-key-id",
      "secret": "your-secret-access-key"
    }
  }
  ```

### Save Settings
- **URL**: `/settings`
- **Method**: `POST`
- **Description**: Saves AWS configuration settings
- **Request Body**:
  ```json
  {
    "region": "us-east-1",
    "endpoint": "http://localhost:4566",
    "key": "your-access-key-id",
    "secret": "your-secret-access-key"
  }
  ```
- **Response**: Same as Get Settings response

## S3 API

### List Buckets
- **URL**: `/s3/buckets`
- **Method**: `GET`
- **Description**: Lists all S3 buckets
- **Response**:
  ```json
  [
    {
      "Name": "bucket-name-1",
      "CreationDate": "2023-01-01T00:00:00.000Z"
    },
    {
      "Name": "bucket-name-2",
      "CreationDate": "2023-01-02T00:00:00.000Z"
    }
  ]
  ```

### Create Bucket
- **URL**: `/s3/buckets`
- **Method**: `POST`
- **Description**: Creates a new S3 bucket
- **Request Body**:
  ```json
  {
    "bucketName": "new-bucket-name"
  }
  ```
- **Response**: Array of buckets (same as List Buckets)

### Delete Bucket
- **URL**: `/s3/buckets`
- **Method**: `DELETE`
- **Description**: Deletes an S3 bucket
- **Request Body**:
  ```json
  {
    "bucketName": "bucket-to-delete"
  }
  ```
- **Response**: Array of remaining buckets

### Delete Multiple Buckets
- **URL**: `/s3/buckets/multiple`
- **Method**: `DELETE`
- **Description**: Deletes multiple S3 buckets
- **Request Body**:
  ```json
  {
    "bucketName": ["bucket1", "bucket2"]
  }
  ```
- **Response**: Array of remaining buckets

### List Bucket Contents
- **URL**: `/s3/list/{bucketName}`
- **Method**: `GET`
- **Description**: Lists objects in a bucket
- **Response**:
  ```json
  [
    {
      "Key": "file1.txt",
      "LastModified": "2023-01-01T00:00:00.000Z",
      "Size": 1024,
      "StorageClass": "STANDARD"
    }
  ]
  ```

### Upload Object
- **URL**: `/s3/buckets/upload`
- **Method**: `POST`
- **Description**: Uploads an object to a bucket
- **Request Body**:
  ```json
  {
    "bucketName": "my-bucket",
    "fileName": "file.txt"
  }
  ```
- **Response**: Array of objects in the bucket

### Load Object
- **URL**: `/s3/load`
- **Method**: `POST`
- **Description**: Loads/downloads an object from a bucket
- **Request Body**:
  ```json
  {
    "bucketName": "my-bucket",
    "fileName": "file.txt"
  }
  ```
- **Response**: File content as string

### Delete Object
- **URL**: `/s3/buckets/delete/object`
- **Method**: `DELETE`
- **Description**: Deletes an object from a bucket
- **Request Body**:
  ```json
  {
    "bucketName": "my-bucket",
    "fileName": "file.txt"
  }
  ```
- **Response**: Array of remaining objects in the bucket

### File Upload
- **URL**: `/s3/file_upload`
- **Method**: `POST`
- **Description**: Uploads a file to temporary storage
- **Request Body**: Form data with file
- **Response**: No content

## SQS API

### List Queues with Attributes
- **URL**: `/sqs/attributes`
- **Method**: `GET`
- **Description**: Lists all SQS queues with their attributes
- **Response**:
  ```json
  [
    {
      "QueueUrl": "http://localhost:4566/000000000000/queue-name",
      "Attributes": {
        "ApproximateNumberOfMessages": "0",
        "ApproximateNumberOfMessagesNotVisible": "0",
        "ApproximateNumberOfMessagesDelayed": "0",
        "CreatedTimestamp": "1609459200",
        "LastModifiedTimestamp": "1609459200",
        "VisibilityTimeout": "30",
        "MaximumMessageSize": "262144",
        "MessageRetentionPeriod": "345600",
        "DelaySeconds": "0",
        "ReceiveMessageWaitTimeSeconds": "0"
      }
    }
  ]
  ```

### List Queues
- **URL**: `/sqs`
- **Method**: `GET`
- **Description**: Lists all SQS queues (without attributes)
- **Response**:
  ```json
  [
    {
      "QueueUrl": "http://localhost:4566/000000000000/queue-name"
    }
  ]
  ```

### Create Queue
- **URL**: `/sqs`
- **Method**: `POST`
- **Description**: Creates a new SQS queue
- **Request Body**:
  ```json
  {
    "name": "new-queue-name",
    "delaySeconds": 5,
    "maximumMessageSize": 4096
  }
  ```
- **Response**: Queue details

### Delete Queue
- **URL**: `/sqs`
- **Method**: `DELETE`
- **Description**: Deletes an SQS queue
- **Request Body**:
  ```json
  {
    "queueUrl": "http://localhost:4566/000000000000/queue-name"
  }
  ```
- **Response**: No content

### Purge Queue
- **URL**: `/sqs/purge`
- **Method**: `DELETE`
- **Description**: Purges all messages from an SQS queue
- **Request Body**:
  ```json
  {
    "queueUrl": "http://localhost:4566/000000000000/queue-name"
  }
  ```
- **Response**: No content

### Get Queue Attributes
- **URL**: `/sqs/attributes`
- **Method**: `POST`
- **Description**: Gets attributes for a specific queue
- **Request Body**:
  ```json
  {
    "queueUrl": "http://localhost:4566/000000000000/queue-name"
  }
  ```
- **Response**: Queue with attributes

### Send Message
- **URL**: `/sqs/message/send`
- **Method**: `POST`
- **Description**: Sends a message to an SQS queue
- **Request Body**:
  ```json
  {
    "queueUrl": "http://localhost:4566/000000000000/queue-name",
    "delaySeconds": 0,
    "messageBody": "Hello, world!"
  }
  ```
- **Response**: Message details including MessageId

### Receive Message
- **URL**: `/sqs/message/receive`
- **Method**: `POST`
- **Description**: Receives messages from an SQS queue
- **Request Body**:
  ```json
  {
    "queueUrl": "http://localhost:4566/000000000000/queue-name",
    "maxNumberOfMessages": 10
  }
  ```
- **Response**: Array of messages

### Receive and Delete Message
- **URL**: `/sqs/message/receive/delete`
- **Method**: `POST`
- **Description**: Receives and deletes messages from an SQS queue
- **Request Body**: Same as Receive Message
- **Response**: Array of messages

### Delete Message
- **URL**: `/sqs/message/delete`
- **Method**: `POST`
- **Description**: Deletes a message from an SQS queue
- **Request Body**:
  ```json
  {
    "queueUrl": "http://localhost:4566/000000000000/queue-name",
    "receiptHandle": "message-receipt-handle"
  }
  ```
- **Response**: No content

## SNS API

### List Topics with Attributes
- **URL**: `/sns/attributes`
- **Method**: `GET`
- **Description**: Lists all SNS topics with their attributes
- **Response**:
  ```json
  [
    {
      "TopicArn": "arn:aws:sns:us-east-1:000000000000:topic-name",
      "Attributes": {
        "DisplayName": "topic-name",
        "SubscriptionsConfirmed": "0",
        "SubscriptionsPending": "0",
        "SubscriptionsDeleted": "0"
      }
    }
  ]
  ```

### List Topics
- **URL**: `/sns`
- **Method**: `GET`
- **Description**: Lists all SNS topics
- **Response**:
  ```json
  [
    {
      "TopicArn": "arn:aws:sns:us-east-1:000000000000:topic-name"
    }
  ]
  ```

### Create Topic
- **URL**: `/sns`
- **Method**: `POST`
- **Description**: Creates a new SNS topic
- **Request Body**:
  ```json
  {
    "name": "new-topic-name"
  }
  ```
- **Response**: Topic details

### Delete Topic
- **URL**: `/sns`
- **Method**: `DELETE`
- **Description**: Deletes an SNS topic
- **Request Body**:
  ```json
  {
    "name": "arn:aws:sns:us-east-1:000000000000:topic-name"
  }
  ```
- **Response**: No content

### Get Topic Attributes
- **URL**: `/sns/attributes`
- **Method**: `POST`
- **Description**: Gets attributes for a specific topic
- **Request Body**:
  ```json
  {
    "name": "arn:aws:sns:us-east-1:000000000000:topic-name"
  }
  ```
- **Response**: Topic with attributes

### List Subscriptions
- **URL**: `/sns/sub/{topicArn}`
- **Method**: `GET`
- **Description**: Lists subscriptions for a topic
- **Response**:
  ```json
  [
    {
      "SubscriptionArn": "arn:aws:sns:us-east-1:000000000000:topic-name:subscription-id",
      "Owner": "000000000000",
      "Protocol": "http",
      "Endpoint": "http://example.com/webhook",
      "TopicArn": "arn:aws:sns:us-east-1:000000000000:topic-name"
    }
  ]
  ```

### Add HTTP Subscription
- **URL**: `/sns/sub/{topicArn}`
- **Method**: `POST`
- **Description**: Adds an HTTP subscription to a topic
- **Request Body**:
  ```json
  {
    "url": "http://example.com/webhook"
  }
  ```
- **Response**: Subscription details

### Delete Subscription
- **URL**: `/sns/sub/{subscriptionArnName}`
- **Method**: `DELETE`
- **Description**: Deletes a subscription
- **Response**: No content

### Publish Message
- **URL**: `/sns/sub/{topicArn}/publish`
- **Method**: `POST`
- **Description**: Publishes a message to a topic
- **Request Body**:
  ```json
  {
    "message": "Hello, world!"
  }
  ```
- **Response**: Publication details including MessageId

## DynamoDB API

### List Tables
- **URL**: `/dynamodb-list/{limit}/{exclusiveStartTableName?}`
- **Method**: `GET`
- **Description**: Lists DynamoDB tables
- **Parameters**:
  - `limit`: Maximum number of tables to return
  - `exclusiveStartTableName` (optional): Start listing from this table name
- **Response**:
  ```json
  {
    "TableNames": ["table1", "table2"],
    "LastEvaluatedTableName": "table2"
  }
  ```

### Create Table
- **URL**: `/dynamodb`
- **Method**: `POST`
- **Description**: Creates a new DynamoDB table
- **Request Body**:
  ```json
  {
    "TableName": "new-table",
    "KeySchema": [
      {
        "AttributeName": "id",
        "KeyType": "HASH"
      }
    ],
    "AttributeDefinitions": [
      {
        "AttributeName": "id",
        "AttributeType": "S"
      }
    ],
    "ProvisionedThroughput": {
      "ReadCapacityUnits": 5,
      "WriteCapacityUnits": 5
    }
  }
  ```
- **Response**: Table details

### Delete Table
- **URL**: `/dynamodb/{tableName}`
- **Method**: `DELETE`
- **Description**: Deletes a DynamoDB table
- **Response**: Deletion status