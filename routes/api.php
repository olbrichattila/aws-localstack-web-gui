<?php

use App\Http\Controllers\SettingsController;
use App\Http\Controllers\SnsController;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;
use App\Http\Controllers\S3Controller;
use App\Http\Controllers\SqsController;

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider and all of them will
| be assigned to the "api" middleware group. Make something great!
|
*/

Route::middleware('auth:sanctum')->get('/user', function (Request $request) {
    return $request->user();
});


Route::get('/settings', [SettingsController::class, 'index']);
Route::post('/settings', [SettingsController::class, 'store']);

Route::get('/s3/buckets', [S3Controller::class, 'buckets']);
Route::post('/s3/buckets', [S3Controller::class, 'create']);
Route::delete('/s3/buckets', [S3Controller::class, 'delete']);
Route::delete('/s3/buckets/multiple', [S3Controller::class, 'deleteMultiple']);
Route::post('/s3/buckets/upload', [S3Controller::class, 'upload']);
Route::get('/s3/list/{bucketName}', [S3Controller::class, 'list']);
Route::post('/s3/load', [S3Controller::class, 'load']);
Route::delete('/s3/buckets/delete/object', [S3Controller::class, 'deleteObject']);
Route::post('/s3/file_upload', [S3Controller::class, 'fileUpload'])->name('file_upload');

Route::get('/sqs', [SqsController::class, 'list']);
Route::get('/sqs/attributes', [SqsController::class, 'listWithAttributes']);
Route::post('/sqs', [SqsController::class, 'create']);
Route::delete('/sqs', [SqsController::class, 'delete']);
Route::delete('/sqs/purge', [SqsController::class, 'purge']);
Route::post('/sqs/attributes', [SqsController::class, 'attributes']);
Route::post('/sqs/message/send', [SqsController::class, 'sendMessage']);
Route::post('/sqs/message/receive', [SqsController::class, 'receiveMessage']);
Route::post('/sqs/message/receive/delete', [SqsController::class, 'receiveAndDeleteMessage']);
Route::post('/sqs/message/delete', [SqsController::class, 'deleteMessage']);


Route::get('/sns', [SnsController::class, 'listTopics']);
Route::get('/sns/attributes', [SnsController::class, 'listTopicsWithAttributes']);
Route::post('/sns', [SnsController::class, 'createTopic']);
Route::delete('/sns', [SnsController::class, 'deleteTopic']);
Route::post('/sns/attributes', [SnsController::class, 'topicAttributes']);

// subs
Route::get('/sns/sub/{arnName}', [SnsController::class, 'listSubscriptions']);
Route::post('/sns/sub/{arnName}', [SnsController::class, 'addHttpSubscription']);
Route::delete('/sns/sub/{subscriptionArnName}', [SnsController::class, 'deleteSubscription']);
Route::post('/sns/sub/{arnName}/publish', [SnsController::class, 'publish']);
