package main

import (
    "github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Users_20161213_071813 struct {
    migration.Migration
}

// DO NOT MODIFY
func init() {
    m := &Users_20161213_071813{}
    m.Created = "20161213_071813"
    migration.Register("Users_20161213_071813", m)
}

// Run the migrations
func (m *Users_20161213_071813) Up() {
    // use m.SQL("CREATE TABLE ...") to make schema update
    m.SQL("CREATE TABLE users(id serial primary key,mail_address TEXT NOT NULL,nickname TEXT NOT NULL,password TEXT NOT NULL,created TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),updated TIMESTAMP WITHOUT TIME ZONE,deleted TIMESTAMP WITHOUT TIME ZONE)")
}

// Reverse the migrations
func (m *Users_20161213_071813) Down() {
    // use m.SQL("DROP TABLE ...") to reverse schema update
    m.SQL("DROP TABLE users")
}
