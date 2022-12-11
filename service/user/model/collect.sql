
create table collect(
                        collect_id int primary key auto_increment,
                        user_id int not null ,
                        paper_id varchar(20) not null ,
                        create_time timestamp NULL default current_timestamp,
                        primary key (collect_id),
                        key idx_uid (user_id),
                        key idx_pid (paper_id)
)default charset =utf8mb4;