package service

import (
	"encoding/json"
	"errors"
	"geak/job/model"
	"geak/libs/log"
	"github.com/goinggo/mapstructure"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	ssq_host          = "http://www.cwl.gov.cn/cwl_admin/kjxx/findDrawNotice?"
	last_ssq_code_key = "last_ssq_code_key"
	ssq_duration          = 60
	ssq_expiration	  = 7*24*3600*time.Second
)

var fetchCount = 0

func (this *Service)GetSSQFromDBByCode(code string)(ssq *model.SSQ,err error) {
	sqlStr := "SELECT `code`,`date`,`red`,`blue`,`blue2`,`sales`,`pool_money`,`first_count`," +
		"`first_money`,`second_count`,`second_money`,`third_count`,`third_money` FROM `ssq` WHERE `code`=?"
	ssq = new(model.SSQ)
	err = this.dao.DB.Get(ssq,sqlStr,code)
	return
}

func (this *Service)GetSSQCountFromDBByCode(code string)(num int) {
	sqlStr := "SELECT COUNT(*) FROM `ssq` WHERE `code`=?"
	err := this.dao.DB.Get(&num,sqlStr,code)
	if err != nil {
		log.Error(sqlStr,zap.Error(err))
	}
	return
}

func (this *Service)GetLastestSSQ()(ssq *model.SSQ,isRemote bool) {

	lastestSSQCode,err := this.dao.RDB.Get(last_ssq_code_key).Result()
	if err != nil || len(lastestSSQCode) == 0 {
		// remote
		log.Error("redis get failure",zap.Error(err))
		ssq,err = this.GetLastestSSQByRemote()
		if err != nil {
			log.Error("ssq remote api",zap.Error(err))
		}
		isRemote = true
		return
	}
	// local
	ssq,err = this.GetSSQFromDBByCode(lastestSSQCode)
	if err != nil {
		log.Error("db query failure",zap.Error(err))
		ssq,err = this.GetLastestSSQByRemote()
		if err != nil {
			log.Error("ssq remote api",zap.Error(err))
		}
		isRemote = true
		return
	}
	return
}

func (this *Service)GetLastestSSQByRemote()(ssq *model.SSQ,err error){
	var ssqlist []model.SSQ
	ssqlist, err = this.fetchSSQByRemote(1)
	if err != nil {
		return
	}
	if len(ssqlist) == 0 {
		return nil,errors.New("query lastest ssq failure")
	}
	ssq = &ssqlist[0]
	this.dao.RDB.Set(last_ssq_code_key,ssq.Code,ssq_expiration)
	ssq.TransFormat()
	//this.waiter.Done()
	return
}

func (this *Service)StartSSQJob(){
	ticker := time.NewTicker(time.Second * ssq_duration)
	go func(t *time.Ticker) {
		for {
			select {
			case <-t.C:
				ssq,err := this.GetLastestSSQByRemote()
				if err != nil {
					log.Error("ssq remote",zap.Error(err))
				}else {
					isExist := this.GetSSQCountFromDBByCode(ssq.Code) > 0
					if isExist {
						sqlStr := "UPDATE `ssq` SET `date`=?,`red`=?,`blue`=?,`blue2`=?," +
							"`sales`=?,`pool_money`=?,`first_count`=?,`first_money`=?," +
							"`second_count`=?,`second_money`=?,`third_count`=?,`third_money`=? WHERE `code`=?"

						_,err := this.dao.DB.Exec(sqlStr,ssq.Date,ssq.Red,ssq.Blue,ssq.Blue2,ssq.Sales,
							ssq.PoolMoney,ssq.FirstCount,ssq.FirstMoney,ssq.SecondCount,ssq.SecondMoney,
							ssq.ThirdCount,ssq.ThirdMoney,ssq.Code)
						if err != nil {
							log.Error(sqlStr,zap.Error(err))
						}

					}else {
						sqlStr := "INSERT INTO `ssq`(`code`, `date`, `red`, `blue`," +
							"`blue2`, `sales`, `pool_money`," +
							"`first_count`, `first_money`," +
							"`second_count`, `second_money`," +
							"`third_count`, `third_money`, `content`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
						_,err := this.dao.DB.Exec(sqlStr,ssq.Code,ssq.Date,ssq.Red,ssq.Blue,ssq.Blue2,ssq.Sales,
							ssq.PoolMoney,ssq.FirstCount,ssq.FirstMoney,ssq.SecondCount,ssq.SecondMoney,
							ssq.ThirdCount,ssq.ThirdMoney,ssq.Content)
						if err != nil {
							log.Error(sqlStr,zap.Error(err))
						}else {
							this.Push()
						}

						t.Stop()
					}
					fetchCount+=1
					log.Info("ssq job",zap.Int("fetch count",fetchCount))
				}

			}
		}
	}(ticker)
}


func (this *Service)fetchSSQByRemote(count int) (ssqList []model.SSQ, err error) {
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


