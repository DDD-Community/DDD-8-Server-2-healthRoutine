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

CREATE TABLE badge (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    subject VARCHAR(15) NOT NULL,
    sub VARCHAR(30) NOT NULL
);

CREATE INDEX badge_id_idx ON badge(id);

CREATE TABLE badge_users (
    users_id CHAR(36) NOT NULL,
    badge_id BIGINT NOT NULL,
    created_at BIGINT NOT NULL,
    FOREIGN KEY (users_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (badge_id) REFERENCES badge (id) ON DELETE CASCADE,
UNIQUE KEY (users_id, badge_id)
);

CREATE INDEX badge_users_users_id_idx ON badge_users (users_id);
CREATE INDEX badge_users_created_at_idx ON badge_users (created_at);

INSERT INTO badge (subject, sub)
VALUES
    ('운동의 시작' , 'exerciseStart'),
    ('운동의 기쁨' , 'exerciseHappy'),
    ('운동 홀릭', 'exerciseHolic'),
    ('운동 마스터', 'exerciseMaster'),
    ('운동 챔피언', 'exerciseChampion'),
    ('성실 주니어', 'sincerityJunior'),
    ('성실 프로', 'sincerityPro'),
    ('성실 마스터', 'sincerityMaster'),
    ('성실 챔피언', 'sincerityChampion'),
    ('꿀컵꿀컵', 'drinkHoneyHoney'),
    ('벌컵벌컵', 'drinkBulkUpBulkUp'),
    ('물 먹는 하마', 'drinkHippo');