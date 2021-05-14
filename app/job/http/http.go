package http

import (
	"geak/job/conf"
	"geak/job/service"
)

var svr *service.Service

// Init init http router.
func Init(c *conf.Config, s *service.Service) {
	// init internal router
	svr = s
}