-- DB作成
CREATE DATABASE golang_for_everyone; 

-- 作成したDBへ切り替え
\c golang_for_everyone

-- スキーマ作成
CREATE SCHEMA chapter_07;

-- ロールの作成
CREATE ROLE postgres_user WITH LOGIN PASSWORD 'postgres_user_passw0rd';

-- 権限追加
GRANT ALL PRIVILEGES ON SCHEMA chapter_07 TO postgres_user;