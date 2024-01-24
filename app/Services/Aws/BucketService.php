<?php

declare(strict_types=1);

namespace App\Services\Aws;

use Exception;
use Aws\S3\S3Client;
use App\Contracts\Aws\BucketServiceInterface;
use App\Exceptions\Aws\BucketServiceException;
use App\Contracts\Aws\AwsConfigServiceInterface;

class BucketService implements BucketServiceInterface
{
    private S3Client $s3Client;

    public function __construct(AwsConfigServiceInterface $config)
    {
        $this->s3Client = new S3Client($config->get());
    }

    public function create(string $name): array
    {
        try {
            $result = $this->s3Client->createBucket([
                'Bucket' => $name,
            ]);
        } catch (Exception $e) {
            throw new BucketServiceException([$e->getMessage()]);
        }

        return [
            'location' => $result['Location'] ?? null,
            'statusCode' => $result['@metadata']['statusCode'] ?? null,
        ];
    }

    public function buckets(): array
    {
        try {
            $result = $this->s3Client->listBuckets();
        } catch (Exception $e) {
            throw new BucketServiceException([$e->getMessage()]);
        }
        return $result['Buckets'] ?? throw new BucketServiceException(['Cannot load buckets']);
    }

    public function delete(string $name): array
    {

        try {
            $result = $this->s3Client->deleteBucket([
                'Bucket' => $name,
            ]);
        } catch (Exception $e) {
            throw new BucketServiceException([$e->getMessage()]);
        }

        return [
            'statusCode' => $result['@metadata']['statusCode'] ?? null,
        ];
    }

    public function deleteMultiple(string $name): array
    {
        $contents = $this->list($name);

        try {
            $objects = [];
            foreach ($contents as $content) {
                $objects[] = [
                    'Key' => $content['Key'],
                ];
            }
            $this->s3Client->deleteObjects([
                'Bucket' => $name,
                'Delete' => [
                    'Objects' => $objects,
                ],
            ]);
            $check = $this->s3Client->listObjects([
                'Bucket' => $name,
            ]);
            if (count($check) <= 0) {
                throw new BucketServiceException(["Bucket wasn't empty."]);
            }

        } catch (Exception $e) {
            throw new BucketServiceException([$e->getMessage()]);
        }

        return ['deleted' => $objects];
    }

    public function upload(string $bucketName, string $fileName): array
    {
        $filePath = storage_path('app/' . $fileName);
        try {
            $result = $this->s3Client->putObject([
                'Bucket' => $bucketName,
                'Key' => basename($fileName),
                'SourceFile' => $filePath
            ]);
        } catch (Exception $e) {
            throw new (['error' => $e->getMessage()]);
        }

        return [
            'statusCode' => $result['@metadata']['statusCode'] ?? null,
        ];
    }

    public function list(string $bucketName): array
    {
        try {
            $result = $this->s3Client->listObjects([
                'Bucket' => $bucketName,
            ]);
        } catch (Exception $e) {
            throw new BucketServiceException([$e->getMessage()]);
        }

        return (array) $result['Contents'] ?? throw new BucketServiceException(['Cannot list content of buckets']);
    }

    public function load(string $bucketName, string $fileName): string
    {
        try {
            $file = $this->s3Client->getObject([
                'Bucket' => $bucketName,
                'Key' => $fileName,
            ]);
        } catch (Exception $e) {
            throw new BucketServiceException([$e->getMessage()]);
        }

        return (string) $file->get('Body');
    }

    public function deleteObject(string $bucketName, string $fileName): array
    {
        try {
            $result = $this->s3Client->deleteObject(array(
                'Bucket' => $bucketName,
                'Key'    => $fileName
            ));
        } catch (Exception $e) {
            throw new BucketServiceException([$e->getMessage()]);
        }

        return (array) $result;
    }
}
