package alerting

// Config defines configuration for alerting
type Config struct {
	SlackToken    string `yaml:"slack_token"`
	EmailHost     string `yaml:"email_host"`
	EmailUserName string `yaml:"email_username"`
	EmailPassword string `yaml:"email_password"`
}
