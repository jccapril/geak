package lottery

import (
	"errors"
	"geak/api/dao"
	"geak/libs/conf"
	"geak/libs/log"
	"gitee.com/jlab/biz/m"
	"go.uber.org/zap"
	"strconv"
	"strings"
)

type Service struct {
	dao	*dao.Dao
	c   *conf.Config
}

func New(c *conf.Config) (s *Service) {

	s = &Service{
		dao:	dao.New(c),
		c:			c,
	}
	return s
}

func (this *Service)GetLastestSSQ()(ssq *m.Lottery,err error) {
	sqlStr := "SELECT `code`,`date`,`red`,`blue`,`blue2`,`sales`,`pool_money`,`first_count`," +
		"`first_money`,`second_count`,`second_money`,`third_count`,`third_money` FROM `ssq` ORDER BY `code` DESC LIMIT 0,1"
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

	var ssqList []*m.Lottery
	for rows.Next() {
		err = rows.Scan(&code, &date, &red, &blue, &blue2, &sales, &poolMoney, &firstCount,
			&firstMoney, &secondCount, &secondMoney, &thirdCount, &thirdMoney)
		if err != nil {
			log.Error("rows.Scan error",zap.Error(err))
			continue
		}
		ssq :=&m.Lottery{Name:"双色球",Type:1}
		ssq.Code = code
		ssq.Date = date
		ssq.Red = strings.Split(red,",")
		ssq.Blue = []string{blue,blue2}
		ssq.Sales,_ = strconv.ParseInt(sales,10,64)
		ssq.PoolMoney,_ = strconv.ParseInt(poolMoney,10,64)
		ssq.FirstCount,_ = strconv.ParseInt(firstCount,10,64)
		ssq.FirstMoney,_ = strconv.ParseInt(firstMoney,10,64)
		ssq.SecondCount,_ = strconv.ParseInt(secondCount,10,64)
		ssq.SecondMoney,_ = strconv.ParseInt(secondMoney,10,64)
		ssq.ThirdCount,_ = strconv.ParseInt(thirdCount,10,64)
		ssq.ThirdMoney,_ = strconv.ParseInt(thirdMoney,10,64)
		ssqList = append(ssqList,ssq)
	}
	err = rows.Err()
	if err != nil {
		log.Error("rows. error",zap.Error(err))
		return
	}
	if len(ssqList) == 1 {
		return ssqList[0],nil
	}else {
		return nil,errors.New("query rows len > 1")
	}

}
