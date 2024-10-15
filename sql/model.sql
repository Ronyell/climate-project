CREATE DATABASE IF NOT EXISTS elections;
USE elections;

DROP TABLE IF EXISTS cities;

CREATE TABLE cities(
    id int auto_increment primary key,
    name varchar(100) not null,
    uf varchar(2) not null,
    createdAt timestamp default current_timestamp()
) ENGINE=INNODB;