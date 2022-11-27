create table subscribe(
    subscribe_id int primary key auto_increment,
    user_id int not null ,
    scholar_id int not null ,
    primary key (subscribe_id),
    key idx_uid(user_id),
    key idx_sid(scholar_id)
)default charset =utf8mb4;