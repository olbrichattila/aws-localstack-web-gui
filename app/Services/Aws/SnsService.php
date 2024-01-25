<?php

namespace App\Services\Aws;

use Exception;
use Aws\Sns\SnsClient;
use App\Contracts\Aws\SnsServiceInterface;
use App\Exceptions\Aws\SnsServiceException;
use App\Contracts\Aws\AwsConfigServiceInterface;

class SnsService implements SnsServiceInterface
{
    private SnsClient $snsClient;

    public function __construct(AwsConfigServiceInterface $config)
    {
        $this->snsClient = new SnsClient($config->get());
    }

    public function createTopic(string $topicName): array
    {
        try {
            $result = $this->snsClient->createTopic([
                'Name' => $topicName,
            ]);

        } catch (Exception $e) {
            throw new SnsServiceException([$e->getMessage()]);
        }

        return (array) $result->get('TopicArn');
    }

    public function listTopics(): array
    {
        try {
            $result = $this->snsClient->listTopics([
            ]);
        } catch (Exception $e) {
            throw new SnsServiceException([$e->getMessage()]);
        }

        return (array) $result->get('Topics');
    }

    public function listTopicsWithAttributes(): array
    {
        $topics = $this->listTopics();

        return array_map(function (array $topic) {
            return $this->topicAttributes($topic['TopicArn']);
        }, $topics);
    }

    public function deleteTopic(string $topicName): array
    {
        try {
            $result = $this->snsClient->deleteTopic([
                'TopicArn' => $topicName,
            ]);

        } catch (Exception $e) {
            throw new SnsServiceException([$e->getMessage()]);
        }

        return (array) $result->get('@metadata');
    }

    public function topicAttributes(string $arnName): array
    {
        try {
            $result = $this->snsClient->getTopicAttributes([
                'TopicArn' => $arnName,
            ]);
        } catch (Exception $e) {
            throw new SnsServiceException([$e->getMessage()]);
        }

        return (array) $result->get('Attributes');
    }

    public function listSubscriptions(string $arnName): array
    {
        try {
            $result = $this->snsClient->listSubscriptions([
            ]);

        } catch (Exception $e) {
            throw new SnsServiceException([$e->getMessage()]);
        }

        return (array) $result->get('Subscriptions');
    }

    public function addHttpSubscription(string $arnName, $url): array
    {
        try {
            $result = $this->snsClient->subscribe([
                'Protocol' => $this->protocol($url),
                'Endpoint' => $url,
                'ReturnSubscriptionArn' => true,
                'TopicArn' => $arnName,
            ]);

        } catch (Exception $e) {
            throw new SnsServiceException([$e->getMessage()]);
        }

        return (array) $result->get('SubscriptionArn');
    }

    private function protocol(string $url): string
    {
        $urlParts = parse_url($url);

        // Check if the URL is valid
        if ($urlParts === false) {
            throw new SnsServiceException(['Invalid URl:' . $url]);
        }

        return $urlParts['scheme'] ?? throw new SnsServiceException(['Invalid URl:' . $url]);
    }

    public function publish(string $arnName, string $message): array
    {
        try {
            $result = $this->snsClient->publish([
                'Message' => $message,
                'TopicArn' => $arnName,
            ]);

        } catch (Exception $e) {
            throw new SnsServiceException([$e->getMessage()]);
        }

        return (array) $result->get('MessageId');
    }


    public function deleteSubscription(string $subscriptionArn): array
    {
        try {
            $result = $this->snsClient->unsubscribe([
                'SubscriptionArn' => $subscriptionArn,
            ]);
        } catch (Exception $e) {
            throw new SnsServiceException([$e->getMessage()]);  
        } 

        return (array) $result->get('@metadata');
    }
}
