package mysql

import (
	"geak/libs/conf"
	"geak/libs/database"
	"geak/libs/factory"
	"geak/libs/log"
	sql "github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"sync"
)

var db *sql.DB



func Init(config *conf.Config){
	 db = database.NewMySQL(config.DB)
}


var wg = sync.WaitGroup{}

func InitSSQData(){
	err := createSSQTable()
	if err != nil {
		log.Fatal("创建 ssq 表 失败",zap.Error(err))
	}

	ssq := new(SSQClient)
	ssq.Year = 2003
	m := factory.New(10,ssq)
	go m.Produce()
	go m.Consume()
	for i := 2003; i <= 2021; i++ {
		<-m.ResultQueue
	}
}

func ImportDLTData(){
	err := createDLTTable()
	if err != nil {
		log.Fatal("创建 dlt 表 失败",zap.Error(err))
	}
	dlt := new(DLTClient)
	dlt.Year = 2007
	m := factory.New(10,dlt)
	go m.Produce()
	go m.Consume()
	for i := 2007; i <= 2021; i++ {
		<-m.ResultQueue
	}
}








