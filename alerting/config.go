package alerting

// Config defines configuration for alerting
type Config struct {
	Providers []Alerting
}

// Alerting provides definition of provider
type Alerting struct {
	Name          string `yaml:"name"`
	SlackToken    string `yaml:"slack_token"`
	EmailHost     string `yaml:"email_host"`
	EmailUserName string `yaml:"email_username"`
	EmailPassword string `yaml:"email_password"`

	AWSClusterName string `yanl:"aws-cluster-name"`
	AWSEnabled     bool   `yanl:"aws-enabled"`
	AWSRegion      string `yanl:"aws-region"`
	AWSTopicArn    string `yanl:"aws-topic-arn"`
	AWSTemplate    string `yanl:"aws-template"`
}
