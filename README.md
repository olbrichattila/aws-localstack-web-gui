# Web GUI to manage AWS localstack services

## This is a work in progress ##

Services implemented:
    - S3
    - SQS

## Install
@todo

## Use pre-built docker image
```
docker pull aolb/localstack-web-ui:latest
docker run -p 80:80 aolb/localstack-web-ui:latest
```

## Visit:
```
http://localhost
```

## Docker compose example
```
version: "3"

services:
  web-ui:
    image: aolb/localstack-web-ui:latest
    container_name: localstack-web-ui
    ports:
      - "80:80"
    networks:
      - my-stack-dev

  localstack:
    image: localstack/localstack
    container_name: localstack-container
    ports:
      - "4566:4566"
      - "4510-4559:4510-4559"  # external services port range
    environment:
      - SERVICES=s3,sqs  # Include additional services as needed
      - DEBUG=1
      - AWS_ACCESS_KEY_ID=your-access-key-id
      - AWS_SECRET_ACCESS_KEY=your-secret-access-key
      - AWS_DEFAULT_REGION=us-east-1
    networks:
      - my-stack-dev

networks:
  my-stack-dev:
    driver: bridge
```

### If your localstack running on your local machine, use host network.

When you visit the page, it starts with settings, Please enter your region, key, secret and host.
In this example:
```
'region' => 'us-east-1',
'endpoint' => 'http://localstack-container:4566', // LocalStack endpoint visible within your container
'key' => 'your-access-key-id',
'secret' => 'your-secret-access-key',
            
```


@TODO description     

