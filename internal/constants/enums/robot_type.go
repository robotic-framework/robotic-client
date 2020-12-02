package enums

//go:generate eden generate enum --type-name=RobotType
// api:enum
type RobotType uint8

//
const (
	ROBOT_TYPE_UNKNOWN        RobotType = iota
	ROBOT_TYPE__TRI                     // TRI
	ROBOT_TYPE__QUADP                   // QUADP
	ROBOT_TYPE__QUADX                   // QUADX
	ROBOT_TYPE__BI                      // BI
	ROBOT_TYPE__GIMBAL                  // GIMBAL
	ROBOT_TYPE__Y6                      // Y6
	ROBOT_TYPE__HEX6                    // HEX6
	ROBOT_TYPE__FLYING_WING             // FLYING_WING
	ROBOT_TYPE__Y4                      // Y4
	ROBOT_TYPE__HEX6X                   // HEX6X
	ROBOT_TYPE__OCTOX8                  // OCTOX8
	ROBOT_TYPE__OCTOFLATP               // OCTOFLATP
	ROBOT_TYPE__OCTOFLATX               // OCTOFLATX
	ROBOT_TYPE__AIRPLANE                // AIRPLANE
	ROBOT_TYPE__HELI_120_CCPM           // HELI_120_CCPM
	ROBOT_TYPE__HELI_90_DEG             // HELI_90_DEG
	ROBOT_TYPE__VTAIL4                  // VTAIL4
	ROBOT_TYPE__HEX6H                   // HEX6H
)

const (
	ROBOT_TYPE__DUALCOPTER   RobotType = iota + 20 // DUALCOPTER
	ROBOT_TYPE__SINGLECOPTER                       // SINGLECOPTER
)
