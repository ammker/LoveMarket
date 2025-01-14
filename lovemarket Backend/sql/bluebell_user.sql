create table user
(
    user_id          bigint auto_increment
        primary key,
    username    varchar(64)                         not null,
    password    varchar(64)                         not null,
    email       varchar(64)                         null,
    createTime timestamp default CURRENT_TIMESTAMP null,
    updateTime timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    constraint idx_user_id
        unique (user_id),
    constraint idx_username
        unique (username)
)
    collate = utf8mb4_general_ci;
