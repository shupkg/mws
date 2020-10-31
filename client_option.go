package mws

type Option interface {
	Apply(*Client)
}

type OptionFunc func(*Client)

func (f OptionFunc) Apply(c *Client) {
	f(c)
}

var defaultOptions []Option

func AddOption(options ...Option) {
	defaultOptions = append(defaultOptions, options...)
}
