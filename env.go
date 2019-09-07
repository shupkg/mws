package mws

import (
	"fmt"
	"time"
)

//Env 环境方法
var Env env

func init() {
	//名称
	Env.nameMap = map[string]string{
		"ATVPDKIKX0DER":  "美国",   //美
		"A2EUQ1WTGCTBG2": "加拿大",  //加
		"A1AM78C64UM0Y8": "墨西哥",  //墨
		"A2Q3Y263D00KWC": "巴西",   //巴
		"A2VIGQ35RCS4UG": "阿联酋",  //阿
		"A1PA6795UKMFR9": "德国",   //德
		"A1RKKUPIHCS9HS": "西班牙",  //西
		"A13V1IB3VIYZZH": "法国",   //法
		"A1F83G8C2ARO7P": "英国",   //英
		"A21TJRUUN4KGV":  "印度",   //印
		"APJ6JRA9NG5V4":  "意大利",  //意
		"A33AVAJ2PDY3EV": "土耳其",  //土
		"A39IBJ37TRP1C6": "澳大利亚", //澳
		"A1VC38T7YXB528": "日本",   //日
		"AAHKV2X7AFYLW":  "中国",   //中
	}
	//域名
	Env.webDomainMap = map[string]string{
		"ATVPDKIKX0DER":  "amazon.com",    //美
		"A2EUQ1WTGCTBG2": "amazon.ca",     //加
		"A1AM78C64UM0Y8": "amazon.com.mx", //墨
		"A2Q3Y263D00KWC": "amazon.com.br", //巴
		"A2VIGQ35RCS4UG": "amazon.ae",     //阿
		"A1PA6795UKMFR9": "amazon.de",     //德
		"A1RKKUPIHCS9HS": "amazon.es",     //西
		"A13V1IB3VIYZZH": "amazon.fr",     //法
		"A1F83G8C2ARO7P": "amazon.co.uk",  //英
		"A21TJRUUN4KGV":  "amazon.in",     //印
		"APJ6JRA9NG5V4":  "amazon.it",     //意
		"A33AVAJ2PDY3EV": "amazon.com.tr", //土
		"A39IBJ37TRP1C6": "amazon.com.au", //澳
		"A1VC38T7YXB528": "amazon.co.jp",  //日
		"AAHKV2X7AFYLW":  "amazon.cn",     //中
	}
	//开发者平台域名
	Env.mwsDomainMap = map[string]string{
		"ATVPDKIKX0DER":  "mws.amazonservices.com",    //美
		"A2EUQ1WTGCTBG2": "mws.amazonservices.ca",     //加
		"A1AM78C64UM0Y8": "mws.amazonservices.com.mx", //墨
		"A2Q3Y263D00KWC": "mws.amazonservices.com",    //巴
		"A2VIGQ35RCS4UG": "mws-eu.amazonservices.com", //阿
		"A1PA6795UKMFR9": "mws-eu.amazonservices.com", //德
		"A1RKKUPIHCS9HS": "mws-eu.amazonservices.com", //西
		"A13V1IB3VIYZZH": "mws-eu.amazonservices.com", //法
		"A1F83G8C2ARO7P": "mws-eu.amazonservices.com", //英
		"APJ6JRA9NG5V4":  "mws-eu.amazonservices.com", //印
		"A21TJRUUN4KGV":  "mws.amazonservices.in",     //意
		"A33AVAJ2PDY3EV": "mws-eu.amazonservices.com", //土
		"A1VC38T7YXB528": "mws.amazonservices.jp",     //澳
		"A39IBJ37TRP1C6": "mws.amazonservices.com.au", //日
		"AAHKV2X7AFYLW":  "mws.amazonservices.com.cn", //中
	}
	//时区 INNA时区
	Env.tzMap = map[string]string{
		"ATVPDKIKX0DER":  "America/Los_Angeles", //美 洛杉矶
		"A2EUQ1WTGCTBG2": "America/Toronto",     //加 多伦多
		"A1AM78C64UM0Y8": "America/Mexico_City", //墨 墨西哥城
		"A2Q3Y263D00KWC": "America/Sao_Paulo",   //巴 圣保罗
		"A2VIGQ35RCS4UG": "Asia/Dubai",          //阿 迪拜
		"A1PA6795UKMFR9": "Europe/Berlin",       //德 柏林
		"A1RKKUPIHCS9HS": "Europe/Madrid",       //西 马德里
		"A13V1IB3VIYZZH": "Europe/Paris",        //法 巴黎
		"A1F83G8C2ARO7P": "Europe/London",       //英 伦敦
		"A21TJRUUN4KGV":  "Asia/Kolkata",        //印 加尔各答
		"APJ6JRA9NG5V4":  "Europe/Rome",         //意 罗马
		"A33AVAJ2PDY3EV": "Europe/Istanbul",     //土 伊斯坦布尔
		"A39IBJ37TRP1C6": "Australia/Sydney",    //澳 悉尼
		"A1VC38T7YXB528": "Asia/Tokyo",          //日 东京
		"AAHKV2X7AFYLW":  "Asia/Shanghai",       //中 上海
	}
	//国家分组
	Env.areaMap = map[string][]string{
		"NA": {"ATVPDKIKX0DER", "A2EUQ1WTGCTBG2", "A1AM78C64UM0Y8", "A2Q3Y263D00KWC"},                                                                        //南美:美加墨巴
		"EU": {"A2VIGQ35RCS4UG", "A1PA6795UKMFR9", "A1RKKUPIHCS9HS", "A13V1IB3VIYZZH", "A1F83G8C2ARO7P", "A21TJRUUN4KGV", "APJ6JRA9NG5V4", "A33AVAJ2PDY3EV"}, //欧洲:德西法英印意土
		"FE": {"A39IBJ37TRP1C6", "A1VC38T7YXB528"},                                                                                                           //远东:澳日
		"CN": {"AAHKV2X7AFYLW"},                                                                                                                              //中国:中
	}
	//国家关联
	Env.linkMap = map[string][]string{
		"ATVPDKIKX0DER":  Env.areaMap["NA"], //美 洛杉矶
		"A2EUQ1WTGCTBG2": Env.areaMap["NA"], //加 多伦多
		"A1AM78C64UM0Y8": Env.areaMap["NA"], //墨 墨西哥城
		"A2Q3Y263D00KWC": Env.areaMap["NA"], //巴 圣保罗
		"A2VIGQ35RCS4UG": Env.areaMap["EU"], //阿 迪拜
		"A1PA6795UKMFR9": Env.areaMap["EU"], //德 柏林
		"A1RKKUPIHCS9HS": Env.areaMap["EU"], //西 马德里
		"A13V1IB3VIYZZH": Env.areaMap["EU"], //法 巴黎
		"A1F83G8C2ARO7P": Env.areaMap["EU"], //英 伦敦
		"A21TJRUUN4KGV":  Env.areaMap["EU"], //印 加尔各答
		"APJ6JRA9NG5V4":  Env.areaMap["EU"], //意 罗马
		"A33AVAJ2PDY3EV": Env.areaMap["EU"], //土 伊斯坦布尔
		"A39IBJ37TRP1C6": Env.areaMap["FE"], //澳 悉尼
		"A1VC38T7YXB528": Env.areaMap["FE"], //日 东京
		"AAHKV2X7AFYLW":  Env.areaMap["CN"], //中 上海
	}
}

