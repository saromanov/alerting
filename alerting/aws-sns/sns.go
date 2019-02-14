package sns

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/saromanov/alerting/alerting"
	"github.com/saromanov/alerting/structs"
)

type client struct {
	config *alerting.Alerting
}

// New provides initialization of aws client
func New(c *alerting.Alerting) client {
	return client{
		config: c,
	}
}

func (c *client) Send(m *structs.Message) (*structs.MessageResponse, error) {
	svc := sns.New(session.New(&aws.Config{
		Region: aws.String(c.config.AWSRegion),
	}))

	params := &sns.PublishInput{
		Message: aws.String(m.Text),
		MessageAttributes: map[string]*sns.MessageAttributeValue{
			"Key": {
				DataType:    aws.String("String"),
				StringValue: aws.String("String"),
			},
		},
		MessageStructure: aws.String("structure"),
		Subject:          aws.String(m.Subject),
		TopicArn:         aws.String(c.config.AWSTopicArn),
	}

	resp, err := svc.Publish(params)
	if err != nil {
		return nil, err
	}
	return &structs.MessageResponse{
		ID:       *resp.MessageId,
		Provider: c.String(),
	}, nil
}

// String returns name of provider
func (c *client) String() string {
	return "sns"
}
