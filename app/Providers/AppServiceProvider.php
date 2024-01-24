<?php

namespace App\Providers;

use App\Services\Aws\DbConfigService;
use Illuminate\Support\ServiceProvider;
use App\Contracts\Aws\AwsConfigServiceInterface;

class AppServiceProvider extends ServiceProvider
{
    /**
     * Register any application services.
     */
    public function register(): void
    {
        $this->app->singleton(AwsConfigServiceInterface::class, DbConfigService::class);
    }

    /**
     * Bootstrap any application services.
     */
    public function boot(): void
    {
        //
    }
}
