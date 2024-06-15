package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/aws/smithy-go/logging"
)

func KinesisClient() *kinesis.Client {
	conf, _ := config.LoadDefaultConfig(context.Background(), func(lo *config.LoadOptions) error {
		lo.Region = "ap-northeast-1"
		return nil
	})

	client := kinesis.NewFromConfig(conf, func(o *kinesis.Options) {
		o.Logger = logging.StandardLogger{}
	})
	return client
}
