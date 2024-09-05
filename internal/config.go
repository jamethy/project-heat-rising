package internal

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"strings"
)

// ReplaceSSMValues optionally replaces the string by reading it as an AWS parameter store address.
// Only effects values that start with ssm:
func ReplaceSSMValues(values ...*string) error {
	const ssmPrefix = "ssm:"
	ssmNames := make([]*string, 0)
	for _, value := range values {
		if strings.HasPrefix(*value, ssmPrefix) {
			name := strings.TrimPrefix(*value, ssmPrefix)
			ssmNames = append(ssmNames, &name)
		}
	}

	if len(ssmNames) == 0 {
		return nil
	}

	// if running locally, set AWS_PROFILE and AWS_SDK_LOAD_CONFIG=true
	awsSession, err := session.NewSession(&aws.Config{})
	if err != nil {
		return fmt.Errorf("failed to create aws session: %w", err)
	}
	ssmService := ssm.New(awsSession)
	req := &ssm.GetParametersInput{
		Names:          ssmNames,
		WithDecryption: aws.Bool(true),
	}
	res, err := ssmService.GetParameters(req)
	if err != nil {
		return fmt.Errorf("unable to get ssm parameters: %w", err)
	}
	for _, invalid := range res.InvalidParameters {
		fmt.Println("invalid ssm parameter given " + *invalid)
	}
	if len(res.InvalidParameters) > 0 {
		return errors.New("invalid ssm parameters (see logs)")
	}

	for _, value := range values {
		if !strings.HasPrefix(*value, ssmPrefix) {
			continue
		}
		for _, param := range res.Parameters {
			if ssmPrefix+*param.Name == *value {
				*value = *param.Value
				break
			}
		}
	}
	return nil
}
