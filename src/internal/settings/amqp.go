package settings

type AMQPSettings struct {
	URI                      string `yaml:"uri"`
	Queue                    string `yaml:"queue"`
	ConsumerTag              string `yaml:"consumer_tag"`
	ConsecutiveRequestsLimit int    `yaml:"consecutive-requests-limit"`
}
