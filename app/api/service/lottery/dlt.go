package lottery

import (
	"errors"
	"geak/biz/m"
	"geak/libs/log"
	"go.uber.org/zap"
	"strconv"
	"strings"
)

func (this *Service)GetLastestDLT()(dlt *m.Lottery,err error) {

	sqlStr := "SELECT `code`,`date`,`red`,`blue`,`blue2`,`sales`,`pool_money`,`first_count`," +
		"`first_money`,`second_count`,`second_money`,`third_count`,`third_money` FROM `dlt` ORDER BY `code` DESC LIMIT 0,1"
	var (
		code string
		date  string
		red	  string
		blue  string
		blue2 string
		sales string
		poolMoney string
		firstCount string
		firstMoney string
		secondCount string
		secondMoney string
		thirdCount string
		thirdMoney string
	)
	rows, err := this.dao.DB.Query(sqlStr)
	if err != nil {
		log.Error(sqlStr,zap.Error(err))
		return
	}
	defer rows.Close()

	var dltList []*m.Lottery
	for rows.Next() {
		err = rows.Scan(&code, &date, &red, &blue, &blue2, &sales, &poolMoney, &firstCount,
			&firstMoney, &secondCount, &secondMoney, &thirdCount, &thirdMoney)
		if err != nil {
			log.Error("rows.Scan error",zap.Error(err))
			continue
		}
		dlt :=&m.Lottery{Name: "大乐透",Type:1}
		dlt.Code = code
		dlt.Date = date
		dlt.Red = strings.Split(red,",")
		dlt.Blue = []string{blue,blue2}
		dlt.Sales,_ = strconv.ParseInt(sales,10,64)
		dlt.PoolMoney,_ = strconv.ParseInt(poolMoney,10,64)
		dlt.FirstCount,_ = strconv.ParseInt(firstCount,10,64)
		dlt.FirstMoney,_ = strconv.ParseInt(firstMoney,10,64)
		dlt.SecondCount,_ = strconv.ParseInt(secondCount,10,64)
		dlt.SecondMoney,_ = strconv.ParseInt(secondMoney,10,64)
		dlt.ThirdCount,_ = strconv.ParseInt(thirdCount,10,64)
		dlt.ThirdMoney,_ = strconv.ParseInt(thirdMoney,10,64)
		dltList = append(dltList,dlt)
	}
	err = rows.Err()
	if err != nil {
		log.Error("rows. error",zap.Error(err))
		return
	}
	if len(dltList) == 1 {
		return dltList[0],nil
	}else {
		return nil,errors.New("query rows len > 1")
	}


}
