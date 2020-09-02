package model

import (
	"fmt"
	"go.uber.org/zap"
	"goskeleton/app/global/variable"
	"goskeleton/app/utils/yml_config"
	"log"
	"strconv"
)

func CreateTestFactory(sqlType string) *Test {
	if len(sqlType) == 0 {
		sqlType = yml_config.CreateYamlFactory().GetString("UseDbType") //如果系统的某个模块需要使用非默认（mysql）数据库，例如 sqlserver，那么就在这里
	}
	dbDriver := CreateBaseSqlFactory(sqlType)
	if dbDriver != nil {
		return &Test{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("TestModel 工厂初始化失败")
	return nil
}

type Test struct {
	*BaseModel
	Id     int64
	Name   string
	Sex    int8
	Age    int8
	Addr   string
	Remark string
}

// 插入操作
func (t *Test) InsertData() bool {
	sql := "INSERT  INTO  tb_test(`name`,`sex`,`age`,`addr`,`remark`) VALUES(?,?,?,?,?)"
	// 函数 ExecuteSql 适合一次性（或者小批量）执行就完成的操作
	if t.ExecuteSql(sql, "姓名_测试001", 1, 18, "地址测试数据2020", "备注信息数据，测试使用") > 0 {
		return true
	}
	return false
}

// 查询操作，一般是多行查询
func (t *Test) QueryData() []Test {
	sql := "SELECT   `name`,`sex`,`age`,`addr`,`remark`   FROM  tb_test  ORDER   BY  id  DESC   LIMIT ?"
	// QuerySql 函数 适合一次性（或者小批量）执行就完成的操作
	rows := t.QuerySql(sql, 10)
	if rows != nil {
		var temp []Test
		for rows.Next() {
			_ = rows.Scan(&t.Name, &t.Sex, &t.Age, &t.Addr, &t.Remark)
			temp = append(temp, *t)
		}
		_ = rows.Close()
		return temp
	} else {
		return nil
	}
}

// 单行查询 QueryRow
func (t *Test) QueryRowData() *Test {
	sql := "SELECT  id, `name`,`sex`,`age`,`addr`,`remark`   FROM  tb_test  ORDER   BY  id  DESC   LIMIT ?"
	// 单条查询，这里虽然查询10条数据，但是只返回结果的第一条数据
	err := t.QueryRow(sql, 10).Scan(&t.Id, &t.Name, &t.Sex, &t.Age, &t.Addr, &t.Remark)
	if err == nil {
		return t
	} else {
		variable.ZapLog.Error("单行查询出错", zap.Error(err))
	}
	return nil
}

// 超多数据批量插入的正确姿势
func (t *Test) InsertDataMultiple() bool {
	sql := "INSERT  INTO  tb_test(`name`,`sex`,`age`,`addr`,`remark`) VALUES(?,?,?,?,?)"
	//1.首先独立预处理sql语句，无参数
	if t.PrepareSql(sql) {
		var age int8 = 18
		// 你可以模拟插入更多条数据，例如 1万+
		for i := 1; i <= 100; i++ {
			sex := i % 2
			//2.执行批量插入，注意 该函数的参数全部是 预处理 sql 的参数
			if t.ExecuteSqlForMultiple("姓名_测试_"+strconv.Itoa(i), sex, age, "地址测试数据,序号："+strconv.Itoa(i), "备注信息数据，测试使用，编号："+strconv.Itoa(i)) == -1 {
				variable.ZapLog.Sugar().Warn("sql执行失败，sql:", sql, "姓名_测试_"+strconv.Itoa(i), sex, age, "地址测试数据,序号："+strconv.Itoa(i), "备注信息数据，测试使用，编号："+strconv.Itoa(i))
			}
		}
	}
	return true
}

//  超多数据批量插入的错误姿势
func (t *Test) InsertDataMultipleErrorMethod() bool {
	sql := "INSERT  INTO  tb_test(`name`,`sex`,`age`,`addr`,`remark`) VALUES(?,?,?,?,?)"
	var age int8 = 18
	// 一次性插入 100 条数据，你可以模拟的更多，例如 1万+
	for i := 1; i <= 100; i++ {
		sex := i % 2
		//2.批量数据插入，如果 预处理 语句不独立调用，ExecuteSql 命令每一次都会进行预编译、执行，导致数据库预编译的sql超过系统默认值
		// 系统预处理sql数量超过 max_prepared_stmt_count（默认16382）设置的值
		// 报错信息： Error 1461: Can't create more than max_prepared_stmt_count statements (current value: 16382)
		if t.ExecuteSql(sql, "姓名_测试_"+strconv.Itoa(i), sex, age, "地址测试数据,序号："+strconv.Itoa(i), "备注信息数据，测试使用，编号："+strconv.Itoa(i)) == -1 {
			variable.ZapLog.Sugar().Warn("sql执行失败，sql:", sql, "姓名_测试_"+strconv.Itoa(i), sex, age, "地址测试数据,序号："+strconv.Itoa(i), "备注信息数据，测试使用，编号："+strconv.Itoa(i))
		}
	}
	return true
}

// 超多数据批量查询的正确姿势
func (t *Test) SelectDataMultiple() bool {
	// 如果您要亲自测试，请确保相关表存在，并且有数据
	sql := `
			SELECT
			code,name,company_name,concepts,indudtry,province,city,introduce,created_at 
			FROM
			db_stocks.tb_code_list 
			LIMIT 0, 1000 ;
		`
	//1.首先独立预处理sql语句，无参数
	if t.PrepareSql(sql) {
		// 你可以模拟插入更多条数据，例如 1万+
		var code, name, company_name, concepts, indudtry, province, city, introduce, created_at string

		type Column struct {
			Code         string `json:"code"`
			Name         string `json:"name"`
			Company_name string `json:"company_name"`
			Concepts     string `json:"concepts"`
			Indudtry     string `json:"indudtry"`
			Province     string `json:"province"`
			City         string `json:"city"`
			Introduce    string `json:"introduce"`
			Created_at   string `json:"created_at"`
		}

		for i := 1; i <= 500; i++ {
			var nColumn = make([]Column, 0)
			//2.执行批量查询
			rows := t.QuerySqlForMultiple()
			if rows == nil {
				variable.ZapLog.Sugar().Error("sql执行失败，sql:", sql)
				return false
			} else {
				for rows.Next() {
					_ = rows.Scan(&code, &name, &company_name, &concepts, &indudtry, &province, &city, &introduce, &created_at)
					oneColumn := Column{
						code,
						name,
						company_name,
						concepts,
						indudtry,
						province,
						city,
						introduce,
						created_at,
					}
					nColumn = append(nColumn, oneColumn)

				}
				//// 我们只输出最后一行数据
				if i == 500 {
					fmt.Println("循环结束，最终需要返回的结果成员数量：", len(nColumn))
					fmt.Printf("%#+v\n", nColumn)
				}
			}
			rows.Close()
		}
	}
	variable.ZapLog.Info("批量查询sql执行完毕！")
	return true
}

//  sql事物操作
func (t *Test) TransAction(isCommit bool) bool {
	sql := "INSERT  INTO  tb_test(`name`,`sex`,`age`,`addr`,`remark`) VALUES(?,?,?,?,?)"
	tx := t.BeginTx()
	if tx != nil {
		if vStm, err := tx.Prepare(sql); err == nil {
			if res, err2 := vStm.Exec("姓名_测试_事务测试", 1, 18, "地址测试数据2020，事务测试", "备注信息数据，事务测试"); err2 == nil {
				if _, err3 := res.RowsAffected(); err3 == nil {
					if isCommit {
						_ = tx.Commit()
						return true
					} else {
						_ = tx.Rollback()
						return true
					}
				} else {
					_ = tx.Rollback()
				}
			}
		}
	} else {
		fmt.Println("开启事物失败")
	}
	return false
}

//  测试sql注入
func (t *Test) QueryInject() {
	tmpStr := "1;update tb_test  set  remark='sql注入信息' where   id=1"
	sql := "SELECT  id, `name`,`sex`,`age`,`addr`,`remark`   FROM  tb_test where  id=? ORDER   BY  id  DESC   LIMIT ?"
	// 单条查询，这里虽然查询10条数据，但是只返回结果的第一条数据
	rows := t.QueryRow(sql, tmpStr, 10).Scan(&t.Id, &t.Name, &t.Sex, &t.Age, &t.Addr, &t.Remark)
	if rows == nil {
		log.Println("查询sql执行无数据")
	} else {

		log.Println("查询sql执行完成！请检查是否发生sql注入,行数：", rows)
	}

}
