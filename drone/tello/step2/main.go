package main

import (
	"fmt"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
)

func main() {
	var currentFlightData *tello.FlightData

	fdTicker := time.NewTicker(time.Second * 2)
	stopChan := make(chan bool, 1)

	drone    := tello.NewDriver("8888")

	work := func() {
		fmt.Println("takeoff...")

		drone.On(tello.FlightDataEvent, func(data interface{}) {
			fd := data.(*tello.FlightData)
			currentFlightData = fd
		})

		drone.TakeOff()

		go func(){
			time.Sleep(time.Second * 10)
			stopChan <- true
		}()

		go func() {
			for {
				select {
				case <- fdTicker.C:
					printFlightData(currentFlightData)
				case <-stopChan:
					drone.Land()
				}
			}
		}()

	}

	robot := gobot.NewRobot("tello",
		[]gobot.Connection{},
		[]gobot.Device{drone},
		work,
	)

	robot.Start()
}

func printFlightData(d *tello.FlightData) {
	if d.BatteryLow {
		fmt.Printf(" -- Battery low: %d% --\n", d.BatteryPercentage)
	}

	displayData := `
Height:         %d
Ground Speed:   %d
Light Strength: %d

`
	fmt.Printf(displayData, d.Height, d.GroundSpeed, d.LightStrength)
}
