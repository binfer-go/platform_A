package main

import (
	"github.com/gogf/gf/frame/g"
	_ "platform/app/hander/env"
	_ "platform/boot"
	_ "platform/library/redis"
	_ "platform/router"
)
func main()  {
	g.Server().Run()
}
