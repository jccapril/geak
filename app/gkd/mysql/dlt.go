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

	dltBodyExpr = "<tbody>([\\s\\S]*?)</tbody>"
	dltListExpr = "<tr[\\s\\S]*?>([\\s\\S]*?)</tr>"
	dltElementExpr = "<td[\\s\\S]*?>([\\s\\S]*?)</td>"
	dltRedBallExpr = "<a[\\s\\S]*?>([\\s\\S]*?)<i>"
	dltBlueBallExpr = "<i>([\\s\\S]*?)</i>"
)

func createDLTTable()(err error) {
	sqlStr := `CREATE TABLE IF NOT EXISTS dlt(
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

type DLTClient struct {
	Year int
}

func (this *DLTClient) Produce()(dlt interface{},err error) {

	if this.Year > 2021 {
		return nil,errors.New("year > 2021")
	}
	dltList,err := this.decodeFile(fmt.Sprintf("%s/dlt_history/%d.html",conf.Conf.App.Resources,this.Year))
	this.Year++
	return dltList,err
}

func (this *DLTClient) Consume(dlt interface{})(result *factory.Result,err error){
	dltList,isOK := dlt.([]*model.DLT)
	if !isOK {
		return result,errors.New("dlt format error")
	}
	sqlStr := `INSERT INTO dlt(code, date, red, blue, 
                         blue2, sales, pool_money, 
                         first_count, first_money,
                         second_count, second_money,
 						 third_count, third_money) VALUES`

	vals := []interface{}{}

	for i := len(dltList)-1; i >=0; i-- {
		row := dltList[i]
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
	log.Info(fmt.Sprintf("年份%v数据导入成功,一共导入%d条数据\n",dltList[0].Code,count))

	return
}

func (this *DLTClient) decodeFile(filePath string)(result []*model.DLT,err error){
	var content []byte
	content, err = ioutil.ReadFile(fmt.Sprintf("%s/dlt_history/%d.html",conf.Conf.App.Resources,this.Year)) // just pass the file name
	if err != nil {
		return
	}
	input := string(content)
	body := regexp.MustCompile(dltBodyExpr).FindAllStringSubmatch(input,2)[1][1]
	dltList := regexp.MustCompile(dltListExpr).FindAllStringSubmatch(body,162)

	for _, v := range dltList {
		dlt := new(model.DLT)
		element := regexp.MustCompile(dltElementExpr).FindAllStringSubmatch(v[1],10)
		if len(element) < 10 {
			continue
		}
		dlt.Code = element[0][1]
		dlt.Date = element[1][1]
		dlt.Sales = strings.ReplaceAll(element[2][1],",","")
		ballStr := element[3][1]
		redball := regexp.MustCompile(dltRedBallExpr).FindAllStringSubmatch(ballStr,1)[0][1]
		redball = strings.TrimSpace(redball)
		dlt.Red = strings.ReplaceAll(redball," ",",")
		blueBall := regexp.MustCompile(dltBlueBallExpr).FindAllStringSubmatch(ballStr,2)
		dlt.Blue = blueBall[0][1]
		dlt.Blue2 = blueBall[1][1]
		dlt.FirstCount = element[4][1]
		dlt.FirstMoney = element[5][1]
		dlt.SecondCount = element[6][1]
		dlt.SecondMoney = element[7][1]
		dlt.ThirdCount = element[8][1]
		dlt.ThirdMoney = element[9][1]
		result = append(result, dlt)
	}
	return
}

