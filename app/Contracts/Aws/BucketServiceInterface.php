<?php

namespace App\Contracts\Aws;

interface BucketServiceInterface
{
    public function create(string $name): array;
    public function buckets(): array;
    public function delete(string $name): array;
    public function deleteMultiple(string $name): array;
    public function upload(string $bucketName, string $fileName): array;
    public function list(string $bucketName): array;
    public function load(string $bucketName, string $fileName): string;
    public function deleteObject(string $bucketName, string $fileName): array;
}
