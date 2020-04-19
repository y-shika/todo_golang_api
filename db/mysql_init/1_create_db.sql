-- データベース名を変えたら`.env.local`の`LOCAL_DB_NAME`も変える
CREATE DATABASE IF NOT EXISTS todo_golang_api;
USE todo_golang_api;

CREATE TABLE IF NOT EXISTS todos (
    -- idはuuid v4を使うことを想定して36としている
    id varchar(36) primary key,
    title text,
    active boolean,
    detail text,
    created_at datetime,
    updated_at datetime
);

INSERT INTO todos VALUES ("1", "Title_1", true, "Detail_1", NOW(), NOW()), ("2", "Title_2", true, "Detail_2", NOW(), NOW()), ("3", "Title_3", true, "Detail_3", NOW(), NOW());