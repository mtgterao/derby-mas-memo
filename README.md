# derby-mas-memo

## はじめに

godepを使ってパッケージ管理をしているので、godepをインストールしてください。

```
$ go get github.com/tools/godep

$ cd github.com/mtgterao/derby-mas-memo/

$ godep restore
```


cacheとしてredisを使う想定です。

libraries/cache.goにキャッシュ接続部分は実装してあります。

また、redisへの接続設定周りは、conf/app.confにcache_として切ってあります。

## ヘルパー関数

独自にhelpersディレクトリを切ってあり、form_helpers.goを用意しているので、フォームヘルパーはここに追記お願いします

## マイグレーション

マイグレーション実行時は以下の通り接続名の指定が必要です。

```
$ bee migration -driver="postgres" -conn="postgres://<pg_user>:<pg_pass>@<host>:<port>/<db_name>?sslmode=disable"
```
