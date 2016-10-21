mysql -uroot -p986652618erMAO -e "
drop database player;
create database player;
use player;
create table player(id char(36) not null primary key comment '#id', name varchar(255) not null comment 'player name', crdate datetime not null comment 'create time');"
