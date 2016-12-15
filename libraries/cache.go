package libraries

import (
    "fmt"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/cache"
    _ "github.com/astaxie/beego/cache/redis"
)

var cacheManager cache.Cache

func GetCacheManager() cache.Cache {
    return cacheManager
}

func init() {
    key := beego.AppConfig.String("cache_key")
    conn := beego.AppConfig.String("cache_conn")
    dbNum := beego.AppConfig.String("cache_dbNum")

    settings := fmt.Sprintf(`{"key": "%s", "conn": "%s", "dbNum": "%s"}`, key, conn, dbNum)

    cache, err := cache.NewCache("memory", settings)

    if err != nil {
        panic(err)
    }

    cacheManager = cache
}
