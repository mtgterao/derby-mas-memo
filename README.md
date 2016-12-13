# ダビマスメモアプリ

## はじめに

beego使っているので、はじめに以下のとおりgo getしてください

```
$ go get github.com/astaxie/beego
$ go get github.com/beego/bee
$ go get github.com/astaxie/beego/cache
$ go get github.com/astaxie/beego/cache/redis
```

上記の通り、cacheとしてredisを使う想定です。

libraries/cache.goにキャッシュ接続部分は実装してあります。

また、redisへの接続設定周りは、conf/app.confにcache_として切ってあります。
