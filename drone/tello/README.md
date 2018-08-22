# DJI Tello Drone

The DJI Tello from Ryze Robotics uses a WiFi interface with a UDP-based API.

## What you need

    - DJI Tello
    - Dualshock 3 gamepad, or compatible
    - Personal computer with Go installed
    - Works on Linux (kernel v4.14+), macOS, or Windows

## Installation

```
go get -d -u gobot.io/x/gobot/...
```

## Running the code
When you run any of these examples, you will compile and execute the code on your computer. When you are running the program, you will be communicating with the Tello using the WiFi interface.

Therefore, you must connect to the Tello drone which acts as a WiFi access point before you will be able to run any of the code.


## Code

### step01/main.go

Let's start with a simple takeoff, and then land. Make sure the drone is turned on, then run the code.

### macOS

    GODEBUG=cgocheck=0 go run drone/minidrone/step1/main.go [dronename]

### macOS & Linux

    go run drone/tello/step1/main.go

### Windows

    go run drone\tello\step1\main.go


### step02/main.go

The drone can return some flight data. Run this code:

...

### step03/main.go

The drone can move forward, backward, to the right, and the left, all while maintaining a steady altitude. Run the code.

...

### step04/main.go

The drone can perform flips while flying. Run the code.

...

### step05/main.go

Now it is time for free flight, controlled by you, the human pilot. Plug in the DS3 controller to your computer. The controls are as follows:

Run the code.

...

### step06/main.go

Now that you have mastered the flight controls, let's grab the drone video feed.

Run the code.

...
