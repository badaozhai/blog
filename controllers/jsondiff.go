package controllers

import "github.com/astaxie/beego"

type JsonDiffController struct  {
	beego.Controller
}
func (c *JsonDiffController) Get(){
	aurl:=c.Input().Get("aurl")
	burl:=c.Input().Get("burl")
	ajson := getHtml(aurl)
	bjson := getHtml(burl)
	c.Data["Ajson"] = ajson;
	c.Data["Bjson"] = bjson;
	c.TplName = "json-diff.html";
}

func getHtml(url string) string {
	return ""
}