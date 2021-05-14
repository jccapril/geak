package gkd

import (
	"fmt"
	"geak/libs/conf"
	"geak/libs/database"
	"geak/libs/log"
	sql "github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"io/ioutil"
	"regexp"
	"strings"
	"sync"
)

var db *sql.DB

const (
	bodyExpr = "<tbody class=\"list-tr\">([\\s\\S]*?)</tbody>"
	ssqListExpr = "<tr[\\s\\S]*?>([\\s\\S]*?)</tr>"
	ssqElementExpr = "<td>([\\s\\S]*?)</td>"
	ssqRedBallExpr = "<span class=\"red\">([\\s\\S]*?)</span>"
	ssqBlueBallExpr = "<span class=\"blue\">([\\s\\S]*?)</span>"

)


func Init(config *conf.Config){
	 db = database.NewMySQL(config.DB)
}

//var ch = make(chan []*SSQ)
var wg = sync.WaitGroup{}

func InitData(){
	err := createSSQTable()
	if err != nil {
		log.Fatal("创建 ssq 表 失败",zap.Error(err))
	}
	ch := make(chan []*SSQ)

	procData(ch)
	go consumeData(ch)

	wg.Wait()
}

func procData(ch chan<- []*SSQ) {
	for year := 2003; year <= 2021; year++ {
		wg.Add(1)
		go initDataFrom(year,ch)
	}
}

func consumeData(ch <-chan []*SSQ){
	for {
		result := <-ch
		insertYearData(result)
		wg.Done()
	}
}


func insertYearData(data []*SSQ){
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
		log.Fatal("ls",zap.Error(err))
	}
	defer stmt.Close()
	//format all vals at once
	res, err := stmt.Exec(vals...)
	if err != nil {
		log.Error("sql exec 失败",zap.Error(err))
		return
	}
	count, err := res.RowsAffected()
	if err != nil {
		log.Error("获取影响条目数量失败",zap.Error(err))
	}
	log.Info(fmt.Sprintf("数据导入成功,一共导入%d条数据\n",count))

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
		third_money VARCHAR(20)
        )charset=utf8;`
	_,err = db.Exec(sqlStr)
	return
}


func initDataFrom(year int,ch chan<- []*SSQ){
	log.Info(fmt.Sprintf("开始导入%d年数据\n",year))
	content, err := ioutil.ReadFile(fmt.Sprintf("%s/ssq_history/%d.html",conf.Conf.App.Resources,year)) // just pass the file name
	if err != nil {
		fmt.Print(err)

	}
	input := string(content)
	body := regexp.MustCompile(bodyExpr).FindAllStringSubmatch(input,1)[0][1]
	ssqList := regexp.MustCompile(ssqListExpr).FindAllStringSubmatch(body,162)
	results := make([]*SSQ,0)
	for _, ssq := range ssqList {
		val := new(SSQ)
		result := regexp.MustCompile(ssqElementExpr).FindAllStringSubmatch(ssq[1],12)
		if len(result) < 12 {
			break
		}
		val.Code = result[0][1]

		val.Date = result[1][1]
		//if strings.Split(val.Date,"-")[0] != strconv.Itoa(year) {
		//	log.Fatalf("year:%d 有问题",year)
		//	log.Fatal(f)
		//}
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
	ch<-results
}
