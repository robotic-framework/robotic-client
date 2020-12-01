package downstream_adaptor

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
	RCRawIMU       = 121
	SetRawRC       = 200
	SetRawGPS      = 201
	SetPID         = 202
	SetBox         = 203
	SetRCTuning    = 204
	ACCCalibration = 205
	MAGCalibration = 206
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

type PIDConfig struct {
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

func (p *PIDConfig) UnmarshalBinary(data []byte) error {
	// the valid data has 30 bytes
	_ = data[29]

	p.RollP, p.RollI, p.RollD = data[0], data[1], data[2]
	p.PitchP, p.PitchI, p.PitchD = data[3], data[4], data[5]
	p.YawP, p.YawI, p.YawD = data[6], data[7], data[8]

	return nil
}
