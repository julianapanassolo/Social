CREATE DATABASE IF NOT EXISTS social;
USE social;

DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios{
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick 
}
