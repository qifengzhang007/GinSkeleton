package Admin

import (
	"GinSkeleton/App/Model"
	"fmt"
	"github.com/gin-gonic/gin"
)

//  创建model
var Model_codelist = Model.CreateCodeListFactory()

type CodeList struct {
}

// 1.查询数据
func (u *CodeList) Show(context *gin.Context) {

	//  get 方式 context.Query("code" ；post 方式 context.name:三六零, code:601366
	fmt.Printf("参数：\nname:%s, code:%s\n", context.Query("name"), context.Query("code"))

	//  1.查询codelist数据
	res := Model_codelist.GetCodeList()
	context.JSON(200, gin.H{
		"code": 0,
		"msg":  "OK for showlist",
		"data": res,
	})
}

// 2.新增数据
func (u *CodeList) Store(context *gin.Context) {

	//  create 模拟 post 数据
	fmt.Printf("参数：\nname:%s, code:%s\n", context.PostForm("name"), context.PostForm("code"))

	context.JSON(200, gin.H{
		"code": 0,
		"msg":  "OK",
		"data": "数据新增成功，返回最新ID：2020",
	})
}

// 3.更新数据
func (u *CodeList) Update(context *gin.Context) {

	//  create 模拟 post 数据
	fmt.Printf("参数：\nname:%s, code:%s， ID：%s\n", context.PostForm("name"), context.PostForm("code"), context.PostForm("id"))

	//  获取中间件设置的键对应的值
	val, _ := context.Get("test_key")
	if test_key_value, ok := val.(string); ok {
		fmt.Printf("获取中间件设置的值：%s\n", test_key_value)
	}

	context.JSON(200, gin.H{
		"code": 0,
		"msg":  "OK",
		"data": "数据 update 成功，返回最新ID：2019",
	})
}

// 3.删除数据
func (u *CodeList) Destroy(context *gin.Context) {

	//  create 模拟 post 数据
	ID := context.PostForm("id")
	fmt.Printf("参数：ID：%s\n", ID)

	context.JSON(200, gin.H{
		"code": 0,
		"msg":  "OK",
		"data": "数据 delete 成功，返回最新ID：" + ID,
	})
}
