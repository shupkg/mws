package mws

const Version = "0.2.4"

func ApiOption(api, version string) Option {
	return OptionFunc(func(c *Client) {
		c.Api = api
		c.Version = version
	})
}

func CredentialOption(credential Credential) Option {
	return OptionFunc(func(c *Client) {
		c.Credential = credential
	})
}
