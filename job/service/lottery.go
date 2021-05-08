package service

import (
	"encoding/json"
	"errors"
	"geak/job/model"

	"time"

	//"time"
	"fmt"
	"github.com/robfig/cron"
	"github.com/goinggo/mapstructure"
	"io/ioutil"
	"net/http"
)

const(
	ssq_url = "http://www.cwl.gov.cn/cwl_admin/kjxx/findDrawNotice?name=ssq&issueCount=1"
	last_ssq_code_key = "last_ssq_code_key"
)

type Lottery struct {
	s 				*Service
	isSSQContinue 	bool
	lastSSQCode		string
}

func (this *Lottery)Run() {
	if this.isSSQContinue {
		ssq,err := this.GetLatestSSQByRemote()
		if err != nil {
			fmt.Println(err)
			return
		}
		if this.lastSSQCode != ssq.Code {
			fmt.Println(ssq)
			this.lastSSQCode = ssq.Code
			this.isSSQContinue = false
			this.s.lotteryDao.RDB.Set(last_ssq_code_key,ssq.Code,time.Second * 3600 * 24 *7)
		}
	}else {
		fmt.Println("已经获取到了最新的结果了")
	}

}

func (this *Lottery)Register(s *Service) {
	this.s = s
	c := cron.New()
	defer c.Stop()
	//spec := "0 */1 12 * * MON,WED,SAT"

	c.AddFunc("0 0 16 * * ?", func() {
		this.isSSQContinue = true
		lastSSQCode, _ := this.s.lotteryDao.RDB.Get(last_ssq_code_key).Result()
		this.lastSSQCode = lastSSQCode
	})

	c.AddJob("0 */1 16 * * ?",this)
	c.Start()
	select{}
}


func (this *Lottery)GetLatestSSQByRemote()(ssq model.SSQ, err error){
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet,ssq_url,nil)
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
	resultList,isOK := result["result"].([]interface{})
	if !isOK {
		return ssq,errors.New("format error")
	}
	if len(resultList) > 0 {
		lastResult,isOK := resultList[0].(map[string]interface{})
		if !isOK {
			return ssq,errors.New("format error")
		}
		err = mapstructure.Decode(lastResult, &ssq)
		if err != nil {
			return
		}
	}else {
		err = errors.New("result is nil")
	}
	return
}