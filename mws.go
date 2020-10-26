package mws

const Version = "0.2.4"

func ApiOption(api, version string) ClientOption {
	return ClientOptionFunc(func(c *Client) {
		//"/Products/2011-10-01", "2011-10-01"
		c.api = api
		c.version = version
	})
}
