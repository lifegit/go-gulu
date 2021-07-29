/**
* @Author: TheLife
* @Date: 2021/7/19 下午4:55
 */
package main_test

import (
	"github.com/lifegit/go-gulu/v2/pkg/ginplay/app"
	"testing"
)

func TestRun(t *testing.T) {
	play := app.GinPlay{
		AppPkg: "go-saletoday",
		AppDir: "./go-saletoday",
		AppAddr: "127.0.0.1",
		AppPort: 8881,

		AuthTable:  "Admin",
		AuthColumn: "password",
		AuthType: app.AuthTypeMobile,

		DbType:     "mysql",
		DbUser:     "com1yema1mysq1ok",
		DbPassword: "new1pwd1yema1ok1_",
		DbAddr:     "127.0.0.1",
		DbPort:     3306,
		DbName:     "db_test",
		DbChar:     "utf8",
	}
	_, err := play.Run()
	if err != nil {
		t.Fatal(err)
	}
}