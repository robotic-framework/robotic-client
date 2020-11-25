package enums

//go:generate eden generate enum --type-name=UpstreamAdaptorType
// api:enum
type UpstreamAdaptorType uint8

// 受控上位机连接类型
const (
	UPSTREAM_ADAPTOR_TYPE_UNKNOWN UpstreamAdaptorType = iota
	UPSTREAM_ADAPTOR_TYPE__PROXY                      // Proxy
	UPSTREAM_ADAPTOR_TYPE__BLE                        // Bluetooth
	UPSTREAM_ADAPTOR_TYPE__WIFI                       // Wifi
)
