/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package arrayconv

// string数组寻找
func StringIn(search string, arr []string) bool {
	for _, val := range arr {
		if val == search {
			return true
		}
	}
	return false
}

// int 数组寻找
func IntIn(search int, arr []int) bool {
	for _, val := range arr {
		if val == search {
			return true
		}
	}
	return false
}

// int 数组寻找
func UIntIn(search uint, arr []uint) bool {
	for _, val := range arr {
		if val == search {
			return true
		}
	}
	return false
}