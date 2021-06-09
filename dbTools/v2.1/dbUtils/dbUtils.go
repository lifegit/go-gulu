/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package dbUtils

//type DbUtils struct {
//	Db      *gorm.DB `json:"-"`
//	Wheres  *DbWheres
//	Updates *DbUpdates
//	Orders  *DbOrder
//	Joins   *DbJoin
//	Offset  *DbOffset
//}
//
//func InitDb(utils *DbUtils, defaultDb *gorm.DB) *gorm.DB {
//	if utils != nil && utils.Db != nil {
//		return utils.Db
//	} else {
//		return defaultDb
//	}
//}
//
//type T struct {
//	TableName string
//	Fields    []Field
//}
//
//type Field struct {
//	Name   string
//	AsName string
//}
//
//func FieldTab(tab ...T) []string {
//	var res []string
//	for _, item := range tab {
//		for _, val := range item.Fields {
//			if val.AsName == "" {
//				res = append(res, fmt.Sprintf("`%s`.%s", item.TableName, val.Name))
//			} else {
//				res = append(res, fmt.Sprintf("`%s`.%s as `%s`", item.TableName, val.Name, val.AsName))
//			}
//		}
//	}
//
//	return res
//}
