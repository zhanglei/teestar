package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/hsluoyz/gitstar/api"
)

func getIPInfo(clientIP string) string {
	if clientIP == "" {
		return ""
	}

	ips := strings.Split(clientIP, ",")
	res := ""
	for i := range ips {
		ip := strings.TrimSpace(ips[i])
		desc := api.GetDescFromIP(ip)
		ipstr := fmt.Sprintf("%s: %s", ip, desc)
		if i != len(ips) - 1 {
			res += ipstr + " -> "
		} else {
			res += ipstr
		}
	}

	return res
}

func LogInfo(ctx *context.Context, f string, v ...interface{}) {
	var ipString string
	clientIP := ctx.Request.Header.Get("x-forwarded-for")
	ipString = fmt.Sprintf("(%s) ", getIPInfo(clientIP))

	logs.Info(ipString + f, v...)
}

func ReadLog() []string {
	f, err := os.Open("logs/gitstar.log")
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(bytes), "\n")
}
