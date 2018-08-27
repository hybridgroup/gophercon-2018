package main

import (
	"fmt"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/joystick"
)

var drone = tello.NewDriver("8888")
var	joystickAdaptor = joystick.NewAdaptor()
var stick = joystick.NewDriver(joystickAdaptor, "dualshock3")

const(
	CommonSpeed = 30
)

func main() {
	var currentFlightData *tello.FlightData

	work := func() {
		configureStickEvents()
		fmt.Println("takeoff...")

		drone.On(tello.FlightDataEvent, func(data interface{}) {
			fd := data.(*tello.FlightData)
			currentFlightData = fd
		})

		drone.On(tello.FlipEvent, func(data interface{}) {
			fmt.Println("Flip")
		})

		drone.TakeOff(())

		gobot.Every(1*time.Second, func() {
			printFlightData(currentFlightData)
		})

		gobot.After(5*time.Second, func() {
			performFlips()
		})

		gobot.After(20*time.Second, func() {
			drone.Land()
		})
	}

	robot := gobot.NewRobot("tello",
		[]gobot.Connection{},
		[]gobot.Device{drone},
		work,
	)

	robot.Start()
}

func configureStickEvents() {
	stick.On(joystick.TrianglePress, func(data interface{}) {
		drone.TakeOff()
	})

	stick.On(joystick.XPress, func(data interface{}) {
		drone.Land()
	})

	// joysticks
	stick.On(joystick.LeftX, func(data interface{}) {
		drone.Clockwise(CommonSpeed)
	})
	stick.On(joystick.LeftY, func(data interface{}) {
		drone.Up(CommonSpeed)
	})
	stick.On(joystick.RightX, func(data interface{}) {
		drone.Left(CommonSpeed)
	})
	stick.On(joystick.RightY, func(data interface{}) {
		drone.Forward(CommonSpeed)
	})
}


func printFlightData(d *tello.FlightData) {
	if d.BatteryLow {
		fmt.Printf(" -- Battery low: %d%% --\n", d.BatteryPercentage)
	}

	displayData := `
Height:         %d
Ground Speed:   %d
Light Strength: %d

`
	fmt.Printf(displayData, d.Height, d.GroundSpeed, d.LightStrength)
}

func performFlips() {
	drone.FrontFlip()
	time.Sleep(time.Second * 3)
	drone.BackFlip()
}

