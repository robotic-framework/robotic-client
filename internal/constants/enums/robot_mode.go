package enums

//go:generate eden generate enum --type-name=RobotMode
// api:enum
type RobotMode uint8

// 行动模式
const (
	ROBOT_MODE_UNKNOWN      RobotMode = iota
	ROBOT_MODE__MANUAL                // 手动控制
	ROBOT_MODE__AUTO                  // 全自动
	ROBOT_MODE__MANUAL_AUTO           // 自动与手动混合
	ROBOT_MODE__TEST                  // 测试
)
