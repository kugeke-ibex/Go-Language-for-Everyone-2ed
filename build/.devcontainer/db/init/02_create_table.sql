-- DB切り替え
\c golang_for_everyone

-- テーブル作成
DROP TABLE IF EXISTS chapter_07.users

CREATE TABLE chapter_07.users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    age INTEGER NOT NULL
);

-- 権限追加
GRANT ALL PRIVILEGES ON chapter_07.users TO postgres_user;