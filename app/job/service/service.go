package service

import (
	"context"
	"geak/job/dao"
	"geak/libs/conf"
	"github.com/robfig/cron"
	"sync"
)

const (
	//ssqCronSpec = "0 0 21 * * TUE,THU,SUN"
	ssqCronSpec = "0 1 * * * ?"
)


type Service struct {
	dao	*dao.Dao
	c			*conf.Config
	waiter   	*sync.WaitGroup
	// cron
	cron *cron.Cron
}

func New(c *conf.Config) (s *Service) {

	s = &Service{
		dao:	dao.New(c),
		c:			c,
		waiter:		new(sync.WaitGroup),
		cron:		cron.New(),
	}
	s.cron.AddFunc(ssqCronSpec, func() {
		s.StartSSQJob()
	})
	s.cron.Start()

	return
}


// Close close service.
func (s *Service) Close() {
	defer s.waiter.Wait()
	s.cron.Stop()
	return
}

// Ping check service health.
func (s *Service) Ping(c context.Context) error {
	return s.dao.Ping(c)
}