<?php

namespace App\Http\Requests;

use Illuminate\Foundation\Http\FormRequest;

class SqsSendMessageRequest extends FormRequest
{
    use S3RequestTrait;

    /**
     * Determine if the user is authorized to make this request.
     */
    public function authorize(): bool
    {
        return true;
    }

    /**
     * Get the validation rules that apply to the request.
     *
     * @return array<string, \Illuminate\Contracts\Validation\ValidationRule|array<mixed>|string>
     */
    public function rules(): array
    {
        return [
            'queueUrl' => 'required|string|max:255',
            'delaySeconds' => 'required|int',
            'messageBody' => 'required|string|max:4096',
        ];
    }
}