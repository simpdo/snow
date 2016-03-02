package main

type PayType struct {
	Name       string       `xml:"name"`
	Version    int          `xml:"version"`
	Valid      int          `xml:"valid"`
	Type       int          `xml:"type"`
	Isp        int          `xml:"ISP"`
	Priority   int          `xml:"priority"`
	Present    int          `xml:"present"`
	DayQuota   int          `xml:"day_quota"`
	MonthQuota int          `xml:"month_quota"`
	ExcludePf  string       `xml:"exclude_platform"`
	PayCodes   []XmlPayCode `xml:"pay_code"`
}

type PayTypeManager struct {
	Pays []PayType
}

type AppConfigManager struct {
	Configs []AppConfig
}

type AppConfig struct {
	Type int
	Apps []XmlAppConfig
}

type XmlPayTypeManager struct {
	Pays []XmlPayType `xml:"paytype"`
}

type XmlPayType struct {
	PayType
	Apps []XmlAppConfig `xml:"app"`
}

type XmlPayCode struct {
	Name  string         `xml:"name,attr"`
	Codes []XmlGoodsCode `xml:"code"`
}

type XmlGoodsCode struct {
	Goods int    `xml:"goods,attr"`
	Code  string `xml:",chardata"`
}

type XmlAppConfig struct {
	CodeName        string `xml:"code_name"`
	InvalidProvince string `xml:"invalid_province"`
	Channel         string `xml:"channel"`
	GoodsList       string `xml:"goods"`
	ExcludeVersion  string `xml:"exclude_version"`
	IncludeVersion  string `xml:"include_version"`
}

type JsonPayTypeManager struct {
}
