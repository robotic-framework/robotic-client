package enums

import (
	"bytes"
	"encoding"
	"errors"

	github_com_eden_framework_enumeration "github.com/eden-framework/enumeration"
)

var InvalidCameraMode = errors.New("invalid CameraMode")

func init() {
	github_com_eden_framework_enumeration.RegisterEnums("CameraMode", map[string]string{
		"OBJECT_DETECTION": "物体识别",
		"TRANSFER":         "视频传输",
	})
}

func ParseCameraModeFromString(s string) (CameraMode, error) {
	switch s {
	case "":
		return CAMERA_MODE_UNKNOWN, nil
	case "OBJECT_DETECTION":
		return CAMERA_MODE__OBJECT_DETECTION, nil
	case "TRANSFER":
		return CAMERA_MODE__TRANSFER, nil
	}
	return CAMERA_MODE_UNKNOWN, InvalidCameraMode
}

func ParseCameraModeFromLabelString(s string) (CameraMode, error) {
	switch s {
	case "":
		return CAMERA_MODE_UNKNOWN, nil
	case "物体识别":
		return CAMERA_MODE__OBJECT_DETECTION, nil
	case "视频传输":
		return CAMERA_MODE__TRANSFER, nil
	}
	return CAMERA_MODE_UNKNOWN, InvalidCameraMode
}

func (CameraMode) EnumType() string {
	return "CameraMode"
}

func (CameraMode) Enums() map[int][]string {
	return map[int][]string{
		int(CAMERA_MODE__OBJECT_DETECTION): {"OBJECT_DETECTION", "物体识别"},
		int(CAMERA_MODE__TRANSFER):         {"TRANSFER", "视频传输"},
	}
}

func (v CameraMode) String() string {
	switch v {
	case CAMERA_MODE_UNKNOWN:
		return ""
	case CAMERA_MODE__OBJECT_DETECTION:
		return "OBJECT_DETECTION"
	case CAMERA_MODE__TRANSFER:
		return "TRANSFER"
	}
	return "UNKNOWN"
}

func (v CameraMode) Label() string {
	switch v {
	case CAMERA_MODE_UNKNOWN:
		return ""
	case CAMERA_MODE__OBJECT_DETECTION:
		return "物体识别"
	case CAMERA_MODE__TRANSFER:
		return "视频传输"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*CameraMode)(nil)

func (v CameraMode) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidCameraMode
	}
	return []byte(str), nil
}

func (v *CameraMode) UnmarshalText(data []byte) (err error) {
	*v, err = ParseCameraModeFromString(string(bytes.ToUpper(data)))
	return
}
