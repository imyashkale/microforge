package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DyanamoDB struct {
	Client *dynamodb.DynamoDB
}

func New() *DyanamoDB {
	// Initialize a session that the SDK uses to load credentials
	// from the shared credentials file ~/.aws/credentials and the shared configuration file ~/.aws/config.
	sess, err := session.NewSession()

	if err != nil {
		panic(err)
	}

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	return &DyanamoDB{
		Client: svc,
	}
}
