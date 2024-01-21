<?php

namespace App\Services\Aws;


use Exception;
use Aws\Sqs\SqsClient;
use App\Services\Aws\AwsConfigService;
use App\Contracts\Aws\SqsServiceInterface;
use App\Exceptions\Aws\BucketServiceException;

class SqsService implements SqsServiceInterface
{
    private SqsClient $SqsClient;

    public function __construct(AwsConfigService $config) 
    {
        $this->SqsClient = new SqsClient($config->get());
    }

    public function list(): array
    {
        try {
            $result = $this->SqsClient->listQueues();
            
        } catch (Exception $e) {
            throw new BucketServiceException([$e->getMessage()]);
        }

        return $result->get('QueueUrls') ?: [];
    }

    public function listWithAttributes(): array
    {
        $urls = $this->list();
        return array_map(function(string $url) {
            return [
                'url' => $url,
                'attributes' => $this->attributes($url),
            ];
        }, $urls);
    }

    public function create(string $name, int $delaySeconds, int $maximumMessageSize): array
    {
        try {
            $result = $this->SqsClient->createQueue([
                'QueueName' => $name,
                'Attributes' => [
                    'DelaySeconds' => (string) $delaySeconds,
                    'MaximumMessageSize' => (string) $maximumMessageSize,
                ],
            ]);
        } catch (Exception $e) {
            throw new BucketServiceException([$e->getMessage()]);
        }

        return (array) $result;
    }

    public function delete(string $queueUrl): array
    {
        try {
            $result = $this->SqsClient->deleteQueue([
                'QueueUrl' => $queueUrl
            ]);
        } catch (Exception $e) {
            throw new BucketServiceException([$e->getMessage()]);
        }

        return (array) $result;
    }

    public function purge(string $queueUrl): array
    {
        try {
            $result = $this->SqsClient->purgeQueue([
                'QueueUrl' => $queueUrl
            ]);
        } catch (Exception $e) {
            throw new BucketServiceException([$e->getMessage()]);
        }

        return (array) $result;
    }

    public function attributes(string $queueUrl): array
    {
        try {
            $result = $this->SqsClient->getQueueAttributes([
                'AttributeNames' => ['All'],
                'QueueUrl' => $queueUrl,
            ]);
        } catch (Exception $e) {
            throw new BucketServiceException([$e->getMessage()]);
        }

        return $result['Attributes'] ?? throw new BucketServiceException(['Cannot retreive attributes of this queue at the moment.']);
    }


    public function sendMessage(string $queueUrl, int $delaySeconds, string $message): array
    {
        try {
            $result = $this->SqsClient->sendMessage([
                'DelaySeconds' => $delaySeconds,
                // 'MessageAttributes' => [
                //     '<String>' => [
                //         'BinaryListValues' => [<string || resource || Psr\Http\Message\StreamInterface>, ...],
                //         'BinaryValue' => <string || resource || Psr\Http\Message\StreamInterface>,
                //         'DataType' => '<string>', // REQUIRED
                //         'StringListValues' => ['<string>', ...],
                //         'StringValue' => '<string>',
                //     ],
                //     // ...
                // ],
                'MessageBody' => $message, // REQUIRED
                // 'MessageDeduplicationId' => '<string>',
                // 'MessageGroupId' => '<string>',
                // 'MessageSystemAttributes' => [
                //     '<MessageSystemAttributeNameForSends>' => [
                //         'BinaryListValues' => [<string || resource || Psr\Http\Message\StreamInterface>, ...],
                //         'BinaryValue' => <string || resource || Psr\Http\Message\StreamInterface>,
                //         'DataType' => '<string>', // REQUIRED
                //         'StringListValues' => ['<string>', ...],
                //         'StringValue' => '<string>',
                //     ],
                //     // ...
                // ],
                'QueueUrl' => $queueUrl, // REQUIRED
            ]);
        } catch (Exception $e) {
            throw new BucketServiceException([$e->getMessage()]);
        }

        return (array) $result;
    }

    public function receiveMessage(string $queueUrl, int $maxNumberOfMessages): array
    {
        try {
            $result = $this->SqsClient->receiveMessage([
                'QueueUrl' => $queueUrl,
                'MaxNumberOfMessages' => $maxNumberOfMessages,
                'VisibilityTimeout' =>  15, // @TODO may want to move to request
                'AttributeNames' => ['All']
            ]);
        } catch (Exception $e) {
            throw new BucketServiceException([$e->getMessage()]);
        }

        return (array) $result->get('Messages') ?? [];
    }


    public function receiveAndDeleteMessage(string $queueUrl): array
    {
        $messages =  $this->receiveMessage(
            $queueUrl,
        );

        foreach($messages as $message) {
            $receiptHandle = $message['ReceiptHandle'] ?? null;
            if ($receiptHandle) {
                $this->deleteMessage(
                    $queueUrl,
                    $receiptHandle,
                );
            }
        }

        return $messages;
    }

    public function deleteMessage(string $queueUrl, string $receiptHandle): array
    {
        try {
            $result = $this->SqsClient->deleteMessage([
                'QueueUrl' => $queueUrl,
                'ReceiptHandle' => $receiptHandle,
            ]);
        } catch (Exception $e) {
            throw new BucketServiceException([$e->getMessage()]);
        }

        return (array) $result;
    }
}