type env struct {
	nameMap      map[string]string
	webDomainMap map[string]string
	mwsDomainMap map[string]string
	tzMap        map[string]string
	linkMap      map[string][]string
	areaMap      map[string][]string
}

//GetTimezone 获取站点时区
func (e env) GetTimezone(marketplaceID string) *time.Location {
	innaName := e.tzMap[marketplaceID]
	if innaName != "" {
		l, err := time.LoadLocation(innaName)
		if err == nil {
			return l
		}
	}
	return time.UTC
}

//GetWebDomain 获取站点网站域名
func (e env) GetWebDomain(marketplaceID string) string {
	d := e.webDomainMap[marketplaceID]
	if d == "" {
		d = "amazon.com"
	}
	return d
}

//GetWebURL 获取站点域名前缀(含scheme)
func (e env) GetWebURL(marketplaceID string) string {
	return "https://www." + e.GetWebDomain(marketplaceID)
}

//GetSellercentralURL 获取站点登录域名(含scheme)
func (e env) GetSellercentralURL(marketplaceID string) string {
	return "https://sellercentral." + e.GetWebDomain(marketplaceID)
}

//GetMWSDomain 获取mws域名
func (e env) GetMWSDomain(marketplaceID string) string {
	d := e.mwsDomainMap[marketplaceID]
	if d == "" {
		d = "mws.amazonservices.com"
	}
	return d
}

//GetCountryName 获取站点国家名称
func (e env) GetCountryName(marketplaceID string) string {
	name := e.nameMap[marketplaceID]
	if name == "" {
		name = fmt.Sprintf("未知[%s]", marketplaceID)
	}
	return name
}

//GetLinkedMarketplace 获取关联的商城国家ID
func (e env) GetLinkedMarketplace(marketplaceID string) []string {
	marketplaces := e.linkMap[marketplaceID]
	if len(marketplaces) == 0 {
		marketplaces = []string{marketplaceID}
	}
	return marketplaces
}
