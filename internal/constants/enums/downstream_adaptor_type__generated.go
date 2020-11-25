package enums

import (
	"bytes"
	"encoding"
	"errors"

	github_com_eden_framework_enumeration "github.com/eden-framework/enumeration"
)

var InvalidDownstreamAdaptorType = errors.New("invalid DownstreamAdaptorType")

func init() {
	github_com_eden_framework_enumeration.RegisterEnums("DownstreamAdaptorType", map[string]string{
		"MSP":     "MultiWii Serial Protocol",
		"FIRMATA": "Firmata Protocol",
	})
}

func ParseDownstreamAdaptorTypeFromString(s string) (DownstreamAdaptorType, error) {
	switch s {
	case "":
		return DOWNSTREAM_ADAPTOR_TYPE_UNKNOWN, nil
	case "MSP":
		return DOWNSTREAM_ADAPTOR_TYPE__MSP, nil
	case "FIRMATA":
		return DOWNSTREAM_ADAPTOR_TYPE__FIRMATA, nil
	}
	return DOWNSTREAM_ADAPTOR_TYPE_UNKNOWN, InvalidDownstreamAdaptorType
}

func ParseDownstreamAdaptorTypeFromLabelString(s string) (DownstreamAdaptorType, error) {
	switch s {
	case "":
		return DOWNSTREAM_ADAPTOR_TYPE_UNKNOWN, nil
	case "MultiWii Serial Protocol":
		return DOWNSTREAM_ADAPTOR_TYPE__MSP, nil
	case "Firmata Protocol":
		return DOWNSTREAM_ADAPTOR_TYPE__FIRMATA, nil
	}
	return DOWNSTREAM_ADAPTOR_TYPE_UNKNOWN, InvalidDownstreamAdaptorType
}

func (DownstreamAdaptorType) EnumType() string {
	return "DownstreamAdaptorType"
}

func (DownstreamAdaptorType) Enums() map[int][]string {
	return map[int][]string{
		int(DOWNSTREAM_ADAPTOR_TYPE__MSP):     {"MSP", "MultiWii Serial Protocol"},
		int(DOWNSTREAM_ADAPTOR_TYPE__FIRMATA): {"FIRMATA", "Firmata Protocol"},
	}
}

func (v DownstreamAdaptorType) String() string {
	switch v {
	case DOWNSTREAM_ADAPTOR_TYPE_UNKNOWN:
		return ""
	case DOWNSTREAM_ADAPTOR_TYPE__MSP:
		return "MSP"
	case DOWNSTREAM_ADAPTOR_TYPE__FIRMATA:
		return "FIRMATA"
	}
	return "UNKNOWN"
}

func (v DownstreamAdaptorType) Label() string {
	switch v {
	case DOWNSTREAM_ADAPTOR_TYPE_UNKNOWN:
		return ""
	case DOWNSTREAM_ADAPTOR_TYPE__MSP:
		return "MultiWii Serial Protocol"
	case DOWNSTREAM_ADAPTOR_TYPE__FIRMATA:
		return "Firmata Protocol"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*DownstreamAdaptorType)(nil)

func (v DownstreamAdaptorType) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidDownstreamAdaptorType
	}
	return []byte(str), nil
}

func (v *DownstreamAdaptorType) UnmarshalText(data []byte) (err error) {
	*v, err = ParseDownstreamAdaptorTypeFromString(string(bytes.ToUpper(data)))
	return
}
