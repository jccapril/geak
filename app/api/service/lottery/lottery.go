package lottery

import (
	"geak/api/dao"
	"geak/libs/conf"
)

type Service struct {
	dao	*dao.Dao
	c   *conf.Config
}

func New(c *conf.Config) (s *Service) {

	s = &Service{
		dao:	dao.New(c),
		c:			c,
	}
	return s
}

