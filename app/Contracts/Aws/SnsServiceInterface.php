<?php

namespace App\Contracts\Aws;

interface SnsServiceInterface
{

    public function createTopic(string $topicName): array;
    public function listTopics(): array;
    public function listTopicsWithAttributes(): array;
    public function deleteTopic(string $topicName): array;
    public function topicAttributes(string $arnName): array;
    public function listSubscriptions(string $arnName): array;
    public function addHttpSubscription(string $arnName, $url): array;
    public function publish(string $arnName, string $message): array;
    public function deleteSubscription(string $subscriptionArn): array;
} 
