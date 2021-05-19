package mysql

import (
	"errors"
	"fmt"
	"geak/gkd/model"
	"geak/libs/conf"
	"geak/libs/factory"
	"geak/libs/log"
	"io/ioutil"
	"regexp"
	"strings"
)

const (
	ssqBodyExpr = "<tbody class=\"list-tr\">([\\s\\S]*?)</tbody>"
	ssqListExpr = "<tr[\\s\\S]*?>([\\s\\S]*?)</tr>"
	ssqElementExpr = "<td>([\\s\\S]*?)</td>"
	ssqRedBallExpr = "<span class=\"red\">([\\s\\S]*?)</span>"
	ssqBlueBallExpr = "<span class=\"blue\">([\\s\\S]*?)</span>"
)

type SSQClient struct {
	Year int
}


func (this *SSQClient) Produce()(dlt interface{},err error) {

	if this.Year > 2021 {
		return nil,errors.New("year > 2021")
	}
	ssqList,err := this.decodeFile(fmt.Sprintf("%s/ssq_history/%d.html",conf.Conf.App.Resources,this.Year))
	this.Year++
	return ssqList,err
}

func (this *SSQClient) Consume(ssq interface{})(result *factory.Result,err error) {
	data,isOK := ssq.([]*model.SSQ)
	if !isOK {
		return result,errors.New("ssq format error")
	}

	sqlStr := `INSERT INTO ssq(code, date, red, blue, 
                         blue2, sales, pool_money, 
                         first_count, first_money,
                         second_count, second_money,
 						 third_count, third_money) VALUES`
	vals := []interface{}{}
	for i := len(data)-1; i >=0; i-- {
		row := data[i]
		sqlStr += "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?),"
		vals = append(vals, row.Code, row.Date, row.Red, row.Blue, row.Blue2, row.Sales, row.PoolMoney,
			row.FirstCount, row.FirstMoney, row.SecondCount, row.SecondMoney, row.ThirdCount, row.ThirdMoney)
	}
	//trim the last ,
	sqlStr = strings.TrimSuffix(sqlStr, ",")
	//prepare the statement
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return
	}
	defer stmt.Close()
	//format all vals at once
	res, err := stmt.Exec(vals...)
	if err != nil {
		return
	}
	count, err := res.RowsAffected()
	if err != nil {
		return
	}
	log.Info(fmt.Sprintf("年份%v数据导入成功,一共导入%d条数据\n",data[0].Code,count))
	return
}

func (this *SSQClient) decodeFile(filePath string)(results []*model.SSQ,err error){
	var content []byte
	content, err = ioutil.ReadFile(filePath) // just pass the file name
	if err != nil {
		return
	}
	input := string(content)
	body := regexp.MustCompile(ssqBodyExpr).FindAllStringSubmatch(input,1)[0][1]
	ssqList := regexp.MustCompile(ssqListExpr).FindAllStringSubmatch(body,162)
	for _, ssq := range ssqList {
		val := new(model.SSQ)
		result := regexp.MustCompile(ssqElementExpr).FindAllStringSubmatch(ssq[1],12)
		if len(result) < 12 {
			continue
		}
		val.Code = result[0][1]
		val.Date = result[1][1]
		ballString := result[2][1]
		redList := regexp.MustCompile(ssqRedBallExpr).FindAllStringSubmatch(ballString,6)
		var reds []string

		for _, red := range redList {
			reds = append(reds, red[1])
		}
		val.Red = strings.Join(reds,",")

		blueList := regexp.MustCompile(ssqBlueBallExpr).FindAllStringSubmatch(ballString,2)
		for i, blue := range blueList {
			if i == 0 {
				val.Blue = blue[1]
			}else if(i == 1) {
				val.Blue2 = blue[1]
			}
		}
		val.Sales = strings.ReplaceAll(result[3][1],",","")
		val.PoolMoney = strings.ReplaceAll(result[4][1],",","")
		val.FirstCount = result[5][1]
		val.FirstMoney = strings.ReplaceAll(result[6][1],",","")
		val.SecondCount = result[7][1]
		val.SecondMoney = strings.ReplaceAll(result[8][1],",","")
		val.ThirdCount = result[9][1]
		val.ThirdMoney = strings.ReplaceAll(result[10][1],",","")
		results = append(results, val)
	}
	return
}



func createSSQTable()(err error) {
	sqlStr := `CREATE TABLE IF NOT EXISTS ssq(
        id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
        code VARCHAR(7) UNIQUE,
        date VARCHAR(20),
        red VARCHAR(20),
		blue VARCHAR(2),
		blue2 VARCHAR(2),
		sales VARCHAR(20),
		pool_money VARCHAR(20),
        first_count VARCHAR(10),
		first_money VARCHAR(20),
		second_count VARCHAR(10),	
		second_money VARCHAR(20),
		third_count VARCHAR(10),
		third_money VARCHAR(20),
		content	VARCHAR(100)
        )charset=utf8;`
	_,err = db.Exec(sqlStr)
	return
}



