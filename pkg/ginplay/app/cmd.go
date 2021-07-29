/**
* @Author: TheLife
* @Date: 2021/7/22 下午4:18
 */
package app

import (
	"github.com/spf13/cobra"
)

var gc GinPlay

var example = `
go build -o main && ./main create \
-d ./go-admin \
-k go-admin \
-b admin \
-u root \
-s pass \
-n db_test
`

// RestCmd represents the rest command
var RestCmd = &cobra.Command{
	Use:     "create",
	Short:   "generate a RESTful codebase from SQL database",
	Long:    `generate a RESTful APIs app with gin and gorm for gophers`,
	Example: example,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		//_, err = Run(gc)
		return err
	},
}

func init() {
	RestCmd.Flags().StringVarP(&gc.AppPkg, "appPkg", "k", "", "go.mod module name")
	RestCmd.Flags().StringVarP(&gc.AppDir, "appDir", "d", "", "code project output directory")
	RestCmd.Flags().StringVarP(&gc.AppAddr, "appAddr", "a", "127.0.0.1", "http service bind address")
	RestCmd.Flags().IntVarP(&gc.AppPort, "appPort", "o", 8080, "http service bind port")
	RestCmd.Flags().StringVarP(&gc.AuthTable, "authTable", "b", "users", "login user table")
	RestCmd.Flags().StringVarP(&gc.AuthColumn, "authColumn", "l", "password", "bcrypt password column")
	RestCmd.Flags().IntVarP((*int)(&gc.AuthType), "authType", "e", int(AuthTypeMobile), "generate 1:mobile or 2:security on register、forget")
	RestCmd.Flags().StringVarP(&gc.DbType, "dbType", "t", "mysql", "database type: mysql/postgres/mssql/sqlite")
	RestCmd.Flags().StringVarP(&gc.DbUser, "dbUser", "u", "root", "database username")
	RestCmd.Flags().StringVarP(&gc.DbPassword, "dbPassword", "s", "", "database user password")
	RestCmd.Flags().StringVarP(&gc.DbAddr, "dbAddr", "r", "127.0.0.1", "database connection addr")
	RestCmd.Flags().IntVarP(&gc.DbPort, "dbPort", "p", 3306, "database connection addr")
	RestCmd.Flags().StringVarP(&gc.DbName, "dbName", "n", "", "database name")
	RestCmd.Flags().StringVarP(&gc.DbChar, "dbChar", "c", "utf8", "database charset")

	RestCmd.MarkFlagRequired("appDir")
	RestCmd.MarkFlagRequired("appPkg")
	RestCmd.MarkFlagRequired("dbName")
}
