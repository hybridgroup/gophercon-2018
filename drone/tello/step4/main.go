package main

import (
	"fmt"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
	"os"
	"os/signal"
)

var sigIntChan = make(chan os.Signal, 1)
var drone = tello.NewDriver("8888")

func main() {
	signal.Notify(sigIntChan, os.Interrupt)
	var currentFlightData *tello.FlightData

	fdTicker := time.NewTicker(time.Second * 2)
	stopChan := make(chan bool, 1)


	work := func() {
		go processSigInt()
		fmt.Println("takeoff...")

		drone.On(tello.FlightDataEvent, func(data interface{}) {
			fd := data.(*tello.FlightData)
			currentFlightData = fd
		})

		drone.On(tello.FlipEvent, func(data interface{}) {
			fmt.Println("Flip")
		})


		drone.TakeOff()

		go func(){
			time.Sleep(time.Second * 10)
			stopChan <- true
		}()

		performFlips()

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

func performFlips()  {
	drone.FrontFlip()
	time.Sleep(time.Second * 3)
	drone.BackFlip()
}

func processSigInt()  {
	<- sigIntChan

	drone.Land()
	drone.Halt()
}
