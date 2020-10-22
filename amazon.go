package mws

import (
	"fmt"
	"strings"
)

var Amazon = newMap()

//国家信息
type Country struct {
	ID            string
	Label         string
	AreaID        string
	MarketplaceID string
	WebDomain     string
	ServiceDomain string
	Tz            Timezone
}

//时区信息
type Timezone struct {
	ID    string
	Label string
}

type countryMap struct {
	ids            []string
	marketplaces   []string
	areas          []string
	countries      map[string]Country
	marketplaceMap map[string]string
	areaMap        map[string][]string
}

func newMap() *countryMap {
	cMap := &countryMap{}
	cMap.ids = []string{"US", "CA", "MX", "BR", "GB", "FR", "DE", "IT", "ES", "IN", "AR", "TR", "AU", "JP", "CN"}
	cMap.areas = []string{"NA", "EU", "AU", "JP"}
	cMap.marketplaces = []string{
		"ATVPDKIKX0DER", "A2EUQ1WTGCTBG2", "A1AM78C64UM0Y8", "A2Q3Y263D00KWC", "A1F83G8C2ARO7P",
		"A13V1IB3VIYZZH", "A1PA6795UKMFR9", "APJ6JRA9NG5V4", "A1RKKUPIHCS9HS", "A21TJRUUN4KGV",
		"A2VIGQ35RCS4UG", "A33AVAJ2PDY3EV", "A39IBJ37TRP1C6", "A1VC38T7YXB528", "AAHKV2X7AFYLW",
	}
	for i := range cMap.ids {
		cMap.marketplaceMap[cMap.marketplaces[i]] = cMap.ids[i]
	}
	cMap.countries = map[string]Country{
		"US": {
			ID:            "US",
			Label:         "美国",
			AreaID:        "NA",
			MarketplaceID: "ATVPDKIKX0DER",
			WebDomain:     "amazon.com",
			ServiceDomain: "mws.amazonservices.com",
			Tz: Timezone{"America/Los_Angeles", "洛杉矶"},
		},
		"CA": {
			ID:            "CA",
			Label:         "加拿大",
			AreaID:        "NA",
			MarketplaceID: "A2EUQ1WTGCTBG2",
			WebDomain:     "amazon.ca",
			ServiceDomain: "mws.amazonservices.ca",
			Tz: Timezone{"America/Toronto", "多伦多"},
		},
		"MX": {
			ID:            "MX",
			Label:         "墨西哥",
			AreaID:        "NA",
			MarketplaceID: "A1AM78C64UM0Y8",
			WebDomain:     "amazon.com.mx",
			ServiceDomain: "mws.amazonservices.com.mx",
			Tz: Timezone{"America/Mexico_City", "墨西哥城"},
		},
		"BR": {
			ID:            "BR",
			Label:         "巴西",
			AreaID:        "NA",
			MarketplaceID: "A2Q3Y263D00KWC",
			WebDomain:     "amazon.com.br",
			ServiceDomain: "mws.amazonservices.com",
			Tz: Timezone{"America/Sao_Paulo", "圣保罗"},
		},
		"GB": {
			ID:            "GB",
			Label:         "英国",
			AreaID:        "EU",
			MarketplaceID: "A1F83G8C2ARO7P",
			WebDomain:     "amazon.co.uk",
			ServiceDomain: "mws-eu.amazonservices.com",
			Tz: Timezone{"Europe/London", "伦敦"},
		},
		"FR": {
			ID:            "FR",
			Label:         "法国",
			AreaID:        "EU",
			MarketplaceID: "A13V1IB3VIYZZH",
			WebDomain:     "amazon.fr",
			ServiceDomain: "mws-eu.amazonservices.com",
			Tz: Timezone{"Europe/Paris", "巴黎"},
		},
		"DE": {
			ID:            "DE",
			Label:         "德国",
			AreaID:        "EU",
			MarketplaceID: "A1PA6795UKMFR9",
			WebDomain:     "amazon.de",
			ServiceDomain: "mws-eu.amazonservices.com",
			Tz: Timezone{"Europe/Berlin", "柏林"},
		},
		"IT": {
			ID:            "IT",
			Label:         "意大利",
			AreaID:        "EU",
			MarketplaceID: "APJ6JRA9NG5V4",
			WebDomain:     "amazon.it",
			ServiceDomain: "mws-eu.amazonservices.com",
			Tz: Timezone{"Europe/Rome", "罗马"},
		},
		"ES": {
			ID:            "ES",
			Label:         "西班牙",
			AreaID:        "EU",
			MarketplaceID: "A1RKKUPIHCS9HS",
			WebDomain:     "amazon.es",
			ServiceDomain: "mws-eu.amazonservices.com",
			Tz: Timezone{"Europe/Madrid", "马德里"},
		},
		"IN": {
			ID:            "IN",
			Label:         "印度",
			AreaID:        "EU",
			MarketplaceID: "A21TJRUUN4KGV",
			WebDomain:     "amazon.in",
			ServiceDomain: "mws.amazonservices.in",
			Tz: Timezone{"Asia/Kolkata", "加尔各答"},
		},
		"AR": {
			ID:            "AR",
			Label:         "阿联酋",
			AreaID:        "EU",
			MarketplaceID: "A2VIGQ35RCS4UG",
			WebDomain:     "amazon.ae",
			ServiceDomain: "mws-eu.amazonservices.com",
			Tz: Timezone{"Asia/Dubai", "迪拜"},
		},
		"TR": {
			ID:            "TR",
			Label:         "土耳其",
			AreaID:        "EU",
			MarketplaceID: "A33AVAJ2PDY3EV",
			WebDomain:     "amazon.com.tr",
			ServiceDomain: "mws-eu.amazonservices.com",
			Tz: Timezone{"Europe/Istanbul", "伊斯坦布尔"},
		},
		"AU": {
			ID:            "AU",
			Label:         "澳大利亚",
			AreaID:        "AU",
			MarketplaceID: "A39IBJ37TRP1C6",
			WebDomain:     "amazon.com.au",
			ServiceDomain: "mws.amazonservices.com.au",
			Tz: Timezone{"Australia/Sydney", "悉尼"},
		},
		"JP": {
			ID:            "JP",
			Label:         "日本",
			AreaID:        "JP",
			MarketplaceID: "A1VC38T7YXB528",
			WebDomain:     "amazon.co.jp",
			ServiceDomain: "mws.amazonservices.jp",
			Tz: Timezone{"Asia/Tokyo", "东京"},
		},
		"CN": {
			ID:            "CN",
			Label:         "中国",
			AreaID:        "CN",
			MarketplaceID: "AAHKV2X7AFYLW",
			WebDomain:     "amazon.cn",
			ServiceDomain: "mws.amazonservices.com.cn",
			Tz: Timezone{"Asia/Shanghai", "上海"},
		},
	}
	cMap.areaMap = map[string][]string{
		"US": {"US", "CA", "MX", "BR"},
		"EU": {"BR", "GB", "FR", "DE", "IT", "ES", "IN", "AR", "TR"},
		"AU": {"AU"},
		"JP": {"JP"},
	}
	return cMap
}

