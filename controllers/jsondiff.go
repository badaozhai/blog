package controllers

import (
	"github.com/astaxie/beego"
	"blog/tool"
	"fmt"
)

type JsonDiffController struct  {
	beego.Controller
}
func (c *JsonDiffController) Get(){
	aurl:=c.Input().Get("aurl")
	burl:=c.Input().Get("burl")
	var ajson string
	var bjson string
	if aurl!="" && burl !=""{
		ajson = tool.GetHtml(aurl)
		bjson = tool.GetHtml(burl)
	}
	c.Data["Aurl"] = aurl;
	c.Data["Burl"] = burl;
	c.Data["Ajson"] = ajson
	c.Data["Bjson"] = bjson
	fmt.Println(ajson)
	c.TplName = "json-diff.html";
}

func getHtml(url string) string {
	return ""
}