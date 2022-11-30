CREATE TABLE users
(
    id CHAR(36) NOT NULL PRIMARY KEY ,
    username VARCHAR(30) NOT NULL ,
    email VARCHAR(100) NOT NULL ,
    password VARCHAR(200) NOT NULL ,
    created_at BIGINT(20) NOT NULL
);

CREATE INDEX users_id_idx ON users (id);