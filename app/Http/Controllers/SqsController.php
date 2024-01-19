<?php

namespace App\Http\Controllers;

use App\Services\Aws\SqsService;
use App\Http\Requests\SqsUrlRequest;
use App\Http\Requests\SqsCreateRequest;
use App\Http\Requests\SqsSendMessageRequest;
use App\Http\Requests\SqsDeleteMessageRequest;

class SqsController extends Controller
{
    public function __construct(private readonly SqsService $sqsService)
    {
    }

    public function list(): array
    {
        return $this->sqsService->list();
    }

    public function listWithAttributes(): array
    {
        return $this->sqsService->listWithAttributes();
    }
    
    public function create(SqsCreateRequest $request): array
    {
        return $this->sqsService->create(
            $request->validated()['name'],
            $request->validated()['delaySeconds'],
            $request->validated()['maximumMessageSize'],
        );
    }

    public function delete(SqsUrlRequest $request): array
    {
        return $this->sqsService->delete(
            $request->validated()['queueUrl']
        );
    }

    public function purge(SqsUrlRequest $request): array
    {
        return $this->sqsService->purge(
            $request->validated()['queueUrl']
        );
    }

    public function attributes(SqsUrlRequest $request): array
    {
        return $this->sqsService->attributes(
            $request->validated()['queueUrl']
        );
    }

    public function sendMessage(SqsSendMessageRequest $request): array
    {
        return $this->sqsService->sendMessage(
            $request->validated()['queueUrl'],
            $request->validated()['delaySeconds'],
            $request->validated()['messageBody'],
        );
    }

    public function receiveMessage(SqsUrlRequest $request): array
    {
        return $this->sqsService->receiveMessage(
            $request->validated()['queueUrl'],
        );
    }

    public function receiveAndDeleteMessage(SqsUrlRequest $request): array
    {
        return $this->sqsService->receiveAndDeleteMessage(
            $request->validated()['queueUrl'],
        );
    }

    public function deleteMessage(SqsDeleteMessageRequest $request): array
    {
        return $this->sqsService->deleteMessage(
            $request->validated()['queueUrl'],
            $request->validated()['receiptHandle'],
        );
    }
}
