# todo_backend

Backend入門のプロジェクト

TODOアプリを作りながらBackendを習得していくことを目的としている

## 初期設定

- 環境変数の読み込み
    1. `.env`ファイルの作成
    
        `$ touch .env`

        `.env`を以下のように編集 (`#`部分を変更する)
        ```
        PORT= # ポート番号

        DB_ROOT_PASSWORD= # MYSQL_ROOT_PASSWORD
        DB_USER= # MYSQL_USER
        DB_PASSWORD=password= # MYSQL_PASSWORD
        ```

    2. 環境変数の適用
        ```sh
        $ chmod +x environment.sh
        $ ./environment.sh
        ```