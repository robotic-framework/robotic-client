package enums

import (
	"bytes"
	"encoding"
	"errors"

	github_com_eden_framework_enumeration "github.com/eden-framework/enumeration"
)

var InvalidUpstreamAdaptorType = errors.New("invalid UpstreamAdaptorType")

func init() {
	github_com_eden_framework_enumeration.RegisterEnums("UpstreamAdaptorType", map[string]string{
		"WIFI":  "Wifi",
		"BLE":   "Bluetooth",
		"PROXY": "Proxy",
	})
}

func ParseUpstreamAdaptorTypeFromString(s string) (UpstreamAdaptorType, error) {
	switch s {
	case "":
		return UPSTREAM_ADAPTOR_TYPE_UNKNOWN, nil
	case "WIFI":
		return UPSTREAM_ADAPTOR_TYPE__WIFI, nil
	case "BLE":
		return UPSTREAM_ADAPTOR_TYPE__BLE, nil
	case "PROXY":
		return UPSTREAM_ADAPTOR_TYPE__PROXY, nil
	}
	return UPSTREAM_ADAPTOR_TYPE_UNKNOWN, InvalidUpstreamAdaptorType
}

func ParseUpstreamAdaptorTypeFromLabelString(s string) (UpstreamAdaptorType, error) {
	switch s {
	case "":
		return UPSTREAM_ADAPTOR_TYPE_UNKNOWN, nil
	case "Wifi":
		return UPSTREAM_ADAPTOR_TYPE__WIFI, nil
	case "Bluetooth":
		return UPSTREAM_ADAPTOR_TYPE__BLE, nil
	case "Proxy":
		return UPSTREAM_ADAPTOR_TYPE__PROXY, nil
	}
	return UPSTREAM_ADAPTOR_TYPE_UNKNOWN, InvalidUpstreamAdaptorType
}

func (UpstreamAdaptorType) EnumType() string {
	return "UpstreamAdaptorType"
}

func (UpstreamAdaptorType) Enums() map[int][]string {
	return map[int][]string{
		int(UPSTREAM_ADAPTOR_TYPE__WIFI):  {"WIFI", "Wifi"},
		int(UPSTREAM_ADAPTOR_TYPE__BLE):   {"BLE", "Bluetooth"},
		int(UPSTREAM_ADAPTOR_TYPE__PROXY): {"PROXY", "Proxy"},
	}
}

func (v UpstreamAdaptorType) String() string {
	switch v {
	case UPSTREAM_ADAPTOR_TYPE_UNKNOWN:
		return ""
	case UPSTREAM_ADAPTOR_TYPE__WIFI:
		return "WIFI"
	case UPSTREAM_ADAPTOR_TYPE__BLE:
		return "BLE"
	case UPSTREAM_ADAPTOR_TYPE__PROXY:
		return "PROXY"
	}
	return "UNKNOWN"
}

func (v UpstreamAdaptorType) Label() string {
	switch v {
	case UPSTREAM_ADAPTOR_TYPE_UNKNOWN:
		return ""
	case UPSTREAM_ADAPTOR_TYPE__WIFI:
		return "Wifi"
	case UPSTREAM_ADAPTOR_TYPE__BLE:
		return "Bluetooth"
	case UPSTREAM_ADAPTOR_TYPE__PROXY:
		return "Proxy"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*UpstreamAdaptorType)(nil)

func (v UpstreamAdaptorType) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidUpstreamAdaptorType
	}
	return []byte(str), nil
}

func (v *UpstreamAdaptorType) UnmarshalText(data []byte) (err error) {
	*v, err = ParseUpstreamAdaptorTypeFromString(string(bytes.ToUpper(data)))
	return
}
