/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package mobileCode

type mobileCodeInterface interface {
	Set(key, value string) error
	Get(key string) (string, error)
	Del(key string) error
}

type MobileCode struct {
	tagPrefix string
	storage   mobileCodeInterface
}

func New(storage mobileCodeInterface) *MobileCode {
	return &MobileCode{
		storage: storage,
	}
}

// 发送
func (m *MobileCode) Send(sendFunc func() (MobileMes, error)) error {
	// 发送验证码
	sendMes, err := sendFunc()
	if err != nil {
		return err
	}

	// 放到缓存
	return m.storage.Set(sendMes.Mobile, sendMes.Code)
}

// 是否正确
func (m *MobileCode) IsCheck(ms MobileMes) bool {
	str, _ := m.storage.Get(ms.Mobile)
	return str == ms.Code
}

// 是否存在
func (m *MobileCode) IsExist(mobile string) bool {
	str, _ := m.storage.Get(mobile)
	return str != ""
}

// 删除
func (m *MobileCode) Del(mobile string) error {
	return m.storage.Del(mobile)
}
