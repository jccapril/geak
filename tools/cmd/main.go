package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)


type SSQ struct {
	Code 		string
	Date 		string
	Red		string
	Blue		string
	Blue2		string
	Sales 		string
	PoolMoney 	string
	FirstCount 	string
	FirstMoney 	string
	SecondCount string
	SecondMoney	string
	ThirdCount 	string
	ThirdMoney 	string
}

func main(){

	for i := 2003; i <= 2021; i++ {
		importDataFrom(i)
	}

}

func importDataFrom(year int){
	b, err := ioutil.ReadFile(fmt.Sprintf("../resources/ssq_history/%d.html",year)) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	expr :="<tbody class=\"list-tr\">([\\s\\S]*?)</tbody>"
	r,_ := regexp.Compile(expr)
	ssqListStr := r.FindAllStringSubmatch(str,1)[0][1]
	expr = "<tr[\\s\\S]*?>([\\s\\S]*?)</tr>"
	r,_ = regexp.Compile(expr)
	ssqList := r.FindAllStringSubmatch(ssqListStr,162)
	expr = "<td>([\\s\\S]*?)</td>"
	r,_ = regexp.Compile(expr)
	results := make([]*SSQ,0)
	for _, ssq := range ssqList {
		val := new(SSQ)
		result := r.FindAllStringSubmatch(ssq[1],12)
		if len(result) < 12 {
			break
		}
		val.Code = result[0][1]

		val.Date = result[1][1]
		if strings.Split(val.Date,"-")[0] != strconv.Itoa(year) {
			log.Fatalf("year:%d 有问题",year)
		}
		ballString := result[2][1]
		redList := regexp.MustCompile("<span class=\"red\">([\\s\\S]*?)</span>").FindAllStringSubmatch(ballString,6)
		var reds []string

		for _, red := range redList {
			reds = append(reds, red[1])
		}
		val.Red = strings.Join(reds,",")

		blueList := regexp.MustCompile("<span class=\"blue\">([\\s\\S]*?)</span>").FindAllStringSubmatch(ballString,2)
		for i, blue := range blueList {
			if i == 0 {
				val.Blue = blue[1]
			}else if(i == 1) {
				val.Blue2 = blue[1]
			}
		}
		val.Sales = result[3][1]
		val.PoolMoney = result[4][1]
		val.FirstCount = result[5][1]
		val.FirstMoney = result[6][1]
		val.SecondCount = result[7][1]
		val.SecondMoney = result[8][1]
		val.ThirdCount = result[9][1]
		val.ThirdMoney = result[10][1]
		results = append(results, val)
		fmt.Println(val)
	}
	fmt.Println("--------------------")
	fmt.Println("当年开出",len(results),"期")
	fmt.Println("--------------------")
}
