create table comment(
    comment_id int primary key auto_increment,
    user_id int not null ,
    paper_id int not null ,
    content varchar(255) not null ,
    likes int not null default 0, #点赞量
    create_time timestamp NULL default current_timestamp,
    primary key (comment_id)
)default charset =utf8mb4;