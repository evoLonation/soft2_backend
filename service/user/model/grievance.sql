create table grievance(
    grievance_id int primary key auto_increment,
    plaintiff_id int not null comment '申诉学者id',
    defendant_id int not null comment '被申诉学者id',
    paper_id int not null ,
    primary key (grievance_id),
    key idx_pid (plaintiff_id),
    key idx_did (defendant_id)
)default charset =utf8mb4;# 应该还有一个时间作为区分