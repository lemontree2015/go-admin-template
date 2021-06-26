package sys

import (
	"sscmgroup.com/app/repository/sys"
)

type MgrService struct {
	rep *sys.Repository
}

func New() *MgrService {
	ser := &MgrService{
		rep: sys.New(),
	}
	return ser
}
