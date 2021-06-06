-- 创建数据库,例如：  db_goskeleton
USE [master]
IF NOT EXISTS(SELECT 1 FROM sysdatabases WHERE NAME=N'db_goskeleton')
BEGIN
CREATE DATABASE db_goskeleton
END
GO
use db_goskeleton ;
--  创建用户表
CREATE TABLE [dbo].[tb_users](
    [id] [int] IDENTITY(1,1) NOT NULL,
    [user_name] [nvarchar](50) NOT NULL ,
    [pass] [varchar](128) NOT NULL ,
    [real_name] [nvarchar](30)   DEFAULT (''),
    [phone] [char](11)    DEFAULT (''),
    [status] [tinyint]   DEFAULT (1),
    [remark] [nvarchar](120)    DEFAULT (''),
    [last_login_time] [datetime] DEFAULT (getdate()),
    [last_login_ip] [varchar](128) DEFAULT (''),
    [login_times] [int] DEFAULT ((0)),
    [created_at] [datetime]   DEFAULT (getdate()),
    [updated_at] [datetime]  DEFAULT (getdate())
    );
-- --  创建token表

CREATE TABLE [dbo].[tb_oauth_access_tokens](
    [id] [int] IDENTITY(1,1) NOT NULL,
    [fr_user_id] [int]  DEFAULT ((0)),
    [client_id] [int]  DEFAULT ((0)),
    [token] [varchar](600)  DEFAULT (''),
    [action_name] [varchar](50)   DEFAULT ('login') ,
    [scopes] [varchar](128) DEFAULT ('*') ,
    [revoked] [tinyint] DEFAULT ((0)),
    [client_ip] [varchar](128) DEFAULT (''),
    [created_at] [datetime]  DEFAULT (getdate()) ,
    [updated_at] [datetime]  DEFAULT (getdate()) ,
    [expires_at] [datetime]  DEFAULT (getdate()) ,
    [remark] [nchar](120) DEFAULT ('')
    ) ;