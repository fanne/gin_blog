package main

import (
	"gin_blog/model"
	"gin_blog/routes"
)

func main() {
	//引用数据库
	model.InitDb()
	routes.InitRouter()
}
