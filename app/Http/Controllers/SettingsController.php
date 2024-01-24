<?php

namespace App\Http\Controllers;

use App\Http\Requests\SettingRequest;
use App\Contracts\Aws\AwsConfigServiceInterface;

class SettingsController extends Controller
{
    public function __construct(private readonly AwsConfigServiceInterface $configService)
    {
    }

    public function index(): array
    {
        return $this->configService->get();
    }

    public function store(SettingRequest $request): array
    {
        $validated = $request->validated();
        return $this->configService->store(
            $validated["region"],
            $validated["endpoint"],
            $validated["key"],
            $validated["secret"],
        );
    }
}
