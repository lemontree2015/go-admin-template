package api

import (
	"sscmgroup.com/app/service/api"
)

var srv *api.Service

func Init(s *api.Service) {
	srv = s
}
