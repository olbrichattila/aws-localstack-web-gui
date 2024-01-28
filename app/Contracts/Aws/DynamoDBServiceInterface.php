<?php

namespace App\Contracts\Aws;


interface DynamoDBServiceInterface
{
    public function listTables(string $prefix, int $limit): array;
    public function createTable(string $tableName, array $attributes): array;
    public function deleteTable(string $tableName): array;
}
