package routes

import (
	v1 "gin_blog/api/v1"
	"gin_blog/middleware"
	"gin_blog/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	// r := gin.Default()
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())

	{
		//User模块路由接口

		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DelUser)
		//Category模块路由接口
		auth.POST("cate/add", v1.AddCate)
		auth.PUT("cate/:id", v1.EditCate)
		auth.DELETE("cate/:id", v1.DelCate)
		//article模块路由接口
		auth.POST("article/add", v1.AddArt)

		auth.PUT("article/:id", v1.EditArt)
		auth.DELETE("article/:id", v1.DeleteArt)

	}

	route := r.Group("api/v1")
	{
		route.POST("user/add", v1.AddUser)
		route.GET("user", v1.GetUsers)
		route.GET("cate", v1.GetCate)
		route.GET("article", v1.GetArt)
		route.GET("article/list/:id", v1.GetCateArt)
		route.GET("article/info/:id", v1.GetArtInfo)
		route.POST("login", v1.Login)

	}

	r.Run(utils.HttpPoort)
}
