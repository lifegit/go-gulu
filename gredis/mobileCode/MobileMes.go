/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package mobileCode

type MobileMes struct {
	Mobile string
	Code   string
}

func (m *MobileMes) key(t string) string {
	return tag(t, m.Mobile)
}
