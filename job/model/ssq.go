package model

type SSQ struct {
	// 名字
	Name string 		`json:"name"`
	// 期数
	Code string			`json:"code"`
	// 开奖日期
	Date string			`json:"date"`
	// 星期几
	Week string 		`json:"week"`
	// 红球
	Red  string			`json:"red"`
	// 蓝球
	Blue string			`json:"blue"`
	Blue2 string 		`json:"blue2"`
	// 本期销量
	Sales string 		`json:"sales"`
	// 累计奖池
	PoolMoney string 	`json:"poolMoney"`
	// 中奖详情简单
	Content string		`json:"content"`
	// 中奖详情详细
	Prizegrades []prizegrade	`json:"prizegrades"`
}

type prizegrade struct {
	Type int			`json:"type"`
	Typenum string		`json:"typenum"`
	Typemoney string	`json:"typemoney"`
}
