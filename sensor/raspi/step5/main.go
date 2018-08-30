package main

import (
	"fmt"
	"net/http"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/api"
	"gobot.io/x/gobot/drivers/aio"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/raspi"
)

// Message displays a message on the attached RGB LCD.
func Message(msg string) {
	fmt.Println(msg)

	lcd.Clear()
	lcd.Home()
	lcd.Write(msg)
}

var board *raspi.Adaptor

var gp *i2c.GrovePiDriver
var lcd *i2c.GroveLcdDriver

var button *gpio.ButtonDriver
var led *gpio.LedDriver

var dial *aio.GroveRotaryDriver

func main() {
	master := gobot.NewMaster()

	rotaryValue := 0

	a := api.NewAPI(master)
	a.Get("/devices/rotary", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "{\"value\": %d}", rotaryValue)
	})
	a.Start()

	board = raspi.NewAdaptor()

	gp = i2c.NewGrovePiDriver(board)
	lcd = i2c.NewGroveLcdDriver(board)

	button = gpio.NewButtonDriver(gp, "D3")
	led = gpio.NewLedDriver(gp, "D2")

	dial = aio.NewGroveRotaryDriver(gp, "A1", 500*time.Millisecond)

	work := func() {
		Message("ready")

		button.On(gpio.ButtonPush, func(data interface{}) {
			Message("On")

			led.On()
			lcd.SetRGB(0, 255, 0)
		})

		button.On(gpio.ButtonRelease, func(data interface{}) {
			Message("Off")

			led.Off()
			lcd.SetRGB(0, 0, 255)
		})

		dial.On(aio.Data, func(data interface{}) {
			msg := fmt.Sprint("Dial: ", data)
			Message(msg)

			val, ok := data.(int)
			if !ok {
				fmt.Println("bad rotary value")
				return
			}
			rotaryValue = val
		})
	}

	robot := gobot.NewRobot("sensors",
		[]gobot.Connection{board, gp},
		[]gobot.Device{gp, button, led, lcd, dial},
		work,
	)

	master.AddRobot(robot)
	master.Start()
}
