/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package dbUtils

func main() {
	//user := models.TbUsersInfo{
	//	Privilege: "山东",
	//	DbUtils: &dbUtils.DbUtils{
	//		Where:  &dbUtils.DbWheres{
	//			In: []dbUtils.In{
	//				{
	//					Field:"id",
	//					In:[]int{1,2},
	//				},
	//				{
	//					Field:"city",
	//					In:[]string{"济南","淄博"},
	//				},
	//			},
	//			Range: []dbUtils.Range{
	//				{
	//					Field: "time_created",
	//					Start: 0,
	//					End:   99,
	//				},
	//			},
	//			Like: []dbUtils.Like{
	//				{
	//					Field:"nickname",
	//					Text: "华",
	//				},
	//			},
	//		},
	//		Updates: &dbUtils.DbUpdates{
	//			Arithmetic:[]dbUtils.Arithmetic{
	//				{
	//					Field: "money_normal",
	//					Type: dbUtils.ArithmeticReduce,
	//					Number:5,
	//
	//				},
	//				{
	//					Field: "money_frozen",
	//					Type: dbUtils.ArithmeticIncrease,
	//					Number:5,
	//				},
	//			},
	//			Set: []dbUtils.Set{
	//				{
	//					Field:  "sex",
	//					Value:  1,
	//				},
	//			},
	//		},
	//	},
	//}

	//b := user.IsExists()
	//fmt.Println("isExists",b)
	//
	//one,err := user.One(nil)
	//fmt.Println("one",one,err)

	//list,err := user.All(nil)
	//fmt.Println("list",list,err)
	//
	//user.Id = one.Id
	//err = user.Update()
	//fmt.Println("update",err)

	//err = user.Delete()
	//fmt.Println("delete",err)
}
