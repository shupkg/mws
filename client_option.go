package mws

import (
	"log"
	"net/http"
	"os"
)

type Option func(*Client)

var defaultOptions []Option

func AddDefaultOption(options ...Option) {
	defaultOptions = append(defaultOptions, options...)
}

//HttpOpt 更改Http属性
func HttpOpt(opt func(client *http.Client)) Option {
	return func(client *Client) {
		opt(client.cli)
	}
}

//DebugOpt debug模式
func DebugOpt(debug bool) Option {
	return func(client *Client) {
		client.debug = debug
		if debug && client.log == nil {
			LogSTD()(client)
		}
	}
}

//ApiOption 引用Api名称和版本
func ApiOption(api, version string) Option {
	return func(c *Client) {
		c.Api = api
		c.Version = version
	}
}

//CredentialOption 授权信息
func CredentialOption(credential Credential) Option {
	return func(c *Client) {
		c.Credential = credential
	}
}

func LogOpt(printer Printer) Option {
	return func(client *Client) {
		client.log = printer
	}
}

func LogSTD() Option {
	return func(client *Client) {
		client.log = log.New(os.Stderr, "[MWS] ", log.LstdFlags)
	}
}
