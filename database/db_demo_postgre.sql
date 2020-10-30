--  请自行创建  数据库,例如：  db_goskeleton ,以及方案，例如：  web

CREATE TABLE web.tb_users
(
    id serial ,
    user_name   varchar(30) ,
    sex int,
    age int,
    addr varchar(120),
    remark text COLLATE pg_catalog."default",
	created_at  DATE,
	updated_at  DATE

);

insert  into  web.tb_users(user_name,sex,age,addr,remark,created_at,updated_at)  values
('goskeleton1',1,18,'postgresql 测试数据_postgre','备注信息001',current_date,current_date),
('goskeleton2',1,18,'postgresql 测试数据_postgre2','备注信息002',current_date,current_date),
('goskeleton3',1,18,'postgresql 测试数据_postgre3','备注信息003',current_date,current_date),
('goskeleton4',1,18,'postgresql 测试数据_postgre4','备注信息004',current_date,current_date),
('goskeleton5',1,18,'postgresql 测试数据_postgre5','备注信息005',current_date,current_date);