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

func turnOff() {
	blue.Off()
	green.Off()
}

func reset() {
	turnOff()
	fmt.Println("Airlock ready.")
	green.On()
}

func lock() {
	turnOff()
	fmt.Println("On!")
	blue.On()
}

func main() {
	board := firmata.NewAdaptor(os.Args[1])

	// digital devices
	button = gpio.NewButtonDriver(board, "2")
	blue = gpio.NewLedDriver(board, "3")
	green = gpio.NewLedDriver(board, "4")

	work := func() {
		reset()

		button.On(gpio.ButtonPush, func(data interface{}) {
			lock()
		})

		button.On(gpio.ButtonRelease, func(data interface{}) {
			reset()
		})
	}

	// cf. https://en.wikipedia.org/wiki/Airlock
	robot := gobot.NewRobot("airlock",
		[]gobot.Connection{board},
		[]gobot.Device{button, blue, green},
		work,
	)

	robot.Start()
}
