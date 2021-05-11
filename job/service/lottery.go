package service

import (
	"encoding/json"
	"errors"
	"geak/job/model"
	"github.com/robfig/cron"
	"net/url"
	"strconv"

	"time"

	//"time"
	"fmt"
	"github.com/goinggo/mapstructure"
	"io/ioutil"
	"net/http"
)

const (
	ssq_host          = "http://www.cwl.gov.cn/cwl_admin/kjxx/findDrawNotice?"
	last_ssq_code_key = "last_ssq_code_key"
	duration          = 60
	ssq_expiration	  = 4*24*3600*time.Second
)

type Lottery struct {
	s             *Service
	lastSSQCode   string
}

func (this *Lottery) Run() {
	this.runSSQ()
}

func (this *Lottery) runSSQ() {

	ticker := time.NewTicker(time.Second * duration)
	this.lastSSQCode,_ = this.s.lotteryDao.RDB.Get(last_ssq_code_key).Result()
	go func(t *time.Ticker) {
		for {
			select {
			case <-t.C:
				ssq, err := this.GetLatestSSQByRemote()
				if err != nil {
					fmt.Println("error:", err)
				} else {
					if this.lastSSQCode != ssq.Code {
						this.lastSSQCode = ssq.Code
						this.s.lotteryDao.RDB.Set(last_ssq_code_key,ssq.Code,ssq_expiration)
						fmt.Println(ssq.Code)
						fmt.Println("数据刷新了")
						ticker.Stop()
					}else {
						fmt.Println(ssq.Code)
						fmt.Println("数据还没有刷新")
					}

				}
			}
		}
	}(ticker)

}

func (this *Lottery) Register(s *Service) {
	this.s = s

	c := cron.New()
	defer c.Stop()
	//spec := "0 */1 12 * * MON,WED,SAT"
	spec := "0 0 21 * * TUE,THU,SUN"
	c.AddJob(spec, this)
	c.Start()
	select {}
}

func (this *Lottery) GetLatestSSQByRemote() (ssq model.SSQ, err error) {
	var ssqList []model.SSQ
	ssqList, err = this.getSSQByRemote(1)
	if len(ssqList) >= 1 {
		ssq = ssqList[0]
	} else {
		err = errors.New("ssqlist count == 0")
	}
	return
}

func (this *Lottery) getSSQByRemote(count int) (ssqList []model.SSQ, err error) {
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

func (this *Lottery) Get100SSQByRemote() (ssqList []model.SSQ, err error) {
	return this.getSSQByRemote(100)
}
