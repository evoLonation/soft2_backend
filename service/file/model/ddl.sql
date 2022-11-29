create table user_avatar
(
    user_id   int ,
    file_name varchar(255) not null,
    primary key (user_id)
);
create table help_file
(
    help_id   int,
    file_name varchar(255) not null,
    PRIMARY KEY (help_id)
);
# create table apply_file
# (
#     apply_id  int,
#     file_name varchar(255) not null,
#     primary key (apply_id)
# );
# create table scholar_avatar
# (
#     scholar_id  varchar(255),
#     file_name varchar(255) not null,
#     primary key (scholar_id)
# );
