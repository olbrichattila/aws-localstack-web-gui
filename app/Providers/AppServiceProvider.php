<?php

namespace App\Providers;

use App\Services\Aws\AwsConfigService;
use Illuminate\Support\ServiceProvider;
use App\Contracts\Aws\AwsConfigServiceInterface;

class AppServiceProvider extends ServiceProvider
{
    /**
     * Register any application services.
     */
    public function register(): void
    {
        $this->app->singleton(AwsConfigServiceInterface::class, AwsConfigService::class);
    }

    /**
     * Bootstrap any application services.
     */
    public function boot(): void
    {
        //
    }
}
