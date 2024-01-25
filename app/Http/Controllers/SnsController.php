<?php

namespace App\Http\Controllers;

use App\Services\Aws\SnsService;
use App\Http\Requests\SnsCreateTopicRequest;
use App\Http\Requests\SnsPublishMessageRequest;
use App\Http\Requests\SnsCreateSubscriptionRequest;

class SnsController extends Controller
{
    public function __construct(private readonly SnsService $snsService)
    {
    }

    public function listTopics(): array
    {
        return $this->snsService->listTopics();
    }

    public function listTopicsWithAttributes(): array
    {
        return $this->snsService->listTopicsWithAttributes();
    }

    public function createTopic(SnsCreateTopicRequest $request): array
    {
        return $this->snsService->createTopic(
            $request->validated()['name']
        );
    }

    public function deleteTopic(SnsCreateTopicRequest $request): array
    {
        return $this->snsService->deleteTopic(
            $request->validated()['name']
        );
    }

    public function topicAttributes(SnsCreateTopicRequest $request): array
    {
        return $this->snsService->topicAttributes(
            $request->validated()['name']
        );
    }

    public function listSubscriptions(string $arnName): array
    {
        return $this->snsService->listSubscriptions($arnName);
    }

    public function addHttpSubscription(string $arnName, SnsCreateSubscriptionRequest $request): array
    {
        return $this->snsService->addHttpSubscription($arnName, $request->validated('url'));
    }

    public function publish(string $arnName, SnsPublishMessageRequest $request): array
    {
        return $this->snsService->publish($arnName, $request->validated('message'));
    }

    public function deleteSubscription(string $subscriptionArnName): array
    {
        return $this->snsService->deleteSubscription($subscriptionArnName);
    }
}
