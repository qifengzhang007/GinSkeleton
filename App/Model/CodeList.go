package Model

import (
	"GinSkeleton/App/Utils/Helper"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type CodeList struct {
	*BaseModel
	Code string `json:"code"`
	Name string `json:"name"`
}

func (c *CodeList) GetCodeList() []CodeList {
	var sql string
	sql = "SELECT `code`,`name` FROM   `db_stocks`.`tb_code_list`  WHERE   ?=? LIMIT 5"
	rows := c.QuerySql(sql, 1, 1)
	if rows != nil {
		var temp = CodeList{}
		var res = make([]CodeList, 0)
		for rows.Next() {
			err := rows.Scan(&temp.Name, &temp.Code)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("name:  %s, code:  %s\n", temp.Name, temp.Code)
				res = append(res, temp)
			}
		}
		//  凡是查询类记得释放记录集
		rows.Close()
		return res
	} else {
		fmt.Println("没有查询到数据", sql, "\n%v\n", rows)
	}
	return nil
}

func CreateCodeListFactory() *CodeList {
	configFac := Helper.CreateYamlFactory()
	DbType := configFac.GetString("DbType")
	dbDriver := CreateBaseSqlFactory(DbType)
	if dbDriver != nil {
		return &CodeList{
			BaseModel: dbDriver,
		}
	}
	log.Fatal("CodeListModel工厂初始化失败")
	return nil
}
