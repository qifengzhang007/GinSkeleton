package Admin

import (
	CacheModule2 "GinSkeleton/App/Cache/CacheModule"
	MyJwt2 "GinSkeleton/App/Http/Middleware/MyJwt"
	"GinSkeleton/App/Model"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Users struct {
}

// 1.模拟用户注册
func (u *Users) Register(context *gin.Context) {
	//  1.模拟用户注册
	context.JSON(200, gin.H{
		"code": 0,
		"msg":  "OK",
		"data": gin.H{
			"user":   "张三丰,模拟用户注册信息",
			"age":    20,
			"remark": "备注信息，2020-02-09",
		},
	})

}

//  2.模拟用户登录, 账号、密码 换token
func (u *Users) Login(context *gin.Context) {

	testRedis() // 测试redis

	userid := 2020 // 模拟一个用户id
	username := context.PostForm("username")
	pass := context.PostForm("pass")
	phone := "16601770915" // 模拟一个用户phone
	var Usertoken string

	fmt.Printf("username:%s,pass:%s, 验证账号、密码在数据库的合法性\n", username, pass)
	// 根据以上参数生成token，返回
	custome_claims := MyJwt2.CustomClaims{
		ID:    userid,
		Name:  username,
		Phone: phone,
		// 特别注意，针对前文的匿名结构体，初始化的时候必须指定键名，并且不带 jwt. 否则报错：Mixture of field: value and value initializers
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 10),   // 生效开始时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 失效截止时间
		},
	}

	if token, err := MyJwt2.CreateMyJWT().CreateToken(custome_claims); err == nil {
		fmt.Printf("账号密码换取token成功：%s\n", token)
		Usertoken = token
	}

	u.ParseUserToken(Usertoken)
	// 返回 json数据
	context.JSON(200, gin.H{
		"code": 0,
		"msg":  "登录成功",
		"data": gin.H{
			"user":   username,
			"age":    20,
			"token":  Usertoken,
			"remark": "备注信息，2020-02-09",
		},
	})

}

//  测试redis操作，假设登录成功将用户信息存储在redis
func testRedis() {

	cacheFactory := CacheModule2.CreateCacheFactory()
	res := cacheFactory.KeyExists("username")
	fmt.Printf("username键是否存在：%v\n", res)
	if res == false {
		res := cacheFactory.Set("username", "张三丰2012")
		fmt.Printf("username Set 值：%v\n", res)
	}

	res2 := cacheFactory.Get("username")

	fmt.Printf("username键是否存在：%v,取出相关值：%v\n", res, res2)
	cacheFactory.Release()

	/*	RedisClient:=RedisFactory.GetOneRedisClient()
		RedisClient.Execute("hSet","zhangqifeng","NO","070370122")
		RedisClient.Execute("hSet","zhangqifeng","universe","河北工程大学")
		res1,err1:=RedisClient.Execute("hGet","zhangqifeng","NO")
		res2,err:=RedisClient.Execute("hGet","zhangqifeng","universe")

		v_string1,v_err1:=RedisClient.String(res1,err1)
		v_string2,v_err2:=RedisClient.String(res2,err)
		if v_err2==nil && v_err1==nil{
			fmt.Printf("username= %#v, ex_key= %s\n",v_string1,v_string2)
		}

		RedisClient.RelaseOneRedisClientPool()*/

}

// 解析用户token的数据信息
func (c *Users) ParseUserToken(token string) {

	fmt.Println(token)
	if custome_claims, err := MyJwt2.CreateMyJWT().ParseToken(token); err == nil {
		fmt.Printf("token解析：%#v", custome_claims)
	} else {
		fmt.Printf("token解析失败：%v", err)
	}

}

//3.模拟查询一条用户记录
func (c *Users) ShowList(context *gin.Context) {

	Model.CreateUserFactory().ShowList(context.DefaultQuery("username", ""))

	/*	userid := context.Param("id")
		fmt.Printf("userid:%d\n", userid)
		context.JSON(http.StatusOK, gin.H{
			"code": "200",
			"msg":  "OK for select",
			"data": gin.H{
				"userid":   1,
				"username": "张三丰",
			},
		},
		)*/
}

//3.模拟新增一条用户记录
func (c *Users) Create(context *gin.Context) {
	username := context.PostForm("username")
	age := context.DefaultPostForm("age", "")
	sex := context.DefaultPostForm("sex", "")
	fmt.Printf("username:%s， age：%s, sex:%s \n", username, age, sex)
	context.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "OK for create",
		"data": gin.H{
			"userid":   2020,
			"username": username,
		},
	},
	)
}

//4.模拟更新一条用户记录
func (c *Users) Update(context *gin.Context) {
	userid := context.PostForm("id")
	username := context.PostForm("username")
	age := context.PostForm("age")
	sex := context.PostForm("sex")
	fmt.Printf("userid:%s, username:%s， age：%s, sex:%s \n", userid, username, age, sex)
	context.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "OK for update",
		"data": gin.H{
			"userid":   userid,
			"username": username,
		},
	},
	)
}

//5.模拟删除一条用户记录
func (c *Users) Delete(context *gin.Context) {
	userid := context.PostForm("id")
	fmt.Printf("userid:%s 将被删除\n", userid)
	context.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "OK for Delete",
		"data": gin.H{
			"userid": userid,
		},
	},
	)
}

//  6.上传头像
func (c *Users) UploadAvatar(context *gin.Context) {
	//  1.获取上传的文件名
	file, error := context.FormFile("avatar") //  file 是一个文件结构体（文件对象）
	if error != nil {
		context.String(http.StatusBadRequest, "上传文件发生错误")
		return
	}
	fmt.Printf("%s", file)
	//  保存文件
	if err := context.SaveUploadedFile(file, "./src/GinSkeleton/"+file.Filename); err != nil {
		context.String(http.StatusBadRequest, "文件保存失败，%v", err.Error())
		return
	}
	//  上传成功
	context.String(http.StatusCreated, "文件上传成功！")

}

/*// 7.展示用户的全部数据showlist
func (c *Users) ShowList(ctx gin.Context){

}
*/
