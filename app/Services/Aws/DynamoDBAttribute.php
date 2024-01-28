<?php

namespace App\Services\Aws;
use App\Exceptions\Aws\DynamoDBServiceException;

class DynamoDBAttribute
{
    public function __construct(
        public readonly string $AttributeName,
        public readonly string $AttributeType,
        public readonly string $KeyType = '',
        public readonly string $Value = ''
    ) {
        if (!in_array($AttributeType, ['S','N','B'])) {
            throw new DynamoDBServiceException(['Attribute type ' . $AttributeType . 'does not match S or N or B!']);
        }

        if (!in_array($KeyType, ['HASH','RANGE'])) {
            throw new DynamoDBServiceException(['Key type ' . $KeyType . 'does not match HASH or RANGE!']);
        }
    }
}
