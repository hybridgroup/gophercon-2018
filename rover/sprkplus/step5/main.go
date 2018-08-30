package main

import (
	"fmt"
	"os"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
	"gobot.io/x/gobot/platforms/keyboard"
	"gobot.io/x/gobot/platforms/mqtt"
	"gobot.io/x/gobot/platforms/sphero/sprkplus"
)

func main() {
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	rover := sprkplus.NewDriver(bleAdaptor)

	mqttAdaptor := mqtt.NewAdaptor(os.Args[2], "rover")
	mqttAdaptor.SetAutoReconnect(true)

	personalizedTopic := fmt.Sprintf("basestation/heartbeat/%s", bleAdaptor.Address())
	heartbeat := mqtt.NewDriver(mqttAdaptor, personalizedTopic)

	keys := keyboard.NewDriver()

	work := func() {
		rover.On("collision", func(data interface{}) {
			fmt.Printf("collision detected = %+v \n", data)
			rover.SetRGB(255, 0, 0)
		})

		heartbeat.On(mqtt.Data, func(data interface{}) {
			fmt.Println("heartbeat")
			r := uint8(gobot.Rand(255))
			g := uint8(gobot.Rand(255))
			b := uint8(gobot.Rand(255))
			rover.SetRGB(r, g, b)
		})

		keys.On(keyboard.Key, func(data interface{}) {
			key := data.(keyboard.KeyEvent)

			switch key.Key {
			case keyboard.ArrowUp:
				rover.Roll(40, 0)
			case keyboard.ArrowRight:
				rover.Roll(40, 90)
			case keyboard.ArrowDown:
				rover.Roll(40, 180)
			case keyboard.ArrowLeft:
				rover.Roll(40, 270)
			case keyboard.Spacebar:
				rover.Stop()
			}
		})
	}

	robot := gobot.NewRobot("rover",
		[]gobot.Connection{bleAdaptor, mqttAdaptor},
		[]gobot.Device{rover, keys},
		work,
	)

	// We send our own messages to the topic we're listening
	go func() {
		for {
			res := mqttAdaptor.Publish(personalizedTopic, []byte(mqtt.Data))
			if res {
				fmt.Println("published color change message...")
			}
			time.Sleep(5 * time.Second)
		}
	}()

	robot.Start()
}
