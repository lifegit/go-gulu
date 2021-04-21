/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package pagination

func arrayGet(search string, arr []Searched) *Searched {
	for _, val := range arr {
		if val.AsName != "" {
			if val.AsName == search {
				return &val
			}
		} else {
			if val.Name == search {
				return &val
			}
		}
	}

	return nil
}
