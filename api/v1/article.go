package v1

import (
	"gin_blog/model"
	"gin_blog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//添加文章
func AddArt(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBindJSON(&data)
	code = model.CreateArt(&data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询单个文章
func GetArtInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetArtInfo(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询文章列表
func GetArt(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	data, code, total := model.GetArt(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

//编辑文章名称
func EditArt(c *gin.Context) {
	var data model.Article
	id, _ := strconv.Atoi(c.Param("id"))

	code = model.EditArt(id, &data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})

}

//删除文章
func DeleteArt(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteArt(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})

}

//查询分类下的所有文章
func GetCateArt(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	data, code, total := model.GetCateArt(id, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})

}
