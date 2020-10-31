package mws

import "context"

//SellerStats Seller
func SellerStats(credential Credential) *Client {
	return createClient(ApiOption("/SellerStats/2011-07-01", "2011-07-01"), CredentialOption(credential))
}

// ListMarketplaceParticipations Returns a list of marketplaces that the seller submitting the request can sell in, and a list of participations that include seller-specific information in that marketplace.
// The ListMarketplaceParticipations operation gets a list of marketplaces a seller can participate in and a list of participations that include seller-specific information in that marketplace. Note that the operation returns only those marketplaces where the seller's account is in an active state.
// The ListMarketplaceParticipations and ListMarketplaceParticipationsByNextToken operations together share a maximum request quota of 15 and a restore rate of one request per minute. For definitions of throttling terminology and for a complete explanation of throttling, see 限制：针对提交请求频率的限制 in the 亚马逊MWS开发者指南.
func (c *Client) ListMarketplaceParticipations(ctx context.Context, nextToken string) (result MarketplaceParticipationsResult, err error) {
	if nextToken != "" {
		var resp struct {
			ResponseMetadata
			Result MarketplaceParticipationsResult `xml:"ListMarketplaceParticipationsByNextTokenResult"`
		}
		err = c.getResult(ctx, "ListMarketplaceParticipationsByNextToken", ParamNexToken(nextToken), &resp)
		result = resp.Result
	} else {
		var resp struct {
			ResponseMetadata
			Result MarketplaceParticipationsResult `xml:"ListMarketplaceParticipationsResult"`
		}
		err = c.getResult(ctx, "ListMarketplaceParticipations", nil, &resp)
		result = resp.Result
	}
	return
}

//MarketplaceParticipationsResult MarketplaceParticipationsResult
type MarketplaceParticipationsResult struct {
	NextToken      string
	Participations []SellerParticipation     `xml:"ListParticipations>Participation"`
	Marketplaces   []SellerMarketplace `xml:"ListMarketplaces>Marketplace"`
}

//SellerParticipation Participation
type SellerParticipation struct {
	MarketplaceID              string `xml:"MarketplaceId"`
	SellerID                   string `xml:"SellerId"`
	HasSellerSuspendedListings string `xml:"HasSellerSuspendedListings"`
}

//SellerMarketplace Marketplace
type SellerMarketplace struct {
	MarketplaceID       string `xml:"MarketplaceId"`
	Name                string
	DefaultCountryCode  string
	DefaultCurrencyCode string
	DefaultLanguageCode string
	DomainName          string
}
