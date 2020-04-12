-- DROP DATABASE IF EXISTS todo_golang_api;
CREATE DATABASE IF NOT EXISTS todo_golang_api;
USE todo_golang_api;
-- DROP TABLE IF EXISTS todos;

CREATE TABLE IF NOT EXISTS todos (
    -- idはuuid v4を使っているため
    id varchar(36) primary key,
    title text,
    active boolean,
    detail text
);

INSERT INTO todos VALUES ("1", "Title_1", true, "Detail_1"), ("2", "Title_2", true, "Detail_2");