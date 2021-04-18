package util

import (
	"regexp"
	"strconv"
	"strings"
)
/*
检测是否是域名https或http格式
*/
func DomainCheck(domain string) bool {
	var match bool
	//支持以http://或者https://开头并且域名中间有/的情况
	IsLine := "^((http://)|(https://))?([a-zA-Z0-9]([a-zA-Z0-9\\-]{0,61}[a-zA-Z0-9])?\\.)+[a-zA-Z]{2,6}(/)"
	//支持以http://或者https://开头并且域名中间没有/的情况
	NotLine := "^((http://)|(https://))?([a-zA-Z0-9]([a-zA-Z0-9\\-]{0,61}[a-zA-Z0-9])?\\.)+[a-zA-Z]{2,6}"
	match, _ = regexp.MatchString(IsLine, domain)
	if !match {
		match, _ = regexp.MatchString(NotLine, domain)
	}
	return match
}

/*
检测地址格式IP+端口的格式
127.0.0.1:8080
*/
func CheckAddr(addr string) bool {
	op := strings.Split(addr, ":")
	if len(op) != 2 {
		return false
	}
	if op[0] == "" {
		return false
	}
	val, err := strconv.Atoi(op[1])
	if err != nil {
		return false
	}
	if 0 > val && val > 65535 {
		return false
	}
	return true
}



