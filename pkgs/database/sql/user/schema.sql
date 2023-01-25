CREATE TABLE users (
    id CHAR(36) NOT NULL PRIMARY KEY ,
    nickname VARCHAR(30) NOT NULL ,
    email VARCHAR(100) NOT NULL ,
    password VARCHAR(200) NOT NULL ,
    profile_img VARCHAR(300) NOT NULL ,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL
);

CREATE INDEX users_id_idx ON users (id);
CREATE INDEX users_email_idx ON users (email);
CREATE INDEX users_nickname_idx ON users (nickname);