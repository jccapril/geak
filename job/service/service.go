package service

import (
	"context"
	"geak/job/conf"
	"geak/job/dao"
	"sync"

)

type Service struct {
	lotteryDao	*dao.Dao
	c			*conf.Config
	waiter   	*sync.WaitGroup
}

func New(c *conf.Config) (s *Service) {

	s = &Service{
		lotteryDao:	dao.New(c),
		c:			c,
		waiter:		new(sync.WaitGroup),
	}

	lottery := new(Lottery)
	go lottery.Register(s)


	return
}


// Close close service.
func (s *Service) Close() {

}

// Wait wait routine unitl all close.
func (s *Service) Wait() {
	s.waiter.Wait()
}

// Ping check service health.
func (s *Service) Ping(c context.Context) error {
	return s.lotteryDao.Ping(c)
}