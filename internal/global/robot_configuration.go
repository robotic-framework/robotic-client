package global

import (
	"github.com/profzone/envconfig"
	"github.com/robotic-framework/robotic-client/internal/constants/enums"
)

var RobotConfig = struct {
	Swarm RobotConfiguration
}{}

type RobotConfiguration struct {
	// 机器人名称
	Name string `json:"name"`
	// 机器人模式
	RobotMode enums.RobotMode `json:"robotMode"`

	// 受控端配置
	// 受控上位机连接类型
	UpstreamAdaptorType enums.UpstreamAdaptorType `json:"upstreamAdaptorType"`
	// 上位机转发连接地址
	UpstreamAdaptorProxyRemoteAddr string `json:"upstreamAdaptorProxyRemoteAddr"`
	// 上位机远程转发端口
	UpstreamAdaptorProxyRemotePort int `json:"upstreamAdaptorProxyRemotePort,string"`
	// 上位机转发重试间隔
	UpstreamAdaptorProxyRetryInterval envconfig.Duration `json:"upstreamAdaptorProxyRetryInterval"`
	// 上位机转发最大重试次数
	UpstreamAdaptorProxyRetryMaxTime int `json:"upstreamAdaptorProxyRetryMaxTime,string"`

	// 自身下位机配置
	// 自身下位机连接类型
	SelfDownstreamAdaptorType enums.DownstreamAdaptorType `json:"selfDownstreamAdaptorType"`
	// Firmata串口设备名
	SelfDownstreamAdaptorFirmataName string `json:"selfDownstreamAdaptorFirmataName"`
	// MSP串口设备名
	SelfDownstreamAdaptorMSPName string `json:"selfDownstreamAdaptorMSPName"`
	// MSP串口准备等待时间
	SelfDownstreamAdaptorMSPReadyDuration envconfig.Duration `json:"selfDownstreamAdaptorMSPReadyDuration"`
}
