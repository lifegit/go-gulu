package fire_test

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// init mysql
var DB *DbUtils
var DBDryRun *DbUtils

// 测试的开始位置
func init() {
	if db, err := OpenTestConnection(); err != nil {
		log.Printf("failed to connect database, got error %v", err)
		os.Exit(1)
	} else {
		DB = NewInstance(db.Session(&gorm.Session{Logger: &(Diary{})}))
		DBDryRun = NewInstance(DB.Session(&gorm.Session{DryRun: true, NewDB: true, Logger: &(Diary{})}))
		sqlDB, err := db.DB()
		if err == nil {
			err = sqlDB.Ping()
		}

		if err != nil {
			log.Printf("failed to connect database, got error %v", err)
		}

		//RunMigrations()
		if DB.Dialector.Name() == "sqlite" {
			DB.Exec("PRAGMA foreign_keys = ON")
		}
	}
}

func OpenTestConnection() (db *gorm.DB, err error) {
	gormConf := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		// https://gorm.io/docs/gorm_config.html#NamingStrategy
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			NameReplacer:  strings.NewReplacer("Tb", ""),
		}, // 表名不加复数s
	}
	dbDSN := os.Getenv("GORM_DSN")
	switch os.Getenv("GORM_DIALECT") {
	case "mysql":
		log.Println("testing mysql...")
		if dbDSN == "" {
			dbDSN = "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
		}
		db, err = gorm.Open(mysql.Open(dbDSN), gormConf)
	case "postgres":
		log.Println("testing postgres...")
		if dbDSN == "" {
			dbDSN = "user=gorm password=gorm dbname=gorm host=localhost port=9920 sslmode=disable TimeZone=Asia/Shanghai"
		}
		db, err = gorm.Open(postgres.New(postgres.Config{
			DSN:                  dbDSN,
			PreferSimpleProtocol: true,
		}), gormConf)
	case "sqlserver":
		// CREATE LOGIN gorm WITH PASSWORD = 'LoremIpsum86';
		// CREATE DATABASE gorm;
		// USE gorm;
		// CREATE USER gorm FROM LOGIN gorm;
		// sp_changedbowner 'gorm';
		// npm install -g sql-cli
		// mssql -u gorm -p LoremIpsum86 -d gorm -o 9930
		log.Println("testing sqlserver...")
		if dbDSN == "" {
			dbDSN = "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
		}
		db, err = gorm.Open(sqlserver.Open(dbDSN), gormConf)
	default:
		log.Println("testing sqlite3...")
		db, err = gorm.Open(sqlite.Open(filepath.Join(os.TempDir(), "gorm.db")), gormConf)
	}

	// ConnPool db conn pool
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return
}

type Diary struct {
	Sql []string
}

func (n *Diary) LogMode(logger.LogLevel) logger.Interface {
	return &(*n)
}
func (n *Diary) Info(c context.Context, sql string, a ...interface{}) {
	fmt.Println("Info: ", sql)
}
func (n *Diary) Warn(c context.Context, sql string, a ...interface{}) {
	fmt.Println("Warn: ", sql)
}
func (n *Diary) Error(c context.Context, sql string, a ...interface{}) {
	fmt.Println("Error: ", sql)
}
func (n *Diary) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	sql, _ := fc()
	n.Sql = append(n.Sql, sql)
	fmt.Println("Trace: ", sql)
}
func (n *Diary) LastSql(position ...int) string {
	p := If(position == nil, []int{1}, position).([]int)[0]
	if len(n.Sql) >= p {
		return n.Sql[len(n.Sql)-p]
	}

	return ""
}

// todo
//func RunMigrations() {
//	var err error
//	allModels := []interface{}{&User{}, &Account{}, &Pet{}, &Company{}, &Toy{}, &Language{}}
//	rand.Seed(time.Now().UnixNano())
//	rand.Shuffle(len(allModels), func(i, j int) { allModels[i], allModels[j] = allModels[j], allModels[i] })
//
//	DB.Migrator().DropTable("user_friends", "user_speaks")
//
//	if err = DB.Migrator().DropTable(allModels...); err != nil {
//		log.Printf("Failed to drop table, got error %v\n", err)
//		os.Exit(1)
//	}
//
//	if err = DB.AutoMigrate(allModels...); err != nil {
//		log.Printf("Failed to auto migrate, but got error %v\n", err)
//		os.Exit(1)
//	}
//
//	for _, m := range allModels {
//		if !DB.Migrator().HasTable(m) {
//			log.Printf("Failed to create table for %#v\n", m)
//			os.Exit(1)
//		}
//	}
//}

// user
type TbUser struct {
	ID        uint `gorm:"primarykey"`
	CompanyID int
	Name      string
	Tag       string
	Age       int
	Height    string
}

// company
type TbCompany struct {
	TimeFieldsModel
	ID      uint `gorm:"primarykey"`
	Address string
	Name    string
}

// card
type TbCard struct {
	ID     uint `gorm:"primarykey"`
	UserID uint
	Number int
}

// language
type TbLanguage struct {
	ID   uint `gorm:"primarykey"`
	Name string
}

// 当使用 GORM 的 AutoMigrate 为 User 创建表时，GORM 会自动创建连接表。这里只是展示
// user_language
//type TbUserLanguage struct {
//	LanguageId uint
//	UserId     uint
//}
