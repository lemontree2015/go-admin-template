package sys

import (
	"sscmgroup.com/app/service/sys"
)

var srv *sys.MgrService

func Init(s *sys.MgrService) {
	srv = s
}
