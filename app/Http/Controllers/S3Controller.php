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
    public function buckets(BucketService $bucketService): array
    {
        return $bucketService->buckets();
    }

    public function create(S3Request $request, BucketService $bucketService): array
    {
        return $bucketService->create(
            $request->validated()['bucketName']
        );
    }

    public function delete(S3Request $request, BucketService $bucketService): array
    {
        return $bucketService->delete(
            $request->validated()['bucketName']
        );
    }

    public function deleteMultiple(S3Request $request, BucketService $bucketService): array
    {
        return $bucketService->deleteMultiple(
            $request->validated()['bucketName']
        );
    }

    public function upload(S3UploadRequest $request, BucketService $bucketService): array
    {
        return $bucketService->upload(
            $request->validated()['bucketName'],
            $request->validated()['fileName'],
        );
    }

    public function list(string $bucketName, BucketService $bucketService): array
    {
        return $bucketService->list($bucketName);
    }

    public function load(S3UploadRequest $request, BucketService $bucketService): string
    {
        return $bucketService->load(
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
