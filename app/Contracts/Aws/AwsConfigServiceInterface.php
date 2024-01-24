<?php

namespace App\Contracts\Aws;

interface AwsConfigServiceInterface
{
    public function get(): array;
    public function store(string $region, string $enpoint, string $key, string $secret): array;

}
