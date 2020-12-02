package enums

import (
	"bytes"
	"encoding"
	"errors"

	github_com_eden_framework_enumeration "github.com/eden-framework/enumeration"
)

var InvalidRobotType = errors.New("invalid RobotType")

func init() {
	github_com_eden_framework_enumeration.RegisterEnums("RobotType", map[string]string{
		"SINGLECOPTER":  "SINGLECOPTER",
		"DUALCOPTER":    "DUALCOPTER",
		"HEX6H":         "HEX6H",
		"VTAIL4":        "VTAIL4",
		"HELI_90_DEG":   "HELI_90_DEG",
		"HELI_120_CCPM": "HELI_120_CCPM",
		"AIRPLANE":      "AIRPLANE",
		"OCTOFLATX":     "OCTOFLATX",
		"OCTOFLATP":     "OCTOFLATP",
		"OCTOX8":        "OCTOX8",
		"HEX6X":         "HEX6X",
		"Y4":            "Y4",
		"FLYING_WING":   "FLYING_WING",
		"HEX6":          "HEX6",
		"Y6":            "Y6",
		"GIMBAL":        "GIMBAL",
		"BI":            "BI",
		"QUADX":         "QUADX",
		"QUADP":         "QUADP",
		"TRI":           "TRI",
	})
}

func ParseRobotTypeFromString(s string) (RobotType, error) {
	switch s {
	case "":
		return ROBOT_TYPE_UNKNOWN, nil
	case "SINGLECOPTER":
		return ROBOT_TYPE__SINGLECOPTER, nil
	case "DUALCOPTER":
		return ROBOT_TYPE__DUALCOPTER, nil
	case "HEX6H":
		return ROBOT_TYPE__HEX6H, nil
	case "VTAIL4":
		return ROBOT_TYPE__VTAIL4, nil
	case "HELI_90_DEG":
		return ROBOT_TYPE__HELI_90_DEG, nil
	case "HELI_120_CCPM":
		return ROBOT_TYPE__HELI_120_CCPM, nil
	case "AIRPLANE":
		return ROBOT_TYPE__AIRPLANE, nil
	case "OCTOFLATX":
		return ROBOT_TYPE__OCTOFLATX, nil
	case "OCTOFLATP":
		return ROBOT_TYPE__OCTOFLATP, nil
	case "OCTOX8":
		return ROBOT_TYPE__OCTOX8, nil
	case "HEX6X":
		return ROBOT_TYPE__HEX6X, nil
	case "Y4":
		return ROBOT_TYPE__Y4, nil
	case "FLYING_WING":
		return ROBOT_TYPE__FLYING_WING, nil
	case "HEX6":
		return ROBOT_TYPE__HEX6, nil
	case "Y6":
		return ROBOT_TYPE__Y6, nil
	case "GIMBAL":
		return ROBOT_TYPE__GIMBAL, nil
	case "BI":
		return ROBOT_TYPE__BI, nil
	case "QUADX":
		return ROBOT_TYPE__QUADX, nil
	case "QUADP":
		return ROBOT_TYPE__QUADP, nil
	case "TRI":
		return ROBOT_TYPE__TRI, nil
	}
	return ROBOT_TYPE_UNKNOWN, InvalidRobotType
}

func ParseRobotTypeFromLabelString(s string) (RobotType, error) {
	switch s {
	case "":
		return ROBOT_TYPE_UNKNOWN, nil
	case "SINGLECOPTER":
		return ROBOT_TYPE__SINGLECOPTER, nil
	case "DUALCOPTER":
		return ROBOT_TYPE__DUALCOPTER, nil
	case "HEX6H":
		return ROBOT_TYPE__HEX6H, nil
	case "VTAIL4":
		return ROBOT_TYPE__VTAIL4, nil
	case "HELI_90_DEG":
		return ROBOT_TYPE__HELI_90_DEG, nil
	case "HELI_120_CCPM":
		return ROBOT_TYPE__HELI_120_CCPM, nil
	case "AIRPLANE":
		return ROBOT_TYPE__AIRPLANE, nil
	case "OCTOFLATX":
		return ROBOT_TYPE__OCTOFLATX, nil
	case "OCTOFLATP":
		return ROBOT_TYPE__OCTOFLATP, nil
	case "OCTOX8":
		return ROBOT_TYPE__OCTOX8, nil
	case "HEX6X":
		return ROBOT_TYPE__HEX6X, nil
	case "Y4":
		return ROBOT_TYPE__Y4, nil
	case "FLYING_WING":
		return ROBOT_TYPE__FLYING_WING, nil
	case "HEX6":
		return ROBOT_TYPE__HEX6, nil
	case "Y6":
		return ROBOT_TYPE__Y6, nil
	case "GIMBAL":
		return ROBOT_TYPE__GIMBAL, nil
	case "BI":
		return ROBOT_TYPE__BI, nil
	case "QUADX":
		return ROBOT_TYPE__QUADX, nil
	case "QUADP":
		return ROBOT_TYPE__QUADP, nil
	case "TRI":
		return ROBOT_TYPE__TRI, nil
	}
	return ROBOT_TYPE_UNKNOWN, InvalidRobotType
}

func (RobotType) EnumType() string {
	return "RobotType"
}

