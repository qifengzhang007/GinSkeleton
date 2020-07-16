### Sql操作命令集合  
>本文档主要介绍了核心的sql操作命令，您也可以直接查看示例代码  [sql示例文档](../App\Model\Test.go)       
#### 查询类，主要是 select 性质的sql  
```sql
    #1.多条查询： 
        QuerySql
    #2.单条查询： 
        QueryRow
```

#### 执行类  
```sql
    #1.执行命令，主要有 insert 、 updated 、 delete   
       ExecuteSql
```       

#### 预处理类，如果场景需要批量插入很多条数据，那么久独立预编译
```sql
    #1.预编译
       PrepareSql
```        

#### 预处理类之后，执行批量语句
```sql
    #1.执行类
       ExecuteSqlForMultiple
    
    #2.查询类
       QuerySqlForMultiple             

``` 

#### 事务类操作
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
  