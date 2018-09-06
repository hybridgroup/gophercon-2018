package main

import (
	"fmt"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/upboard/up2"
)

func main() {
	board := up2.NewAdaptor()
	gp := i2c.NewGrovePiDriver(board)

	button := gpio.NewButtonDriver(gp, "D8")
	led := gpio.NewLedDriver(gp, "D7")

	work := func() {
		button.On(gpio.ButtonPush, func(data interface{}) {
			fmt.Println("On!")
			led.On()
		})

		button.On(gpio.ButtonRelease, func(data interface{}) {
			fmt.Println("off...")
			led.Off()
		})
	}

	robot := gobot.NewRobot("sensors",
		[]gobot.Connection{board},
		[]gobot.Device{gp, button, led},
		work,
	)

	robot.Start()
}
