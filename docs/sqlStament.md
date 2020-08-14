### Sql操作命令集合  
>本文档主要介绍了sql操作的核心命令，您也可以直接查看示例代码  [sql示例文档](../app/models/test.go)       
#### 1.查询类： 不会修改数据的sql、存储过程、视图
```sql
    #1.多条查询： 
        QuerySql
    #2.单条查询： 
        QueryRow
```

#### 2.执行类： 会修改数据的sql、存储过程等  
```sql
    #1.执行命令，主要有 insert 、 updated 、 delete   
       ExecuteSql
```       

#### 3.预处理类：如果场景需要批量插入很多条数据，那么就需要独立调用预编译
>   1.如果你的sql语句需要循环插入1万、5万、10万+数据。  
>   2.那么可能会报错:  Error 1461: Can't create more than max_prepared_stmt_count statements (current value: 16382)  
>   3.此时需要以下解决方案  
```sql
    #1.预编译，预处理类之后，执行批量语句
       PrepareSql
    #2.（多条）执行类
       ExecuteSqlForMultiple
    #3.（多条）查询类
       QuerySqlForMultiple    
```        

#### 4.事务类操作
```sql
    #1.开启一个事务
       tx:=BeginTx
    
    #2.预编译sql
       tx.Prepare

    #3.执行sql
       tx.Exec

    #4.提交
       tx.Commit

    #5.回滚
       tx.Rollback         
``` 
  