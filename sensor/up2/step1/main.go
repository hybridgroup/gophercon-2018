package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/upboard/up2"
)

func main() {
	board := up2.NewAdaptor()
	gp := i2c.NewGrovePiDriver(board)

	// digital devices
	led := gpio.NewLedDriver(gp, "D7")

	work := func() {
		gobot.Every(1*time.Second, func() {
			led.Toggle()
		})
	}

	robot := gobot.NewRobot("sensors",
		[]gobot.Connection{board},
		[]gobot.Device{gp, led},
		work,
	)

	robot.Start()
}
