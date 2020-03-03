package main

import (
	"database/sql"
	"github.com/astaxie/beego/toolbox"
)

type DatabaseCheck struct {
}

func (*DatabaseCheck) Check() error {
	_, err := sql.Open("mysql", "handsome:handsome@/guess?charset=utf8")
	return err
}

func init() {
	toolbox.AddHealthCheck("database", &DatabaseCheck{})
}
