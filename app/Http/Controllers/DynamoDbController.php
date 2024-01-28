<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Services\Aws\DynamoDBService;
use App\Http\Requests\DynamoDbCreateTableRequest;

class DynamoDbController extends Controller
{
    public function __construct(private readonly DynamoDBService $dynamoDBService)
    {

    }

    public function listTables(int $limit, string $exclusiveStartTableName = ''): array
    {
        return $this->dynamoDBService->listTables($exclusiveStartTableName, $limit);
    }

    public function createTable(DynamoDbCreateTableRequest $request): array
    {
        return $this->dynamoDBService->createTable(
            $request->validated('name'),
            $request->validated('fields')
        );
    }

    public function deleteTable(string $tableName): array
    {
        return $this->dynamoDBService->deleteTable($tableName);
    }
}
