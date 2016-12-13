package main

import (
    "fmt"
    _ "derby-mas-memo/routers"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    _ "github.com/lib/pq"
)

func init() {
    initDatabase()
    initTemplate()
}

func main() {
    beego.Run()
}

func initDatabase() {
    driver := beego.AppConfig.String("db_driver")
    dbLink := fmt.Sprintf(
        "user=%s dbname=%s sslmode=disable",
        beego.AppConfig.String("db_user"),
        beego.AppConfig.String("db_name"),
    )

    orm.RegisterDriver(driver, orm.DRPostgres)
    orm.RegisterDataBase("default", driver, dbLink)
}

func initTemplate() {
    helpers.RegistFormHelper()
}
