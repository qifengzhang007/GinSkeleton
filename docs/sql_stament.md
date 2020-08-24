### Sql操作命令集合  
>本文档主要介绍了sql操作的核心命令，详细操作命令示例代码参见 [mysql示例文档](../app/model/test.go).  [sqlserver测试用例](../test/db_sqlserver_test.go) , [postgreSql测试用例](../test/db_postgresql_test.go) 操作方式同 mysql .      

#### 1.查询类： 不会修改数据的sql、存储过程、视图
```sql
    // 首先获取一个数据连接
    sqlservConn := sql_factory.GetOneSqlClient("postgre")  // 参数为空,默认就是mysql驱动,您还可以传递 sqlserver 、 postgresql 参数获取对应数据库的一个连接.
    #1.多条查询： 
        sqlservConn.QuerySql
    #2.单条查询： 
         sqlservConn.QueryRow
```

#### 2.执行类： 会修改数据的sql、存储过程等  
```sql
    #1.执行命令，主要有 insert 、 updated 、 delete   
       sqlservConn.ExecuteSql
```       

#### 3.预处理类：如果场景需要批量插入很多条数据，那么就需要独立调用预编译
>   1.如果你的sql语句需要循环插入1万、5万、10万+数据。  
>   2.那么可能会报错:  Error 1461: Can't create more than max_prepared_stmt_count statements (current value: 16382)  
>   3.此时需要以下解决方案  
```sql
    #1.预编译，预处理类之后，执行批量语句
       sqlservConn.PrepareSql
    #2.（多条）执行类
       sqlservConn.ExecuteSqlForMultiple
    #3.（多条）查询类
       sqlservConn.QuerySqlForMultiple    
```        

#### 4.事务类操作
```sql
    #1.开启一个事务
       tx:=sqlservConn.BeginTx()
    
    #2.预编译sql
       tx.Prepare

    #3.执行sql
       tx.Exec

    #4.提交
       tx.Commit

    #5.回滚
       tx.Rollback         
``` 
  