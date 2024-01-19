<?php

namespace App\Contracts\Aws;


interface SqsServiceInterface
{
    public function list(): array;
    public function listWithAttributes(): array;
    public function create(string $name, int $delaySeconds, int $maximumMessageSize): array;
    public function delete(string $queueUrl): array;
    public function purge(string $queueUrl): array;
    public function attributes(string $queueUrl): array;
    public function sendMessage(string $queueUrl, int $delaySeconds, string $message): array;
    public function receiveMessage(string $queueUrl): array;
    public function receiveAndDeleteMessage(string $queueUrl): array;
    public function deleteMessage(string $queueUrl, string $receiptHandle): array;
    // @todo add deleteMessage, pugeQueue for now
}

