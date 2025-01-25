-- 作成したDBへ切り替え
\c golang_for_everyone

-- スキーマ作成
CREATE SCHEMA IF NOT EXISTS chapter_07;

-- 権限追加
GRANT ALL PRIVILEGES ON SCHEMA chapter_07 TO postgres_user;