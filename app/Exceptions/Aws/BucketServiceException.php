<?php

namespace App\Exceptions\Aws;

use Illuminate\Http\JsonResponse;
use Illuminate\Http\Exceptions\HttpResponseException;
use Throwable;

class BucketServiceException extends HttpResponseException implements Throwable
{
    public function __construct(array $errors)
    {
        parent::__construct(
            response()->json([
                'message' => 'S3 Bucket operation failed',
                'errors'  => $errors,
            ], JsonResponse::HTTP_UNPROCESSABLE_ENTITY)
        );
    }
}
