<?php

namespace App\Services\Aws;

use Illuminate\Http\JsonResponse;
use Illuminate\Http\Exceptions\HttpResponseException;

class BucketServiceException extends HttpResponseException
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
