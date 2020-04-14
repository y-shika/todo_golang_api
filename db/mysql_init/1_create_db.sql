-- データベース名を変えたら`.env.local`の`LOCAL_DB_NAME`も変える
CREATE DATABASE IF NOT EXISTS todo_golang_api;
USE todo_golang_api;

CREATE TABLE IF NOT EXISTS todos (
    -- idはuuid v4を使っているため36としている
    id varchar(36) primary key,
    title text,
    active boolean,
    detail text
);

INSERT INTO todos VALUES ("1", "Title_1", true, "Detail_1"), ("2", "Title_2", true, "Detail_2"), ("3", "Title_3", true, "Detail_3");