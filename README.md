# derby-mas-memo

## はじめに

beego使っているので、はじめに以下のとおりgo getしてください

```
$ go get github.com/astaxie/beego
$ go get github.com/beego/bee
$ go get github.com/astaxie/beego/cache
$ go get github.com/astaxie/beego/cache/redis
$ go get github.com/astaxie/beego/orm
$ go get github.com/lib/pq
```

上記の通り、cacheとしてredisを使う想定です。

libraries/cache.goにキャッシュ接続部分は実装してあります。

また、redisへの接続設定周りは、conf/app.confにcache_として切ってあります。

## ヘルパー関数

独自にhelpersディレクトリを切ってあり、form_helpers.goを用意しているので、フォームヘルパーはここに追記お願いします
