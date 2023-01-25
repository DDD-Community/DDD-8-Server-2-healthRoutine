CREATE TABLE exercise_category (
    id BIGINT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    subject VARCHAR(30) NOT NULL
);

CREATE TABLE exercise (
    id BIGINT AUTO_INCREMENT NOT NULL PRIMARY KEY ,
    subject VARCHAR(50) NOT NULL ,
    category_id BIGINT NOT NULL ,
    user_id CHAR(36) NULL,
    FOREIGN KEY (category_id) REFERENCES exercise_category (id)
);

CREATE TABLE health (
    id CHAR(36) NOT NULL PRIMARY KEY ,
    user_id CHAR(36) NOT NULL ,
    exercise_id BIGINT NOT NULL ,
    weight SMALLINT NOT NULL DEFAULT 0,
    reps SMALLINT NOT NULL DEFAULT 0,
    `set` SMALLINT NOT NULL DEFAULT 0,
    created_at BIGINT NOT NULL
);

CREATE INDEX health_user_id_idx ON health (user_id);
CREATE INDEX health_exercise_id_idx ON health (exercise_id);
CREATE INDEX health_created_at_idx ON health (created_at);

INSERT INTO exercise_category (subject)
VALUES
    ('가슴'),
    ('어깨'),
    ('등'),
    ('복근'),
    ('팔'),
    ('하체'),
    ('엉덩이'),
    ('전신');

INSERT INTO exercise (subject, category_id)
VALUES
    ('덤벨 플라이', 1),
    ('딥스', 1),
    ('디클라인 벤치프레스', 1),
    ('스퀴즈 프레스', 1),
    ('시티드 체스트 프레스', 1),
    ('인클라인 해머 프레스', 1),
    ('인클라인 덤벨 프레스', 1),
    ('케이블 크로스 오버', 1),
    ('푸쉬 업', 1),
    ('덤벨 숄더 프레스', 2),
    ('덤벨 레이즈', 2),
    ('리어 델트 로우', 2),
    ('리버스 풀 다운', 2),
    ('리버스 펙덱 플라이', 2),
    ('바벨 업라이트 로우', 2),
    ('바벨 숄더 프레스', 2),
    ('밀리터리 프레스', 2),
    ('리어델트 레이즈', 2),
    ('벤트 오버 레터럴 레이즈', 2),
    ('사이드 레터럴 레이즈', 2),
    ('비하인드 숄더 프레스', 2),
    ('비하인드 넥 프레스', 2),
    ('오버헤드 숄더 프레스', 2),
    ('아놀드 프레스', 2),
    ('익스터널 로테이션', 2),
    ('케이블 플라이', 2),
    ('랫 풀 다운', 3),
    ('로우풀리', 3),
    ('루마니안 데드리프트', 3),
    ('바벨 로우', 3),
    ('바벨 슈러그', 3),
    ('벤트 오버 덤벨 로우', 3),
    ('스트레이트 암 킥백', 3),
    ('시티 드로우', 3),
    ('팬레이 로우', 3),
    ('케이블 하이로우', 3),
    ('AB 슬라이더', 4),
    ('로만 체어 사이드밴드', 4),
    ('사이드 크런치', 4),
    ('크런치', 4),
    ('라잉 트라이셉스 익스텐션', 5),
    ('덤벨 컬', 5),
    ('로프 풀오버', 5),
    ('로프 암 컬', 5),
    ('바벨 컬', 5),
    ('바이셉스 컬', 5),
    ('벤치 딥스', 5),
    ('스컬크러셔', 5),
    ('얼터네이트 덤벨 컬', 5),
    ('컨센트레이션 컬', 5),
    ('케이블 푸쉬다운', 5),
    ('레그컬', 6),
    ('레그레이즈', 6),
    ('레그 익스텐션', 6),
    ('런지', 6),
    ('백 스쿼트', 6),
    ('브이 스쿼트', 6),
    ('스테퍼', 6),
    ('와이드 스쿼트', 6),
    ('스티드 레그 컬', 6),
    ('원 레그 프레스', 6),
    ('워킹 런지', 6),
    ('카프 프레스', 6),
    ('핵 스쿼트', 6),
    ('힙 런지', 7),
    ('힙 어덕션', 7),
    ('힙 리프트', 7),
    ('덤벨 쓰러스터', 8),
    ('데드리프트', 8),
    ('버피테스트', 8),
    ('플랭크', 8);