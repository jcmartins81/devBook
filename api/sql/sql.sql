CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS publicacoes;
DROP TABLE IF EXISTS seguidores;
DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
                         id int auto_increment primary key,
                         nome varchar(50) not null,
                         nick varchar(50) not null unique,
                         email varchar(50) not null unique,
                         senha varchar(100) not null,
                         criadoEm timestamp default current_timestamp()
) ENGINE=INNODB;

$ docker run --name some-mysql -e MYSQL_ROOT_PASSWORD=my-secret-pw -d mysql:tag
