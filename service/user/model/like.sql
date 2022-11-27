create table `like`(
    like_id int primary key auto_increment,
    user_id int not null ,
    comment_id int not null ,
    primary key (like_id),
    key idx_uid(user_id),
    key idx_cid(comment_id)
)default charset =utf8mb4;