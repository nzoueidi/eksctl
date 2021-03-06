package cloudformation

import (
	"encoding/json"
)

// AWSLambdaFunction_DeadLetterConfig AWS CloudFormation Resource (AWS::Lambda::Function.DeadLetterConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-deadletterconfig.html
type AWSLambdaFunction_DeadLetterConfig struct {

	// TargetArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-deadletterconfig.html#cfn-lambda-function-deadletterconfig-targetarn
	TargetArn *Value `json:"TargetArn,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSLambdaFunction_DeadLetterConfig) AWSCloudFormationType() string {
	return "AWS::Lambda::Function.DeadLetterConfig"
}

func (r *AWSLambdaFunction_DeadLetterConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
