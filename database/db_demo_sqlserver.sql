-- 创建数据库,例如：  db_goskeleton
USE [master]
IF NOT EXISTS(SELECT 1 FROM sysdatabases WHERE NAME=N'db_goskeleton')
BEGIN
CREATE DATABASE db_goskeleton
END
GO
use db_goskeleton ;
--  创建表  tb_users 以及 写入模拟数据
CREATE TABLE   tb_users (
	id int   IDENTITY(1,1) NOT NULL  primary  key,
	user_name nvarchar(60) DEFAULT(''),
	pass varchar(60) DEFAULT(''),
	age int DEFAULT(0),
	sex int DEFAULT(1),
	remark nvarchar(120) DEFAULT(''),
	created_at datetime ,
	updated_at datetime DEFAULT (getdate())
) ;


-- 模拟插入数据

insert   into tb_users (
user_name,
pass,
sex,
age,
remark,
created_at,
updated_at
)
values
('goskeleton1','123456789',1,18,'备注信息，测试！',getdate(),getdate())  ;

insert   into tb_users (
user_name,
pass,
sex,
age,
remark,
created_at,
updated_at
)
values
('goskeleton2','123456789',1,18,'备注信息，测试！',getdate(),getdate())  ;

insert   into tb_users (
user_name,
pass,
sex,
age,
remark,
created_at,
updated_at
)
values
('goskeleton3','123456789',1,18,'备注信息，测试！',getdate(),getdate())  ;