package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/slb"
)

/**
流程：
1.创建slb
2.添加后端服务器
3.创建监听
4.开始监听
*/

type AliSlbClient struct {
	client *slb.Client
}

func NewSlb(regionId, accessKeyID, accessKeySecret string) (client *AliSlbClient, err error) {
	c, err := slb.NewClientWithAccessKey(regionId, accessKeyID, accessKeySecret)
	return &AliSlbClient{c}, err
}

// DeleteLoadBalancer 删除slb
func (c *AliSlbClient) DeleteLoadBalancer(balancerId string) (response *slb.DeleteLoadBalancerResponse, err error) {
	request := slb.CreateDeleteLoadBalancerRequest()
	request.Scheme = "https"
	request.LoadBalancerId = balancerId

	response, err = c.client.DeleteLoadBalancer(request)

	return
}

// CreateLoadBalancerTCPListener 创建监听
func (c *AliSlbClient) CreateLoadBalancerTCPListener(balancerId string, listenerPort int, backendServerPort int) (response *slb.CreateLoadBalancerTCPListenerResponse, err error) {
	request := slb.CreateCreateLoadBalancerTCPListenerRequest()
	request.Scheme = "https"
	request.LoadBalancerId = balancerId
	request.Bandwidth = requests.NewInteger(-1)
	request.ListenerPort = requests.NewInteger(listenerPort)
	request.BackendServerPort = requests.NewInteger(backendServerPort)

	response, err = c.client.CreateLoadBalancerTCPListener(request)

	return
}

// StartLoadBalancerListener 开始监听
func (c *AliSlbClient) StartLoadBalancerListener(balancerId string, listenerPort int) (response *slb.StartLoadBalancerListenerResponse, err error) {
	request := slb.CreateStartLoadBalancerListenerRequest()
	request.Scheme = "https"
	request.LoadBalancerId = balancerId
	request.ListenerPort = requests.NewInteger(listenerPort)

	response, err = c.client.StartLoadBalancerListener(request)

	return
}

// CreateLoadBalancer 创建slb
func (c *AliSlbClient) CreateLoadBalancer(internetChargeType string, bandwidth int) (response *slb.CreateLoadBalancerResponse, err error) {
	request := slb.CreateCreateLoadBalancerRequest()
	request.Scheme = "https"

	request.InternetChargeType = internetChargeType
	request.Bandwidth = requests.NewInteger(bandwidth)

	response, err = c.client.CreateLoadBalancer(request)

	return
}

// AddBackendServers 添加后端服务器
func (c *AliSlbClient) AddBackendServers(balancerId string, serverId string) (response *slb.AddBackendServersResponse, err error) {
	request := slb.CreateAddBackendServersRequest()
	request.Scheme = "https"

	request.LoadBalancerId = balancerId
	request.BackendServers = `[{ "ServerId": "` + serverId + `", "Weight": "100", "Type": "ecs"}]`
	response, err = c.client.AddBackendServers(request)

	return
}
