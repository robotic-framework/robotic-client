package downstream_adaptor

import (
	"encoding/binary"
	"errors"
	"github.com/robotic-framework/robotic-client/internal/constants/enums"
)

const (
	HeaderStart     = byte('$')
	HeaderM         = byte('M')
	HeaderArrowReq  = byte('<')
	HeaderArrowResp = byte('>')

	Ident          = 100
	Status         = 101
	RawIMU         = 102
	Servo          = 103
	Motor          = 104
	RC             = 105
	RawGPS         = 106
	CompGPS        = 107
	Attitude       = 108
	Altitude       = 109
	Analog         = 110
	RCTuning       = 111
	PID            = 112
	BOX            = 113
	MISC           = 114
	MotorPins      = 115
	BoxNames       = 116
	PinNames       = 117
	WP             = 118
	BoxIDs         = 119
	ServoConfig    = 120
	RCRawIMU       = 121
	SetRawRC       = 200
	SetRawGPS      = 201
	SetPID         = 202
	SetBox         = 203
	SetRCTuning    = 204
	AccCalibration = 205
	MagCalibration = 206
	SetMISC        = 207
	ResetConf      = 208
	SetWP          = 209
	SwitchRCSerial = 210
	IsSerial       = 211
	DEBUG          = 254
	VTXConfig      = 88
	VtxSetConfig   = 89
	EEPROMWrite    = 250
	Reboot         = 68
)

type IdentityResp struct {
	Version    uint8
	Type       enums.RobotType
	MSPVersion uint8
	Capability uint32
}

func (p *IdentityResp) UnmarshalBinary(data []byte) error {
	if len(data) < 7 {
		return errors.New("IdentityResp data must have at least 7 bytes")
	}

	p.Version = data[0]
	p.Type = enums.RobotType(data[1])
	p.MSPVersion = data[2]
	p.Capability = binary.LittleEndian.Uint32(data[3:])

	return nil
}

type StatusResp struct {
	CycleTime     uint16
	I2CErrorCount uint16
	SensorAcc     bool
	SensorBaro    bool
	SensorMag     bool
	SensorGPS     bool
	SensorSonar   bool
	Flag          uint32
	Set           uint8
}

func (s *StatusResp) UnmarshalBinary(data []byte) error {
	if len(data) < 11 {
		return errors.New("StatusResp data must have at least 11 bytes")
	}

	s.CycleTime = binary.LittleEndian.Uint16(data[0:2])
	s.I2CErrorCount = binary.LittleEndian.Uint16(data[2:4])
	sensor := binary.LittleEndian.Uint16(data[4:6])
	s.Flag = binary.LittleEndian.Uint32(data[6:10])
	s.Set = data[10]

	s.SensorAcc = uin16ToBool(sensor & 0x01)
	s.SensorBaro = uin16ToBool(sensor & 0x02)
	s.SensorMag = uin16ToBool(sensor & 0x04)
	s.SensorGPS = uin16ToBool(sensor & 0x08)
	s.SensorSonar = uin16ToBool(sensor & 0x10)

	return nil
}

type PIDResp struct {
	RollP  uint8
	RollI  uint8
	RollD  uint8
	PitchP uint8
	PitchI uint8
	PitchD uint8
	YawP   uint8
	YawI   uint8
	YawD   uint8
}

func (p *PIDResp) UnmarshalBinary(data []byte) error {
	// the valid data has 30 bytes
	if len(data) < 30 {
		return errors.New("PIDResp data must have at least 30 bytes")
	}

	p.RollP, p.RollI, p.RollD = data[0], data[1], data[2]
	p.PitchP, p.PitchI, p.PitchD = data[3], data[4], data[5]
	p.YawP, p.YawI, p.YawD = data[6], data[7], data[8]

	return nil
}

type AttitudeResp struct {
	AngleX  int16
	AngleY  int16
	Heading int16
}

func (a *AttitudeResp) UnmarshalBinary(data []byte) error {
	// the valid data has 6 bytes
	if len(data) < 6 {
		return errors.New("AttitudeResp data must have at least 6 bytes")
	}

	a.AngleX = int16(binary.LittleEndian.Uint16(data[0:2]))
	a.AngleY = int16(binary.LittleEndian.Uint16(data[2:4]))
	a.Heading = int16(binary.LittleEndian.Uint16(data[4:6]))

	return nil
}

