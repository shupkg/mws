package mws

//SellerService 卖家服务
type SellerService struct {
	*Client
}

//Seller Seller
func Seller() *SellerService {
	return &SellerService{
		Client: newClient("/Sellers/2011-07-01", "2011-07-01"),
	}
}

//MarketplaceParticipationsResult MarketplaceParticipationsResult
type MarketplaceParticipationsResult struct {
	NextToken      string
	Participations []*Participation `xml:"ListParticipations>Participation"`
	Marketplaces   []*Marketplace   `xml:"ListMarketplaces>Marketplace"`
}

//Participation Participation
type Participation struct {
	MarketplaceID              string `xml:"MarketplaceId"`
	SellerID                   string `xml:"SellerId"`
	HasSellerSuspendedListings string `xml:"HasSellerSuspendedListings"`
}

//Marketplace Marketplace
type Marketplace struct {
	MarketplaceID       string `xml:"MarketplaceId"`
	Name                string
	DefaultCountryCode  string
	DefaultCurrencyCode string
	DefaultLanguageCode string
	DomainName          string
}

// ListMarketplaceParticipations Returns a list of marketplaces that the seller submitting the request can sell in, and a list of participations that include seller-specific information in that marketplace.
//
// The ListMarketplaceParticipations operation gets a list of marketplaces a seller can participate in and a list of participations that include seller-specific information in that marketplace. Note that the operation returns only those marketplaces where the seller's account is in an active state.
//
// The ListMarketplaceParticipations and ListMarketplaceParticipationsByNextToken operations together share a maximum request quota of 15 and a restore rate of one request per minute. For definitions of throttling terminology and for a complete explanation of throttling, see 限制：针对提交请求频率的限制 in the 亚马逊MWS开发者指南.
func (s *SellerService) ListMarketplaceParticipations(c *Credential, nextToken string) (string, *MarketplaceParticipationsResult, error) {
	data := ActionValues("ListMarketplaceParticipations")
	if nextToken != "" {
		data = ActionValues("ListMarketplaceParticipationsByNextToken")
		var result struct {
			BaseResponse
			Result *MarketplaceParticipationsResult `xml:"ListMarketplaceParticipationsByNextTokenResult"`
		}
		if err := s.GetModel(c, data, &result); err != nil {
			return "", nil, err
		}
		return result.RequestID, result.Result, nil
	}

	var result struct {
		BaseResponse
		Result *MarketplaceParticipationsResult `xml:"ListMarketplaceParticipationsResult"`
	}
	if err := s.GetModel(c, data, &result); err != nil {
		return "", nil, err
	}
	return result.RequestID, result.Result, nil
}
