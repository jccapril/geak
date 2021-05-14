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
	PoolMoney string 	`json:"poolMoney" db:"pool_money"`
	// 中奖详情简单
	Content string		`json:"content"`
	// 中奖详情详细
	Prizegrades []Prizegrade	`json:"prizegrades"`

	FirstCount 	string	`db:"first_count"`
	FirstMoney 	string	`db:"first_money"`
	SecondCount string	`db:"second_count"`
	SecondMoney	string	`db:"second_money"`
	ThirdCount 	string	`db:"third_count"`
	ThirdMoney 	string	`db:"third_money"`
}

type Prizegrade struct {
	Type int			`json:"type"`
	Typenum string		`json:"typenum"`
	Typemoney string	`json:"typemoney"`
}

func (this *SSQ)TransPrizegradesFmt(){
	if len(this.Prizegrades) > 0 {
		this.FirstCount = this.Prizegrades[0].Typenum
		this.FirstMoney = this.Prizegrades[0].Typemoney
		this.SecondCount = this.Prizegrades[1].Typenum
		this.SecondMoney = this.Prizegrades[1].Typemoney
		this.ThirdCount = this.Prizegrades[2].Typenum
		this.ThirdMoney = this.Prizegrades[2].Typemoney
	}
}