//获取所有国家信息
func (m *countryMap) GetCountries() []Country {
	countries := make([]Country, len(m.ids))
	for i, id := range m.ids {
		countries[i] = m.countries[id]
	}
	return countries
}

//返回所有国家编号
func (m *countryMap) GetIDs() []string {
	return m.ids
}

//获取所有商城编号
func (m *countryMap) GetMarketplaces() []string {
	return m.marketplaces
}

//获取所有国家
func (m *countryMap) GetCountry(key string) Country {
	id := key
	if len(key) > 3 {
		id = m.marketplaceMap[key]
	}
	if id == "" {
		id = m.ids[0]
	}
	c := m.countries[id]
	if c.ID == "" {
		c = m.countries[m.ids[0]]
	}
	return c
}

//获取国家标签(中文名称)
func (m *countryMap) GetLabel(key string) string {
	return m.GetCountry(key).Label
}

//获取国家编号
func (m *countryMap) GetId(key string) string {
	return m.GetCountry(key).ID
}

//获取商城编号
func (m *countryMap) GetMarketplace(key string) string {
	return m.GetCountry(key).MarketplaceID
}

//获取域名
func (m *countryMap) GetMallHost(key string) string {
	return m.GetCountry(key).WebDomain
}

//拼接指定商城的地址
func (m *countryMap) GetMallURL(key string, path ...string) string {
	return fmt.Sprintf("https://www.%s/%s", m.GetMallHost(key), strings.Join(path, "/"))
}

//获取Asin商品链接
func (m *countryMap) GetAsinURL(key, asin string) string {
	if asin == "" {
		return ""
	}
	return m.GetMallURL(key, "dp", asin)
}

//获取商家后台地址
func (m *countryMap) GetSellerCenterURL(key string) string {
	return fmt.Sprintf("https://sellercenter.%s", m.GetMallHost(key))
}

//获取时区文字
func (m *countryMap) GetTimezone(key string) string {
	return m.GetCountry(key).Tz.ID
}

//获取时区标签
func (m *countryMap) GetTimezoneLabel(key string) string {
	return m.GetCountry(key).Tz.Label
}

//获取服务(MWS)域名
func (m *countryMap) GetServiceHost(key string) string {
	return m.GetCountry(key).ServiceDomain
}

//获取区域编号
func (m *countryMap) GetAreaID(key string) string {
	return m.GetCountry(key).AreaID
}

//获取所有区域编号
func (m *countryMap) GetAreas() []string {
	return m.areas
}

//获取所有区域标签
func (m *countryMap) GetAreaLabel(area string) string {
	area = strings.ToUpper(area)
	switch area {
	case "NA":
		return "北美"
	case "EU":
		return "欧洲"
	case "AU":
		return "澳洲"
	case "JP":
		return "日本"
	default:
		return area
	}
}
