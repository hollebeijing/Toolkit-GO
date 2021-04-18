package request

import (
	"Toolkit-GO/request/rest"
	"Toolkit-GO/request/util"
	"fmt"
	"strings"
)

var DBClient *DBServer

type DBServer struct {
	Client   rest.ClientInterface
	TimeTask int
}

func Init(addr, subpath string, c *util.Capability, timeTask int) (*DBServer, error) {
	var base string
	addr = strings.Trim(addr, "/")
	subpath = strings.Trim(subpath, "/")

	if s := util.DomainCheck(addr); s {
		base = fmt.Sprintf("%s/%s/", addr, subpath)
	} else if i := util.CheckAddr(addr); i {
		base = fmt.Sprintf("http://%s/%s/", addr, subpath)
	} else {
		return nil, fmt.Errorf("addr format err！")
	}
	DBClient = &DBServer{
		Client:   rest.NewRESTClient(c, base),
		TimeTask: timeTask,
	}
	return DBClient, nil
}

