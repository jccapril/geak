package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"geak/job/model"
	"geak/libs/log"
	"github.com/goinggo/mapstructure"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	dlt_host          = "https://webapi.sporttery.cn/gateway/lottery/getDigitalDrawInfoV1.qry?param=85,0&isVerify=1"
	last_dlt_code_key = "last_dlt_code_key"
	dlt_expiration	  = 7*24*3600*time.Second
)

func (this *Service)StartDLTJob(){
	dlt,err := this.GETLastestDLTByRemote()
	if err != nil {
		log.Error("dlt remote",zap.Error(err))
	}else {
		code := fmt.Sprintf("20%s",dlt.LotteryDrawNum)
		date := strings.Split(dlt.LotteryDrawTime," ")[0]
		balls := strings.Split(dlt.LotteryDrawResult," ")
		redBalls := balls[0:5]
		blueBalls := balls[5:7]
		red := strings.Join(redBalls,",")
		poolMoney := strings.ReplaceAll(dlt.PoolBalanceAfterdraw,",","")
		isExist := this.GetDLTCountFromDBByCode(code) > 0
		if isExist {
			log.Info("update dlt",zap.Any("dlt",dlt))
			sqlStr := "UPDATE `dlt` SET `date`=?,`red`=?,`blue`=?,`blue2`=?,`pool_money`=?,`content`=? WHERE `code`=?"
			_,err := this.dao.DB.Exec(sqlStr,date,red,blueBalls[0],blueBalls[1],poolMoney,dlt.DrawPdfUrl,code)
			if err != nil {
				log.Error(sqlStr,zap.Error(err))
			}
		}else {
			log.Info("insert new dlt",zap.Any("dlt",dlt))
			sqlStr := "INSERT INTO `dlt`(`code`, `date`, `red`, `blue`," +
				"`blue2`, `pool_money`,`content`) VALUES (?, ?, ?, ?, ?, ?, ?)"
			_,err := this.dao.DB.Exec(sqlStr,code,date,red,blueBalls[0],blueBalls[1],poolMoney,dlt.DrawPdfUrl)
			if err != nil {
				log.Error(sqlStr,zap.Error(err))
			}else {
				this.Push()
			}
		}
	}


}



func (this *Service)GetDLTCountFromDBByCode(code string)(num int) {
	sqlStr := "SELECT COUNT(*) FROM `dlt` WHERE `code`=?"
	err := this.dao.DB.Get(&num,sqlStr,code)
	if err != nil {
		log.Error(sqlStr,zap.Error(err))
	}
	return
}

func (this *Service)GETLastestDLTByRemote()(dlt model.DLT, err error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, dlt_host, nil)
	//req.Header.Set("Referer", "http://www.cwl.gov.cn/")
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
	value,isOK := result["value"].(map[string]interface{})
	if !isOK {
		return dlt,errors.New("format error")
	}
    d,isOK := value["dlt"].(map[string]interface{})
	if !isOK {
		return dlt,errors.New("format error")
	}
	err = mapstructure.Decode(d, &dlt)
	if err == nil {
		this.dao.RDB.Set(last_dlt_code_key,dlt.LotteryDrawNum,dlt_expiration)
	}

	return
}

