package service

import (
	"context"
	"geak/job/dao"
	"geak/libs/conf"
	"github.com/robfig/cron"
)

const (
	ssqCronSpec = "0 0 21 * * TUE,THU,SUN"
	//ssqCronSpec = "0 10 * * * ?"
)


type Service struct {
	dao	*dao.Dao
	c			*conf.Config
	//waiter   	*sync.WaitGroup
	// cron
	cron *cron.Cron
}

func New(c *conf.Config) (s *Service) {

	s = &Service{
		dao:	dao.New(c),
		c:			c,
		//waiter:		new(sync.WaitGroup),
		cron:		cron.New(),
	}

	s.cron.AddFunc(ssqCronSpec, func() {
		s.FetchLastSSQByRemote()
	})
	s.cron.Start()

	return
}


// Close close service.
func (s *Service) Close() {
	s.cron.Stop()
}

// Wait wait routine unitl all close.
func (s *Service) Wait() {
	//s.waiter.Wait()
}

// Ping check service health.
func (s *Service) Ping(c context.Context) error {
	return s.dao.Ping(c)
}