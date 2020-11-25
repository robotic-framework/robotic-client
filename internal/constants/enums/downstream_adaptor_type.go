package enums

//go:generate eden generate enum --type-name=DownstreamAdaptorType
// api:enum
type DownstreamAdaptorType uint8

// 下位机驱动类型
const (
	DOWNSTREAM_ADAPTOR_TYPE_UNKNOWN  DownstreamAdaptorType = iota
	DOWNSTREAM_ADAPTOR_TYPE__FIRMATA                       // Firmata Protocol
	DOWNSTREAM_ADAPTOR_TYPE__MSP                           // MultiWii Serial Protocol
)
