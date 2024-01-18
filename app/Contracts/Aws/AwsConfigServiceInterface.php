<?php

namespace App\Contracts\Aws;

interface AwsConfigServiceInterface
{
    public function get(): array;
}
