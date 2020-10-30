package ginbro

import (
	"os"
	"testing"
)

// sudo go test -v ./ginbro && sudo chmod -R 777 /Users/yxs/GolandProjects/src/felix/ginbro/TEST0
func TestRun(t *testing.T) {

	gc := Ginbro{
		AppSecret:  "sdfsadfewdddcd",
		AppDir:     "./dream-admin",
		AppAddr:    "127.0.0.1:4444",
		AppPkg:     "dream-admin",
		AuthTable:  "users",
		AuthColumn: "password",
		DbUser:     "dream_d_hsouyvd",
		DbPassword: "daiudAhgeowhf2",
		DbAddr:     "rm-bp1a2mm5kn82b8sp5o.mysql.rds.aliyuncs.com:3306",
		DbType:     "mysql",
		DbName:     "db_clouddream_development",
		DbChar:     "utf8",
	}

	_, err := Run(gc)
	if err != nil {
		t.Fatal(err)
	}
	_ = os.Chmod(gc.AppDir, 0777)
}
