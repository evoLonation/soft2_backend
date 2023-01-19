create table user(
                     user_id int auto_increment,
                     login_id varchar(20) not null ,
                     password varchar(20) not null ,
                     nickname varchar(20) default 'DEFAULT',
                     email varchar(20) not null ,
                     introduction varchar(255),
                     complaints int default 0,
                     requests int default 0 comment '求助次数',
                     help int default 0 comment '应助次数',
                     follows int default 0,
                     primary key (user_id),
                     unique key (login_id)
)default charset =utf8mb4;