package main

import (
	"fmt"
	"os"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

var button *gpio.ButtonDriver
var blue *gpio.LedDriver
var green *gpio.LedDriver

func TurnOff() {
	blue.Off()
	green.Off()
}

func Green() {
	TurnOff()

	fmt.Println("Green!")
	green.On()
}

func Blue() {
	TurnOff()

	fmt.Println("Blue!")
	blue.On()
}

var Startup = Green

func main() {
	board := firmata.NewAdaptor(os.Args[1])

	// digital devices
	button = gpio.NewButtonDriver(board, "2")
	blue = gpio.NewLedDriver(board, "3")
	green = gpio.NewLedDriver(board, "4")

	work := func() {
		Startup()

		button.On(gpio.ButtonPush, func(data interface{}) {
			Blue()
		})

		button.On(gpio.ButtonRelease, func(data interface{}) {
			Green()
		})
	}

	// Cf. https://en.wikipedia.org/wiki/Airlock
	robot := gobot.NewRobot("airlock",
		[]gobot.Connection{board},
		[]gobot.Device{button, blue, green},
		work,
	)

	robot.Start()
}
