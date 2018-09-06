package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
	"gobot.io/x/gobot/platforms/joystick"
	"gobot.io/x/gobot/platforms/mqtt"
	"gobot.io/x/gobot/platforms/sphero/sprkplus"
)

var robot *gobot.Robot
var mqttAdaptor *mqtt.Adaptor

//ReportCollision - Some info
func ReportCollision(data interface{}) {
	buf := new(bytes.Buffer)
	msg, _ := json.Marshal(data)
	binary.Write(buf, binary.LittleEndian, msg)
	mqttAdaptor.Publish("rovers/"+robot.Name+"/collision", buf.Bytes())
}

func main() {
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	rover := sprkplus.NewDriver(bleAdaptor)

	adaptor := joystick.NewAdaptor()
	adaptor.Connect()

	stick := joystick.NewDriver(adaptor, "dualshock3")

	mqttAdaptor = mqtt.NewAdaptor(os.Args[2], "rover")
	mqttAdaptor.SetAutoReconnect(true)

	personalizedTopic := fmt.Sprintf("basestation/heartbeat/%s", bleAdaptor.Address())
	heartbeat := mqtt.NewDriver(mqttAdaptor, personalizedTopic)

	work := func() {
		rover.On("collision", func(data interface{}) {
			fmt.Printf("collision detected = %+v \n", data)
			rover.SetRGB(255, 0, 0)
			ReportCollision(data)
		})

		heartbeat.On(mqtt.Data, func(data interface{}) {
			r := uint8(gobot.Rand(255))
			g := uint8(gobot.Rand(255))
			b := uint8(gobot.Rand(255))
			rover.SetRGB(r, g, b)
		})

		stick.On(joystick.UpPress, func(data interface{}) {
			rover.Roll(40, 90)
		})
		stick.On(joystick.DownPress, func(data interface{}) {
			rover.Roll(40, 270)
		})
		stick.On(joystick.LeftPress, func(data interface{}) {
			rover.Roll(40, 0)
		})
		stick.On(joystick.RightPress, func(data interface{}) {
			rover.Roll(40, 180)
		})

		stick.On(joystick.LeftX, func(data interface{}) {
			speed := uint8(data.(int16))
			if speed > 0 {
				rover.Roll(speed, 0)
			} else {
				rover.Roll(speed, 180)
			}
		})
		stick.On(joystick.LeftY, func(data interface{}) {
			speed := uint8(data.(int16))
			if speed == 0 {
				return
			}
			if speed > 0 {
				rover.Roll(speed, 90)
			} else {
				rover.Roll(speed, 270)
			}
		})

		stick.On(joystick.XPress, func(data interface{}) {
			rover.Stop()
		})

	}

	robot = gobot.NewRobot([]gobot.Connection{bleAdaptor, mqttAdaptor},
		[]gobot.Device{rover, stick},
		work,
	)

	go func() {
		for {
			mqttAdaptor.Publish(personalizedTopic, []byte(mqtt.Data))
			time.Sleep(5 * time.Second)
		}
	}()

	robot.Start()
}
