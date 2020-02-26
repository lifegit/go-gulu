/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package aliyun

import (
	"encoding/json"
	"errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

type AliSmsClient struct {
	client *dysmsapi.Client
}

func NewSMS(regionId, accessKeyID, accessKeySecret string) (client *AliSmsClient, err error) {
	c, err := dysmsapi.NewClientWithAccessKey(regionId, accessKeyID, accessKeySecret)
	return &AliSmsClient{c}, err
}

// 发送短信
func (c *AliSmsClient) Send(phoneNumber, signName, templateCode string, templateParam map[string]interface{}) (response *dysmsapi.SendSmsResponse, err error) {
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = phoneNumber
	request.SignName = signName
	request.TemplateCode = templateCode
	if len(templateParam) > 0 {
		mjson, _ := json.Marshal(templateParam)
		request.TemplateParam = string(mjson)
	}

	response, err = c.client.SendSms(request)
	if response != nil && response.Code != "OK" {
		err = errors.New(response.Message)
	}

	return
}
