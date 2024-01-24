<?php

declare(strict_types= 1);

namespace App\Services\Aws;

use App\Models\Settings;
use App\Contracts\Aws\AwsConfigServiceInterface;

class DbConfigService implements AwsConfigServiceInterface
{
    public function get(): array
    {
        $settings = Settings::first();
        if ($settings === null) {
            return [];
        }

        return [
            'version' => 'latest',
            'region' => $settings->region,
            'endpoint' => $settings->endpoint,
            'use_path_style_endpoint' => true,
            'credentials' => [
                'key' => $settings->key,
                'secret' => $settings->secret,
            ],
        ];
    }

    public function store(string $region, string $enpoint, string $key, string $secret): array
    {
         Settings::updateOrCreate([],
            [
                'region' => $region,
                'endpoint' => $enpoint,
                'key' => $key,
                'secret' => $secret,
            ]
        );

        return $this->get();
    }
}
