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
var session *request.RequestSession = request.NewSession()
func (c *MainController) Get() {

	URL := `https://www.google.com.hk`

	html,err:= http_get(session,URL)
	if err != nil{
		fmt.Println(err)
		c.Ctx.WriteString("hello,world")
	}else{
		c.Ctx.WriteString(html)
	}
}
func (c *MainController) Search() {
	q := c.Input().Get("q")
	URL := `https://www.google.com.hk`
	URL = URL + `?q=`+q
	fmt.Println(URL)
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
	req.Add("Host", " www.google.com.hk")
	req.Add("Connection", " keep-alive")
	req.Add("Cache-Control", " no-cache")
	req.Add("Accept", " text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Add("upgrade-insecure-requests", " 1")
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