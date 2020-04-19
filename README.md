# todo_golang_api

Backend入門のプロジェクト

TODOアプリを作りながらBackendを習得していくことを目的としている

## 初期設定

- 環境変数の読み込み
    - `.env.local`ファイルの作成

        環境変数を記載しているファイル

        `$ touch .env.local`

        `.env.local`を以下のように編集 (`=`の後の`#`部分を変更する)

        注) Makefileの処理で、`.env.local`にコメントなどを残すと動作しない
        ``` 
        USAGE=local
        PORT=8080
        TZ=Asia/Tokyo
        HEROKU_DB_NAME= # HerokuDBの名前
        HEROKU_DB_USERNAME= # HerokuDBのユーザ名
        HEROKU_DB_PASSWORD= # HerokuDBのユーザパスワード
        HEROKU_DB_HOST= # HerokuDBのホスティング先
        LOCAL_DB_NAME=todo_golang_api
        LOCAL_DB_HOST=mysql_db:3306
        MYSQL_ROOT_PASSWORD= # MYSQL_ROOT_PASSWORD (docker-composeの関係でこの環境変数名は固定とする)
        MYSQL_USER= # MYSQL_USER (docker-composeの関係でこの環境変数名は固定とする)
        MYSQL_PASSWORD= # MYSQL_PASSWORD (docker-composeの関係でこの環境変数名は固定とする)
        ENABLE_MOCK= # 空白: gateway, true: stub (Mock)
        CONNECTED_DATABASE= # heroku: HerokuDBに接続, local: ローカルDBに接続
        ```

    - `.env`ファイルの作成

        Herokuに適用する環境変数を定義したファイル
    
        `$ touch .env`

        `.env`を以下のように編集 (`#`部分を変更する)
        ```
        ## Usage ##
        USAGE=heroku

        HEROKU_DB_NAME= # HerokuDBの名前
        HEROKU_DB_USERNAME= # HerokuDBのユーザ名
        HEROKU_DB_PASSWORD= # HerokuDBのユーザパスワード
        HEROKU_DB_HOST= # HerokuDBのホスティング先

        ## TimeZone ##
        TZ=Asia/Tokyo
        ```