func (RobotType) Enums() map[int][]string {
	return map[int][]string{
		int(ROBOT_TYPE__SINGLECOPTER):  {"SINGLECOPTER", "SINGLECOPTER"},
		int(ROBOT_TYPE__DUALCOPTER):    {"DUALCOPTER", "DUALCOPTER"},
		int(ROBOT_TYPE__HEX6H):         {"HEX6H", "HEX6H"},
		int(ROBOT_TYPE__VTAIL4):        {"VTAIL4", "VTAIL4"},
		int(ROBOT_TYPE__HELI_90_DEG):   {"HELI_90_DEG", "HELI_90_DEG"},
		int(ROBOT_TYPE__HELI_120_CCPM): {"HELI_120_CCPM", "HELI_120_CCPM"},
		int(ROBOT_TYPE__AIRPLANE):      {"AIRPLANE", "AIRPLANE"},
		int(ROBOT_TYPE__OCTOFLATX):     {"OCTOFLATX", "OCTOFLATX"},
		int(ROBOT_TYPE__OCTOFLATP):     {"OCTOFLATP", "OCTOFLATP"},
		int(ROBOT_TYPE__OCTOX8):        {"OCTOX8", "OCTOX8"},
		int(ROBOT_TYPE__HEX6X):         {"HEX6X", "HEX6X"},
		int(ROBOT_TYPE__Y4):            {"Y4", "Y4"},
		int(ROBOT_TYPE__FLYING_WING):   {"FLYING_WING", "FLYING_WING"},
		int(ROBOT_TYPE__HEX6):          {"HEX6", "HEX6"},
		int(ROBOT_TYPE__Y6):            {"Y6", "Y6"},
		int(ROBOT_TYPE__GIMBAL):        {"GIMBAL", "GIMBAL"},
		int(ROBOT_TYPE__BI):            {"BI", "BI"},
		int(ROBOT_TYPE__QUADX):         {"QUADX", "QUADX"},
		int(ROBOT_TYPE__QUADP):         {"QUADP", "QUADP"},
		int(ROBOT_TYPE__TRI):           {"TRI", "TRI"},
	}
}

func (v RobotType) String() string {
	switch v {
	case ROBOT_TYPE_UNKNOWN:
		return ""
	case ROBOT_TYPE__SINGLECOPTER:
		return "SINGLECOPTER"
	case ROBOT_TYPE__DUALCOPTER:
		return "DUALCOPTER"
	case ROBOT_TYPE__HEX6H:
		return "HEX6H"
	case ROBOT_TYPE__VTAIL4:
		return "VTAIL4"
	case ROBOT_TYPE__HELI_90_DEG:
		return "HELI_90_DEG"
	case ROBOT_TYPE__HELI_120_CCPM:
		return "HELI_120_CCPM"
	case ROBOT_TYPE__AIRPLANE:
		return "AIRPLANE"
	case ROBOT_TYPE__OCTOFLATX:
		return "OCTOFLATX"
	case ROBOT_TYPE__OCTOFLATP:
		return "OCTOFLATP"
	case ROBOT_TYPE__OCTOX8:
		return "OCTOX8"
	case ROBOT_TYPE__HEX6X:
		return "HEX6X"
	case ROBOT_TYPE__Y4:
		return "Y4"
	case ROBOT_TYPE__FLYING_WING:
		return "FLYING_WING"
	case ROBOT_TYPE__HEX6:
		return "HEX6"
	case ROBOT_TYPE__Y6:
		return "Y6"
	case ROBOT_TYPE__GIMBAL:
		return "GIMBAL"
	case ROBOT_TYPE__BI:
		return "BI"
	case ROBOT_TYPE__QUADX:
		return "QUADX"
	case ROBOT_TYPE__QUADP:
		return "QUADP"
	case ROBOT_TYPE__TRI:
		return "TRI"
	}
	return "UNKNOWN"
}

func (v RobotType) Label() string {
	switch v {
	case ROBOT_TYPE_UNKNOWN:
		return ""
	case ROBOT_TYPE__SINGLECOPTER:
		return "SINGLECOPTER"
	case ROBOT_TYPE__DUALCOPTER:
		return "DUALCOPTER"
	case ROBOT_TYPE__HEX6H:
		return "HEX6H"
	case ROBOT_TYPE__VTAIL4:
		return "VTAIL4"
	case ROBOT_TYPE__HELI_90_DEG:
		return "HELI_90_DEG"
	case ROBOT_TYPE__HELI_120_CCPM:
		return "HELI_120_CCPM"
	case ROBOT_TYPE__AIRPLANE:
		return "AIRPLANE"
	case ROBOT_TYPE__OCTOFLATX:
		return "OCTOFLATX"
	case ROBOT_TYPE__OCTOFLATP:
		return "OCTOFLATP"
	case ROBOT_TYPE__OCTOX8:
		return "OCTOX8"
	case ROBOT_TYPE__HEX6X:
		return "HEX6X"
	case ROBOT_TYPE__Y4:
		return "Y4"
	case ROBOT_TYPE__FLYING_WING:
		return "FLYING_WING"
	case ROBOT_TYPE__HEX6:
		return "HEX6"
	case ROBOT_TYPE__Y6:
		return "Y6"
	case ROBOT_TYPE__GIMBAL:
		return "GIMBAL"
	case ROBOT_TYPE__BI:
		return "BI"
	case ROBOT_TYPE__QUADX:
		return "QUADX"
	case ROBOT_TYPE__QUADP:
		return "QUADP"
	case ROBOT_TYPE__TRI:
		return "TRI"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*RobotType)(nil)

func (v RobotType) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidRobotType
	}
	return []byte(str), nil
}

func (v *RobotType) UnmarshalText(data []byte) (err error) {
	*v, err = ParseRobotTypeFromString(string(bytes.ToUpper(data)))
	return
}
