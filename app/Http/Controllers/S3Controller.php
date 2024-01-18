<?php

namespace App\Http\Controllers;

use App\Http\Requests\S3Request;
use App\Services\Aws\BucketService;
use Illuminate\Support\Facades\File;
use SebastianBergmann\Type\VoidType;
use App\Http\Requests\S3UploadRequest;
use App\Http\Requests\FileUploadRequest;

class S3Controller extends Controller
{

    public function __construct(private readonly BucketService $bucketService)
    {
    }

    public function buckets(): array
    {
        return $this->bucketService->buckets();
    }

    public function create(S3Request $request): array
    {
        return $this->bucketService->create(
            $request->validated()['bucketName']
        );
    }

    public function delete(S3Request $request): array
    {
        return $this->bucketService->delete(
            $request->validated()['bucketName']
        );
    }

    public function deleteMultiple(S3Request $request): array
    {
        return $this->bucketService->deleteMultiple(
            $request->validated()['bucketName']
        );
    }

    public function upload(S3UploadRequest $request): array
    {
        return $this->bucketService->upload(
            $request->validated()['bucketName'],
            $request->validated()['fileName'],
        );
    }

    public function list(string $bucketName): array
    {
        return $this->bucketService->list($bucketName);
    }

    public function load(S3UploadRequest $request): string
    {
        return $this->bucketService->load(
            $request->validated()['bucketName'],
            $request->validated()['fileName'],
        );
    }

    public function fileUpload(FileUploadRequest $request): void
    {
        $tmpFile = storage_path('app/' . $request->file('file')->store('tmp'));
        $fileName = $request->file('file')->getClientOriginalName();
        $storagePath = storage_path('app/' . $fileName);

        File::copy($tmpFile, $storagePath);
        File::delete($tmpFile);
    }
}
