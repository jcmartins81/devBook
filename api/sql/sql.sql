CREATE DATABASE IF NOT EXISTS devBook;
USE devBook;

DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
  id int auto_increment primary key,
  nome varchar(50) not null,
  nick varchar(50) not null unique,
  email varchar(50) not null unique,
  senha varchar(20) not null unique,
  criadoEM timestamp default current_t€ý,€ý,imestamp()
) ENGINE=INNODB;

$ docker run --name some-mysql -e MYSQL_ROOT_PASSWORD=my-secret-pw -d mysql:tag
