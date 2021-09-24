###  表单参数提交介绍  
 - 1.前端提交简单的表单参数示例代码，[请参考已有的接口测试用例文档](./api_doc.md)  
 - 2.本篇我们将介绍复杂表单参数的提交.  

#### 什么是简单的表单参数提交
> 1.如果接口参数都是简单的键值对,没有嵌套关系,就是简单模式.    

![form-parms](https://www.ginskeleton.com/images/formparams1.png)  

#### 什么是复杂的表单参数提交
> 1.表单参数存在嵌套关系,这种数据在 `postman` 都是以 raw 方式提交,本质上就是请求的表单参数头设置为：`Content-Type: application/json`   

![form-parms](https://www.ginskeleton.com/images/formparams2.png)  

#### `ginskeleton` 后台处理复杂表单数据  
> 1.按照提交的数据格式,我们在表单参数验证器部分,定义接受的结构体,例如上图的参数我们在后台的接受参数就可以定义如下：
```code  

type ViewEleCreateUpdate struct {
	FkBigScreenView     float64         `form:"fk_big_screen_view" json:"fk_big_screen_view"`
	EleId               string          `form:"ele_id" json:"ele_id"`
	EleIdTitle          string          `form:"ele_id_title" json:"ele_id_title"`
	Status              *float64        `form:"status" json:"status"`
	Remark              string          `form:"remark" json:"remark"`
	ChildrenTableDelIds string          `form:"children_table_del_ids" json:"children_table_del_ids"`
	ChildrenTable       []ChildrenTable `form:"children_table" json:"children_table"`
}

//  大屏界面元素的子表数据
//  每种元素都有三个状态（1=正常；2=禁止；3=隐藏）
//  被嵌套的数据请独立定义，这样的好处就是后续可以随意精准取出任意一部分
type ChildrenTable struct {
	Id                               float64  `form:"id" json:"id"`
	FkBigScreenViewElement           float64  `form:"fk_big_screen_view_element" json:"fk_big_screen_view_element"`
	FkBigScreenViewElementStatusName float64  `form:"fk_big_screen_view_element_status_name" json:"fk_big_screen_view_element_status_name"`
	Status                           *float64 `form:"status" json:"status"`
	Remark                           string   `form:"remark" json:"remark"`
}

```
#### 接口验证器  ↓ 
> 1.复杂接口参数前端都是通过json格式提交.  
> 2.`go` 语言代码接收语法是 `context.ShouldBindJSON()`  

![form-parms3](https://www.ginskeleton.com/images/formparams3.png)  

#### 接口验证器对应的数据类型  ↓  
![form-parms4](https://www.ginskeleton.com/images/formparams4.png)  

#### 在后续的控制器、model 获取子表数据
```code  
# 在接口验证逻辑部分，通过参数验证后，我们将子表数据已经存储在上线文

// 子表数据设置一个独立的键存储
extraAddBindDataContext.Set(consts.ValidatorPrefix+"children_table_del_ids", v.ChildrenTable)

// 那么后续的控制器、以及model都可以根据相关的键获取原始数据、断言为我们定义的子表数据类型继续操作
    var  childrenTableData = c.MustGet(consts.ValidatorPrefix + "children_table_del_ids")
    
    // 获取子表数据断言为我们定义的子表数据类型
    // 这里需要注意：验证器验证参数ok调用了控制器，如果再验证器文件没有创建独立的数据类型文件夹（包）,在控制器断言会形成包的嵌套、报错，这就是我们一开始将复杂数据类型创建独立的文件件定义的原因

	if subTableStr, ok := childrenTableData.([]data_type_for_create_edit.ChildrenTable); ok {
	    // 这里就相当于获取了go语言切片数据
	    // 继续批量存储、或者挨个遍历就行
	    //  ....   省略业务逻辑
	}

```



