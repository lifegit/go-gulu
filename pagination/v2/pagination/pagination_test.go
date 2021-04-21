/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package pagination

func main() {
	//"gorm.io/gorm"ar param pagination.Param
	//err := c.ShouldBind(&param)
	//if app.HandleError(c, err) { return }

	//maps:= make(map[string]interface{})
	//maps["sex"] = []int{1,2}
	//maps["nickname"] = "华"
	//param := pagination.Param{
	//	Page: 2,
	//	Filtered: &maps,
	//}
	//
	//query := pagination.InitPagination(param.Page)
	//query.AllowFiltered("tb_users_info",param.Filtered,
	//	[]string{"sex"},
	//	[]pagination.Searched{{ Name: "nickname", Vague:true} },
	//)
	//query.AllowSorted("tb_users_info",param.Sorted,
	//	[]string{"time_created","num_star"},
	//	&dbUtils.Order{Field: "id", Type: dbUtils.OrderDesc},
	//)
	//query.AddJoin("tb_users_follow",
	//	dbUtils.JoinOn{Table: "tb_users_follow", Field: "activeuid"},
	//	dbUtils.JoinOn{Table: "tb_users_info",   Field: "id"},
	//)
	//
	//
	//var list = &[]models.TbUsersInfo{}
	//
	//page, err := user.AllPage([]string{ "tb_users_info.id" }, list, query)
	//fmt.Println(page,err,list)
	//if app.HandleError(c, err) {
	//	return
	//}
	//
	//app.JsonPagination2(c,list,page)
}
