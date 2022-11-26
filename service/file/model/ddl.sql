create table user_avatar
(
    user_id   varchar(255) ,
    file_name varchar(255) not null,
    primary key (user_id)
)ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 ;
create table help_file
(
    help_id   int,
    file_name varchar(255) not null,
    PRIMARY KEY (help_id)
)ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 ;
create table apply_file
(
    apply_id  int primary key,
    file_name varchar(255) not null,
    primary key (apply_id)
)ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 ;
