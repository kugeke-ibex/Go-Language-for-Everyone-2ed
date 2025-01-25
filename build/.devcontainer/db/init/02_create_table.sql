-- テーブル作成
DROP TABLE IF EXISTS chapter_07.users;

CREATE TABLE chapter_07.users (
    id SERIAL PRIMARY KEY,
    name TEXT NULL,
    age INTEGER NOT NULL
);

-- 権限追加
GRANT ALL PRIVILEGES ON chapter_07.users TO postgres_user;
