package test

import (
	"fmt"
	"goskeleton/app/model"
	_ "goskeleton/bootstrap"
	"testing"
	"time"
)

// 新增 与查询，可以在 config/config.yml 开启读写分离配置项，进行测试读写分离
func TestSqlInsert(t *testing.T) {
	sqlFact := model.CreateTestFactory("")
	if sqlFact.InsertData() {
		fmt.Println("数据插入成功")
	} else {
		t.Errorf("数据插入操作，单元测试失败")
	}

	//  如果 开启了读写分离，查询则连接的 Read 库
	list := sqlFact.QueryData()
	if list != nil {
		for index, item := range list {
			fmt.Printf("%d, %s,%d, %d, %s, %s\n", index, item.Name, item.Age, item.Sex, item.Addr, item.Remark)
		}
	} else {
		t.Errorf("数据查询操作，单元测试失败")
	}
}

// 查询（多条）
func TestSqlSelect(t *testing.T) {
	list := model.CreateTestFactory("").QueryData()
	if list != nil {
		for index, item := range list {
			fmt.Printf("%d, %s,%d, %d, %s, %s\n", index, item.Name, item.Age, item.Sex, item.Addr, item.Remark)
		}
	} else {
		t.Errorf("数据查询操作，单元测试失败")
	}
}

// 查询（单条）
func TestSqlSelectOne(t *testing.T) {
	oneList := model.CreateTestFactory("").QueryRowData()
	if oneList == nil {
		t.Errorf("单元测试：单条数据查询失败")
	} else {
		fmt.Printf("%#+v\n", *oneList)
	}
}

// 测试提交事务的操作
func TestSqlTransAction(t *testing.T) {
	// 修改以下函数的参数，测试事务的提交（true）与回滚（false）
	if model.CreateTestFactory("").TransAction(true) {
		fmt.Println("数据插入成功(提交事务操作)")
	} else {
		t.Errorf("数据插入（提交事务操作），单元测试失败")
	}
}

// 测试回滚事务的操作
func TestSqlTransAction2(t *testing.T) {
	// 参数 true 表示 提交事务；  false 表示 回滚事务
	if model.CreateTestFactory("").TransAction(false) {
		fmt.Println("数据插入成功(回滚事务操作)")
	} else {
		t.Errorf("数据插入（回滚事务操作），单元测试失败！")
	}
}

// 批量插入数据的正确姿势
func TestSqlInsertMultiple(t *testing.T) {

	if model.CreateTestFactory("").InsertDataMultiple() {
		fmt.Println("批量插入数据OK")
	} else {
		t.Errorf("批量插入数据，单元测试失败！")
	}
}

// 批量插入数据的错误姿势
func TestSqlInsertMultipleError(t *testing.T) {
	if model.CreateTestFactory("").InsertDataMultipleErrorMethod() {
		fmt.Println("批量插入数据OK")
	} else {
		t.Errorf("批量插入数据，单元测试失败！")
	}
}

// 批量查询数据，测试 pprof cpu性能
func TestSqlSelecttMultiple(t *testing.T) {

	if model.CreateTestFactory("").SelectDataMultiple() {
		fmt.Println("批量查询数据OK")
	} else {
		t.Errorf("批量查询数据出错")
	}
}

// 测试sql注入
func TestSqlInject(t *testing.T) {
	model.CreateTestFactory("").QueryInject()
}

// 测试连接池的获取与释放、对比数据库连接数量

func TestConnPool(t *testing.T) {

	//  获取20个连接，数据库客户端使用命令： SHOW  PROCESSLIST 查看正在连接的数量，发现增多了不少
	for i := 1; i <= 20; i++ {
		go func() {
			tmpAddr := model.CreateTestFactory("")
			fmt.Printf("获取的数据库连接池地址：%p\n", tmpAddr)
			time.Sleep(time.Second * 20)
		}()
	}
	fmt.Printf("20秒以后应该释放连接池并不会释放，数据库使用 SHOW  PROCESSLIST  查看依然能看见很多连接...")
	time.Sleep(time.Second * 30)
	//  知道本函数执行完毕，相关的连接会自动释放，继续使用  SHOW  PROCESSLIST 查看验证
}
