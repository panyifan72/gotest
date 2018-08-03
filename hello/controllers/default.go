package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type MainController struct {
	beego.Controller
}
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
func (this *MainController) Test(){
	chs := make([]chan int,10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go b(chs[i],i)
	}

	for _,ch := range chs{
		<-ch
	}
	this.Ctx.WriteString("=========")
}
func b(ch chan int,ds int){
	ch<- ds+10
	fmt.Println("---")
	fmt.Println(ds)
	fmt.Println("---")
}
