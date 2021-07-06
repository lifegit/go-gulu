package ginbro

import (
	"testing"
)

// sudo go test -v ./ginbro && sudo chmod -R 777 /Users/yxs/GolandProjects/src/felix/ginbro/TEST0
func TestRun(t *testing.T) {

	gc := Ginbro{
		AppAddr:    "127.0.0.1:4444",

		AppDir:     "./go-saletoday",
		AppPkg:     "go-saletoday",
		AuthTable:  "users",
		AuthColumn: "password",

		DbUser:     "com1yema1mysq1ok",
		DbPassword: "new1pwd1yema1ok1_",
		DbAddr:     "127.0.0.1:3306",
		DbType:     "mysql",
		DbName:     "db_test",
		DbChar:     "utf8",
	}

	_, err := Run(gc)
	if err != nil {
		t.Fatal(err)
	}
}
