package service

import (
	"context"
	"geak/job/dao"
	"geak/libs/conf"
	"github.com/robfig/cron"
	"sync"
)

const (


	ssqCronSpec = "0 */1 21-23 * * TUE,THU,SUN"

	dltCronSpec1 = "0 30/1 20 * * MON,WED,SAT"
	dltCronSpec2 = "0 */1 21-23 * * MON,WED,SAT"

	otherCronSpec = "0 0 */1 * * ?"

	//otherCronSpec = "*/5 * * * * ?"
)


type Service struct {
	dao	*dao.Dao
	c				*conf.Config
	waiter   		*sync.WaitGroup
	// cron
	cron 			*cron.Cron
	lastestSSQCode 	string
	lastestDLTCode	string
}

func New(c *conf.Config) (s *Service) {

	s = &Service{
		dao:	dao.New(c),
		c:			c,
		waiter:		new(sync.WaitGroup),
		cron:		cron.New(),
	}

	s.cron.AddFunc(ssqCronSpec,s.StartSSQJob)
	s.cron.AddFunc(dltCronSpec1,s.StartDLTJob)
	s.cron.AddFunc(dltCronSpec2,s.StartDLTJob)
	s.cron.AddFunc(otherCronSpec,s.StartSSQJob)
	s.cron.AddFunc(otherCronSpec,s.StartDLTJob)
	s.cron.Start()

	return s
}


// Close close service.
func (this *Service) Close() (err error) {
	defer this.waiter.Wait()
	this.cron.Stop()
	return
}

// Ping check service health.
func (this *Service) Ping(c context.Context) error {
	return this.dao.Ping(c)
}