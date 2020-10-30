package ginbro

import "time"

type Ginbro struct {
	Id        uint      `gorm:"primary_key" json:"id" form:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsSuccess  bool   `json:"is_success" form:"is_success"`
	AppSecret  string `json:"app_secret" form:"app_secret"`
	AppAddr    string `json:"app_addr" form:"app_addr"`
	AppDir     string `json:"app_dir" form:"app_dir"`
	AppPkg     string `json:"app_pkg" form:"app_pkg"`
	AuthTable  string `json:"auth_table" form:"auth_table"`
	AuthColumn string `json:"auth_column" form:"auth_column"`
	DbUser     string `json:"db_user" form:"db_user"`
	DbPassword string `json:"db_password" form:"db_password"`
	DbAddr     string `json:"db_addr" form:"db_addr"`
	DbName     string `json:"db_name" form:"db_name"`
	DbChar     string `json:"db_char" form:"db_char"`
	DbType     string `json:"db_type" form:"db_type"`
}