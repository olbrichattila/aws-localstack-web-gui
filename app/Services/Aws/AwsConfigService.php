<?php

declare(strict_types= 1);

namespace App\Services\Aws;

use App\Contracts\Aws\AwsConfigServiceInterface;

class AwsConfigService implements AwsConfigServiceInterface
{
    public function get(): array
    {
        // @todo get from env, db whatever
        return [
            'version' => 'latest',
            'region' => 'us-east-1', // Set your desired AWS region
            'endpoint' => 'http://localhost:4566', // LocalStack endpoint
            // 'endpoint' => 'http://localstack-container:4566', // LocalStack endpoint
            'use_path_style_endpoint' => true,
            'credentials' => [
                'key' => 'your-access-key-id',
                'secret' => 'your-secret-access-key',
            ],
        ];
    }
}
