package admin

import (
	"myoo/models"
	"strconv"

	"github.com/gin-gonic/gin"
	//	"myoo/redis"
)

type Category struct{}

func (*Category) Get(router *gin.Context) {
	category := new(models.Category)
	DbValue := category.QueryAll()
	router.JSON(200, gin.H{"success": 200, "message": nil, "data": DbValue})
	return
}
func (this *Category) Post(router *gin.Context) {
	switch router.PostForm("type") {
	case "insert":
		this.insert(router)
	case "update":
		this.update(router)
	case "delete":
		this.delete(router)
	default:
		router.JSON(200, gin.H{"success": 403, "message": nil, "data": nil})
	}
	return
}
func (this *Category) insert(router *gin.Context) {
	name := router.PostForm("name")
	if name == "" {
		router.JSON(200, gin.H{"success": 403, "message": "新建分类名称不能为空", "data": nil})
		return
	}
	category := new(models.Category)
	if category.Is(name) {
		router.JSON(200, gin.H{"success": 403, "message": "该分类已存在，不能重复创建", "data": nil})
		return
	}
	if category.Insert(name) > 0 {
		router.JSON(200, gin.H{"success": 200, "message": "新建分类成功", "data": nil})
		return
	}
	router.JSON(200, gin.H{"success": 403, "message": "新建分类失败", "data": nil})
	return
}

func (this *Category) update(router *gin.Context) {
	category := new(models.Category)
	name := router.PostForm("name")
	id, err := strconv.Atoi(router.PostForm("id"))
	if err != nil {
		router.JSON(200, gin.H{"success": 403, "message": err, "data": nil})
		return
	}

	if !category.Is(int64(id)) {
		router.JSON(200, gin.H{"success": 403, "message": "该分类不存在", "data": nil})
		return
	}
	if category.Update(int64(id), name) {
		router.JSON(200, gin.H{"success": 200, "message": "分类已更新", "data": nil})
		return
	}
	router.JSON(200, gin.H{"success": 403, "message": "分类更新失败", "data": nil})
	return
}

func (this *Category) delete(router *gin.Context) {
	path, err := strconv.Atoi(router.PostForm("id"))
	if err != nil {
		router.JSON(200, gin.H{"success": 403, "message": err, "data": nil})
		return
	}
	category := new(models.Category)
	id := int64(path)
	if category.Count() < 1 {
		router.JSON(200, gin.H{"success": 403, "message": "请至少留下一个分类", "data": nil})
		return
	}

	if category.QueryPosts(id) {
		router.JSON(200, gin.H{"success": 403, "message": "当前分类下存在文章,无法删除", "data": nil})
		return
	}
	if category.Delete(id) {
		router.JSON(200, gin.H{"success": 200, "message": "分类删除成功", "data": nil})
		return
	}
	router.JSON(200, gin.H{"success": 403, "message": "分类删除失败", "data": nil})
	return
}
