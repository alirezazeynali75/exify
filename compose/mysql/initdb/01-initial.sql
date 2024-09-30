grant all privileges on *.* to 'root'@'%' with grant option;

create user `exify_test` identified by 'password';
create user `exify_dev` identified by 'password';

create database if not exists `exify`;
grant all on exify_dev.* to exify_dev@'%';

grant all on `exify_test_%`.* to exify_test@'%';
