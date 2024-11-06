package optionargs

import "errors"

type options struct {
	port    int
	content []string
}

type Option func(options *options) error

func WithPort(port int) Option {
	return func(options *options) error {
		if port < 0 {
			return errors.New("port should be positive")
		}
		options.port = port
		return nil
	}
}

func WithContent(content []string) Option {
	return func(options *options) error {
		options.content = content
		return nil
	}
}
