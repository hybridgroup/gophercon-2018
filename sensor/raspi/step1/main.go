package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	board := raspi.NewAdaptor()
	gp := i2c.NewGrovePiDriver(board)

	// digital devices
	led := gpio.NewLedDriver(gp, "D2")

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