type AltitudeResp struct {
	EstAlt int32
	Vario  int16
}

func (a *AltitudeResp) UnmarshalBinary(data []byte) error {
	// the valid data has 6 bytes
	if len(data) < 6 {
		return errors.New("AltitudeResp data must have at least 6 bytes")
	}

	a.EstAlt = int32(binary.LittleEndian.Uint32(data[0:4]))
	a.Vario = int16(binary.LittleEndian.Uint16(data[4:6]))

	return nil
}

type RawIMUResp struct {
	AccX  int16
	AccY  int16
	AccZ  int16
	GyroX int16
	GyroY int16
	GyroZ int16
	MagX  int16
	MagY  int16
	MagZ  int16
}

func (r *RawIMUResp) UnmarshalBinary(data []byte) error {
	// the valid data has 18 bytes
	if len(data) < 18 {
		return errors.New("RawIMUResp data must have at least 18 bytes")
	}

	r.AccX = int16(binary.LittleEndian.Uint16(data[0:2]))
	r.AccY = int16(binary.LittleEndian.Uint16(data[2:4]))
	r.AccZ = int16(binary.LittleEndian.Uint16(data[4:6]))

	r.GyroX = int16(binary.LittleEndian.Uint16(data[6:8]))
	r.GyroY = int16(binary.LittleEndian.Uint16(data[8:10]))
	r.GyroZ = int16(binary.LittleEndian.Uint16(data[10:12]))

	r.MagX = int16(binary.LittleEndian.Uint16(data[12:14]))
	r.MagY = int16(binary.LittleEndian.Uint16(data[14:16]))
	r.MagZ = int16(binary.LittleEndian.Uint16(data[16:18]))

	return nil
}

type ServoResp struct {
	Servos [8]int16
}

func (s *ServoResp) UnmarshalBinary(data []byte) error {
	if len(data) < 16 {
		return errors.New("ServoResp data must have at least 16 bytes")
	}

	for i := 0; i < 16; i += 2 {
		s.Servos[i/2] = int16(binary.LittleEndian.Uint16(data[i : i+2]))
	}

	return nil
}

type servoConfig struct {
	Min    int16
	Max    int16
	Middle int16
	Rate   int8
}

func (s *servoConfig) UnmarshalBinary(data []byte) error {
	if len(data) < 7 {
		return errors.New("servoConfig data must have at least 7 bytes")
	}

	s.Min = int16(binary.LittleEndian.Uint16(data[0:2]))
	s.Max = int16(binary.LittleEndian.Uint16(data[2:4]))
	s.Middle = int16(binary.LittleEndian.Uint16(data[4:6]))
	s.Rate = int8(data[6])
	return nil
}

type ServoConfigResp struct {
	ServoConfigs [8]servoConfig
}

func (s *ServoConfigResp) UnmarshalBinary(data []byte) error {
	if len(data) < 56 {
		return errors.New("ServoConfigResp data must have at least 56 bytes")
	}

	for i := 0; i < 56; i += 7 {
		err := s.ServoConfigs[i/7].UnmarshalBinary(data[i : i+7])
		if err != nil {
			return err
		}
	}

	return nil
}

type MotorResp struct {
	Motors [8]int16
}

func (s *MotorResp) UnmarshalBinary(data []byte) error {
	if len(data) < 16 {
		return errors.New("MotorResp data must have at least 16 bytes")
	}

	for i := 0; i < 16; i += 2 {
		s.Motors[i/2] = int16(binary.LittleEndian.Uint16(data[i : i+2]))
	}

	return nil
}

type MotorPinResp struct {
	Pin [8]uint8
}

func (m *MotorPinResp) UnmarshalBinary(data []byte) error {
	if len(data) < 8 {
		return errors.New("MotorPinResp data must have at least 8 bytes")
	}

	for i := 0; i < 8; i++ {
		m.Pin[i] = data[i]
	}

	return nil
}

func uin16ToBool(i uint16) bool {
	return i != 0
}
