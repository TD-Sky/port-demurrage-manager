drop table if exists login;

create table login
(
    username text primary key,      -- 用户名
    passwd text not NULL            -- 密码
);

insert into login (username, passwd)
values ('Klein',
        'f9ade17244dc5b67b9bf11f177cf019419b7ca6e50d22b3801c2ea4740fc0ade');
