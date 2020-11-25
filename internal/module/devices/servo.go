package devices

import "gobot.io/x/gobot/drivers/gpio"

type Servo struct {
	d *gpio.ServoDriver
}
