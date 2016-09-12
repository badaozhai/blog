package controllers

import (
	"github.com/astaxie/beego"

	"go-executor/util/http_client"
	"go-executor/request"
	"fmt"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	session := request.NewSession()
	URL := `https://www.google.com`
	html,err:= http_get(session,URL)
	if err != nil{
		fmt.Println(err)
		c.Ctx.WriteString("hello,world")
	}else{
		c.Ctx.WriteString(html)
	}

}

func http_get(session *request.RequestSession,URL string)(string,error){
	httpclient := http_client.New(URL, nil)
	req := httpclient.Req.Header
	req.Add("Host", " www.google.com")
	req.Add("Connection", " keep-alive")
	req.Add("Cache-Control", " no-cache")
	req.Add("Accept", " */*")
	//req.Add("X-Requested-With", " XMLHttpRequest")
	//req.Add("If-Modified-Since", " 0")
	req.Add("User-Agent", " " + session.UserAgent)
	//req.Add("Referer", " https://kyfw.12306.cn/otn/leftTicket/init")
	req.Add("Accept-Encoding", " gzip, deflate, sdch, br")
	req.Add("Accept-Language", " zh-CN,zh;q=0.8,en;q=0.6")
	if session.Cookie.String() != "" {
		req.Add("Cookie", session.Cookie.String())
	}
	res, err := httpclient.Do()
	if err != nil {
		return "", err
	}
	// 更新cookie
	session.Cookie.Update(res)
	html, err := http_client.ReadResponse(res)
	if err != nil {
		return "", err
	}
	return html,nil
}