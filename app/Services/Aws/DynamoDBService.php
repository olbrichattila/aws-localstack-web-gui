<?php

namespace App\Services\Aws;

use Exception;
use Aws\DynamoDb\DynamoDbClient;
use App\Contracts\Aws\DynamoDBServiceInterface;
use App\Contracts\Aws\AwsConfigServiceInterface;
use App\Exceptions\Aws\DynamoDBServiceException;

class DynamoDBService implements DynamoDBServiceInterface
{
    private DynamoDbClient $client;
    public function __construct(AwsConfigServiceInterface $config)
    {
        $this->client = new DynamoDbClient($config->get());
    }

    public function listTables(string $exclusiveStartTableName, int $limit): array
    {
        $options = [
            'Limit' => $limit,
        ];

        if ($exclusiveStartTableName !== '') {
            $options['ExclusiveStartTableName'] = $exclusiveStartTableName;
        }

        try {
            $result = $this->client->listTables($options);

        } catch (Exception $e) {
            throw new DynamoDBServiceException([$e->getMessage()]);
        }

        return (array) $result->get('TableNames');
    }

    public function createTable(string $tableName, array $attributes): array
    {

        $fields = array_map(function (array $attribute) {
            return new DynamoDBAttribute($attribute['attributeName'], $attribute['attributeType'], $attribute['keyType']);
        }, $attributes);

        return $this->_createTable(
            $tableName,
            $fields
        );

    }

    private function _createTable(string $tableName, array $attributes): array
    {
        try {
            $keySchema = [];
            $attributeDefinitions = [];
            foreach ($attributes as $attribute) {
                if (is_a($attribute, DynamoDBAttribute::class)) {
                    $keySchema[] = ['AttributeName' => $attribute->AttributeName, 'KeyType' => $attribute->KeyType];
                    $attributeDefinitions[] =
                        ['AttributeName' => $attribute->AttributeName, 'AttributeType' => $attribute->AttributeType];
                }
            }

            $result = $this->client->createTable([
                'TableName' => $tableName,
                'KeySchema' => $keySchema,
                'AttributeDefinitions' => $attributeDefinitions,
                'ProvisionedThroughput' => ['ReadCapacityUnits' => 10, 'WriteCapacityUnits' => 10],
            ]);
        } catch (Exception $e) {
            throw new DynamoDBServiceException([$e->getMessage()]);
        }

        return (array) $result->get('TableDescription');
    }

    public function deleteTable(string $tableName): array
    {
        try {
            $result = $this->client->deleteTable([
                'TableName' => $tableName,
            ]);
        } catch (Exception $e) {
            throw new DynamoDBServiceException([$e->getMessage()]);
        }

        return (array) $result;
    }

    
}
