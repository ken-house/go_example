package consulClient

import (
	"fmt"

	consulApi "github.com/hashicorp/consul/api"
)

type ConsulClient interface {
	RegisterService(serviceName string, ip string, port int) error
	DeregisterService(serviceId string) error
}

type consulClient struct {
	Client *consulApi.Client
}

func NewClient(addr string) (ConsulClient, error) {
	cfg := consulApi.DefaultConfig()
	cfg.Address = addr
	cli, err := consulApi.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return &consulClient{
		Client: cli,
	}, nil
}

// RegisterService 注册服务
func (cli *consulClient) RegisterService(serviceName string, ip string, port int) error {
	// 服务健康检查
	check := &consulApi.AgentServiceCheck{
		Interval:                       "2s",
		Timeout:                        "10s",
		GRPC:                           fmt.Sprintf("%s:%d", ip, port),
		GRPCUseTLS:                     true,
		TLSSkipVerify:                  true,
		DeregisterCriticalServiceAfter: "30s", //check失败后30秒删除本服务
	}
	return cli.Client.Agent().ServiceRegister(&consulApi.AgentServiceRegistration{
		ID:      fmt.Sprintf("%s-%s-%d", serviceName, ip, port),
		Name:    serviceName,
		Tags:    []string{"my_grpc"},
		Port:    port,
		Address: ip,
		Check:   check,
	})
}

// DeregisterService 注销服务
func (cli *consulClient) DeregisterService(serviceId string) error {
	return cli.Client.Agent().ServiceDeregister(serviceId)
}
