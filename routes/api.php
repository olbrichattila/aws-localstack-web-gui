<?php

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;
use  App\Http\Controllers\S3Controller;
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

Route::get('/s3/buckets', [S3Controller::class, 'buckets']);
Route::post('/s3/buckets', [S3Controller::class, 'create']);
Route::delete('/s3/buckets', [S3Controller::class, 'delete']);
Route::delete('/s3/buckets/multiple', [S3Controller::class, 'deleteMultiple']);
Route::post('/s3/buckets/upload', [S3Controller::class, 'upload']);
Route::get('/s3/list/{bucketName}', [S3Controller::class, 'list']);
Route::post('/s3/load', [S3Controller::class, 'load']);
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