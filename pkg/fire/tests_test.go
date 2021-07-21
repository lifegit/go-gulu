package fire_test

import (
	"context"
	"fmt"
	"github.com/lifegit/go-gulu/v2/pkg/fire"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// init mysql
var DB *fire.Fire
var DBDryRun *fire.Fire

// 测试的开始位置
func init() {
	if db, err := OpenTestConnection(); err != nil {
		log.Printf("failed to connect database, got error %v", err)
		os.Exit(1)
	} else {
		DB = fire.NewInstance(db.Session(&gorm.Session{Logger: &(Diary{})}))
		DBDryRun = fire.NewInstance(DB.Session(&gorm.Session{DryRun: true, NewDB: true, Logger: &(Diary{})}))
		sqlDB, err := db.DB()
		if err == nil {
			err = sqlDB.Ping()
		}

		if err != nil {
			log.Printf("failed to connect database, got error %v", err)
		}

		RunMigrations()
		if DB.Dialector.Name() == "sqlite" {
			DB.Exec("PRAGMA foreign_keys = ON")
		}

		// ConnPool db conn pool
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
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
			dbDSN = "root:@tcp(localhost)/gorm?charset=utf8&parseTime=True&loc=Local"
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
	p := fire.If(position == nil, []int{1}, position).([]int)[0]
	if len(n.Sql) >= p {
		return n.Sql[len(n.Sql)-p]
	}

	return ""
}

func RunMigrations() {
	type TbUser struct {
		User
		Card      Card
		Company   Company
		Languages []*Language `gorm:"many2many:user_languages;"`
	}

	var err error
	allModels := []interface{}{&User{}, &Card{}, &Company{}, &Language{}, &TbUser{}}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(allModels), func(i, j int) { allModels[i], allModels[j] = allModels[j], allModels[i] })

	if err = DB.Migrator().DropTable(append(allModels, &UserLanguages{})...); err != nil { // UserLanguages 链接表是自动生成的，所以这里需要删除一下
		log.Printf("Failed to drop table, got error %v\n", err)
		os.Exit(1)
	}

	if err = DB.AutoMigrate(allModels...); err != nil {
		log.Printf("Failed to auto migrate, but got error %v\n", err)
		os.Exit(1)
	}

	for _, m := range allModels {
		if !DB.Migrator().HasTable(m) {
			log.Printf("Failed to create table for %#v\n", m)
			os.Exit(1)
		}
	}
	users := []TbUser{
		{
			User: User{
				Name: "Wang", Tag: "student", Age: 18, Height: 185,
			},
			Card: Card{
				Number: 1,
			},
			Company: Company{
				Address: "Shanghai", Name: "dong",
			},
			Languages: []*Language{
				{Name: "ZH"},
				{Name: "EN"},
			},
		},
		{
			User: User{
				Name: "Zhang", Tag: "student", Age: 20, Height: 180,
			},
			Card: Card{
				Number: 2,
			},
			Company: Company{
				Address: "Shanghai", Name: "dong",
			},
			Languages: []*Language{
				{Name: "ZH"},
			},
		},
		{
			User: User{
				Name: "Li", Tag: "teacher", Age: 30, Height: 155, CompanyID: 1,
			},
			Card: Card{
				Number: 3,
			},
		},
		{
			User: User{
				Name: "Liu", Tag: "boss", Age: 35, Height: 175, CompanyID: 2,
			},
			Card: Card{
				Number: 4,
			},
		},
	}
	DB.Save(&users)

	res := []string{
		fmt.Sprintf("INSERT INTO `company` (`created_at`,`updated_at`,`deleted_at`,`address`,`name`) VALUES (%d,%d,'0','Shanghai','dong'),(%d,%d,'0','Shanghai','dong') ON DUPLICATE KEY UPDATE `id`=`id`", users[0].Company.CreatedAt, users[0].Company.UpdatedAt, users[1].Company.CreatedAt, users[1].Company.UpdatedAt),
		"INSERT INTO `card` (`user_id`,`number`) VALUES (1,1),(2,2),(3,3),(4,4) ON DUPLICATE KEY UPDATE `user_id`=VALUES(`user_id`)",
		"INSERT INTO `language` (`name`) VALUES ('ZH'),('EN'),('ZH') ON DUPLICATE KEY UPDATE `id`=`id`",
		"INSERT INTO `user_languages` (`user_id`,`language_id`) VALUES (1,1),(1,2),(2,3) ON DUPLICATE KEY UPDATE `user_id`=`user_id`",
		"INSERT INTO `user` (`company_id`,`name`,`tag`,`age`,`height`) VALUES (1,'Wang','student',18,185),(2,'Zhang','student',20,180),(1,'Li','teacher',30,155),(2,'Liu','boss',35,175) ON DUPLICATE KEY UPDATE `company_id`=VALUES(`company_id`),`name`=VALUES(`name`),`tag`=VALUES(`tag`),`age`=VALUES(`age`),`height`=VALUES(`height`)",
	}
	for key, item := range res {
		if sql := DB.Logger.(*Diary).LastSql(len(res) - key); sql != item {
			log.Printf("Failed to create data for `%s` != `%s`\n", sql, item)
			os.Exit(1)
		}
	}
}

// user
type User struct {
	ID        uint `gorm:"primarykey"`
	CompanyID int
	Name      string `gormCreate:"required"`
	Tag       string
	Age       int
	Height    int
}

// company
type Company struct {
	fire.TimeFieldsModel
	ID      uint `gorm:"primarykey"`
	Address string
	Name    string
}

// card
type Card struct {
	ID     uint `gorm:"primarykey"`
	UserID uint
	Number int
}

// language
type Language struct {
	ID   uint `gorm:"primarykey"`
	Name string
}

//当使用 GORM 的 AutoMigrate 为 User 创建表时，GORM 会自动创建连接表。这里只是展示
//user_language
type UserLanguages struct {
	LanguageId uint
	UserId     uint
}
