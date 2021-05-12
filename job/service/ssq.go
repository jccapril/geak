package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"geak/job/model"
	"github.com/goinggo/mapstructure"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

//this.s.lotteryDao.RDB.Set(last_ssq_code_key,ssq.Code,ssq_expiration)
//this.lastSSQCode,_ = this.s.lotteryDao.RDB.Get(last_ssq_code_key).Result()
const (
	ssq_host          = "http://www.cwl.gov.cn/cwl_admin/kjxx/findDrawNotice?"
	last_ssq_code_key = "last_ssq_code_key"
	duration          = 60
	ssq_expiration	  = 4*24*3600*time.Second
)

func (this *Service) FetchSSQFromDBByCode(code string) {
	sqlStr := "SELECT `code`,`date`,`red`,`blue`,`blue2`,`sales`,`pool_money`,`first_count`," +
		"`first_money`,`second_count`,`second_money`,`third_count`,`third_money` FROM `ssq` WHERE `code`=?"
	ssq := new(model.SSQ)
	ssq.Name = "双色球"
	this.dao.DB.Get(ssq,sqlStr,"2021051")
}

func (this *Service) FetchSSQCountFromDBByCode(code string)(bool) {
	sqlStr := "SELECT COUNT(*) FROM `ssq` WHERE `code`=?"
	var num int
	err := this.dao.DB.Get(&num,sqlStr,code)
	if err != nil {
		log.Fatal(err)
	}
	return num > 0
}

func (this *Service) FetchLastSSQByRemote() {
	ticker := time.NewTicker(time.Second * duration)
	go func(t *time.Ticker) {
		for {
			select {
			case <-t.C:
				ssqlist, err := this.fetchSSQByRemote(1)
				if err != nil {
					fmt.Printf("双色球接口出问题了: %v\n", err)
				}else {
					if len(ssqlist) > 0 {
						ssq := ssqlist[0]
						isExist := this.FetchSSQCountFromDBByCode(ssq.Code)
						ssq.TransPrizegradesFmt()
						if isExist {
							fmt.Printf("ssq数据库里已经有了,update")
							sqlStr := "UPDATE `ssq` SET `date`=?,`red`=?,`blue`=?,`blue2`=?," +
								"`sales`=?,`pool_money`=?,`first_count`=?,`first_money`=?," +
								"`second_count`=?,`second_money`=?,`third_count`=?,`third_money`=? WHERE `code`=?"
							_,err := this.dao.DB.Exec(sqlStr,ssq.Date,ssq.Red,ssq.Blue,ssq.Blue2,ssq.Sales,
								ssq.PoolMoney,ssq.FirstCount,ssq.FirstMoney,ssq.SecondCount,ssq.SecondMoney,
								ssq.ThirdCount,ssq.ThirdMoney,ssq.Code)
							if err != nil {
								fmt.Println("ssq 数据更新失败：err",err)
							}
						}else {
							fmt.Printf("ssq数据库里没有该数据,insert")
							sqlStr := "INSERT INTO `ssq`(`code`, `date`, `red`, `blue`," +
								"`blue2`, `sales`, `pool_money`," +
								"`first_count`, `first_money`," +
								"`second_count`, `second_money`," +
								"`third_count`, `third_money`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
							_,err := this.dao.DB.Exec(sqlStr,ssq.Code,ssq.Date,ssq.Red,ssq.Blue,ssq.Blue2,ssq.Sales,
								ssq.PoolMoney,ssq.FirstCount,ssq.FirstMoney,ssq.SecondCount,ssq.SecondMoney,
								ssq.ThirdCount,ssq.ThirdMoney)
							if err != nil {
								fmt.Println("ssq 数据插入失败：err",err)
							}else {
								this.Push()
							}

							t.Stop()
						}
					}
				}

			}

		}
	}(ticker)

}

func (this *Service) fetchSSQByRemote(count int) (ssqList []model.SSQ, err error) {
	client := &http.Client{}
	q := url.Values{}
	q.Set("name", "ssq")
	q.Set("issueCount", strconv.Itoa(count))
	req, err := http.NewRequest(http.MethodGet, ssq_host+q.Encode(), nil)
	req.Header.Set("Referer", "http://www.cwl.gov.cn/")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return
	}
	resultList, isOK := result["result"].([]interface{})
	if !isOK {
		return ssqList, errors.New("format error")
	}
	err = mapstructure.Decode(resultList, &ssqList)
	return
}
