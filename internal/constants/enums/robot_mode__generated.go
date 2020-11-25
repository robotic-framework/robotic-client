package enums

import (
	"bytes"
	"encoding"
	"errors"

	github_com_eden_framework_enumeration "github.com/eden-framework/enumeration"
)

var InvalidRobotMode = errors.New("invalid RobotMode")

func init() {
	github_com_eden_framework_enumeration.RegisterEnums("RobotMode", map[string]string{
		"TEST":        "测试",
		"MANUAL_AUTO": "自动与手动混合",
		"AUTO":        "全自动",
		"MANUAL":      "手动控制",
	})
}

func ParseRobotModeFromString(s string) (RobotMode, error) {
	switch s {
	case "":
		return ROBOT_MODE_UNKNOWN, nil
	case "TEST":
		return ROBOT_MODE__TEST, nil
	case "MANUAL_AUTO":
		return ROBOT_MODE__MANUAL_AUTO, nil
	case "AUTO":
		return ROBOT_MODE__AUTO, nil
	case "MANUAL":
		return ROBOT_MODE__MANUAL, nil
	}
	return ROBOT_MODE_UNKNOWN, InvalidRobotMode
}

func ParseRobotModeFromLabelString(s string) (RobotMode, error) {
	switch s {
	case "":
		return ROBOT_MODE_UNKNOWN, nil
	case "测试":
		return ROBOT_MODE__TEST, nil
	case "自动与手动混合":
		return ROBOT_MODE__MANUAL_AUTO, nil
	case "全自动":
		return ROBOT_MODE__AUTO, nil
	case "手动控制":
		return ROBOT_MODE__MANUAL, nil
	}
	return ROBOT_MODE_UNKNOWN, InvalidRobotMode
}

func (RobotMode) EnumType() string {
	return "RobotMode"
}

func (RobotMode) Enums() map[int][]string {
	return map[int][]string{
		int(ROBOT_MODE__TEST):        {"TEST", "测试"},
		int(ROBOT_MODE__MANUAL_AUTO): {"MANUAL_AUTO", "自动与手动混合"},
		int(ROBOT_MODE__AUTO):        {"AUTO", "全自动"},
		int(ROBOT_MODE__MANUAL):      {"MANUAL", "手动控制"},
	}
}

func (v RobotMode) String() string {
	switch v {
	case ROBOT_MODE_UNKNOWN:
		return ""
	case ROBOT_MODE__TEST:
		return "TEST"
	case ROBOT_MODE__MANUAL_AUTO:
		return "MANUAL_AUTO"
	case ROBOT_MODE__AUTO:
		return "AUTO"
	case ROBOT_MODE__MANUAL:
		return "MANUAL"
	}
	return "UNKNOWN"
}

func (v RobotMode) Label() string {
	switch v {
	case ROBOT_MODE_UNKNOWN:
		return ""
	case ROBOT_MODE__TEST:
		return "测试"
	case ROBOT_MODE__MANUAL_AUTO:
		return "自动与手动混合"
	case ROBOT_MODE__AUTO:
		return "全自动"
	case ROBOT_MODE__MANUAL:
		return "手动控制"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*RobotMode)(nil)

func (v RobotMode) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidRobotMode
	}
	return []byte(str), nil
}

func (v *RobotMode) UnmarshalText(data []byte) (err error) {
	*v, err = ParseRobotModeFromString(string(bytes.ToUpper(data)))
	return
}
