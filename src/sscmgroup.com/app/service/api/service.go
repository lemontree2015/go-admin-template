package api

import (
	"golang.org/x/sync/singleflight"
	"sscmgroup.com/app/repository/api"
)

var singleRequest singleflight.Group

type Service struct {
	rep *api.Repository
}

func New() *Service {
	ser := &Service{
		rep: api.New(),
	}
	return ser
}
