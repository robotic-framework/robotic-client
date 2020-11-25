package enums

//go:generate eden generate enum --type-name=CameraMode
// api:enum
type CameraMode uint8

// 摄像头模式
const (
	CAMERA_MODE_UNKNOWN           CameraMode = iota
	CAMERA_MODE__TRANSFER                    // 视频传输
	CAMERA_MODE__OBJECT_DETECTION            // 物体识别
)
