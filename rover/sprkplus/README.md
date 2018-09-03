# Sphero SPRK+

The Sphero SPRK+ and Sphero Ollie, and Sphero BB-8 all use the same API. However,
they have separate Gobot drivers to accommodate their other differences.

## What you need

    - Sphero Ollie, SPRK+, or BB-8
    - Personal computer with Go installed, and a Bluetooth 4.0 radio.
    - Linux or OS X

## Installation

```
go get -d -u gobot.io/x/gobot/...
```

## Running the code
When you run any of these examples, you will compile and execute the code on your computer. When you are running the program, you will be communicating with the robot using the Bluetooth Low Energy (LE) interface.

To compile/run the code, substitute the name of your SPRK+, Ollie or BB-8 as needed.

### OS X

To run any of the Gobot BLE code on OS X, you must use the `GODEBUG=cgocheck=0` flag.

For example:

```
$ GODEBUG=cgocheck=0 go run rover/sprkplus/step1/main.go BB-128E
```

### Linux

On Linux the BLE code will need to run as a root user account. The easiest way to accomplish this is probably to use `go build` to build your program, and then to run the requesting executable using `sudo`.

For example:

```
$ go build -o step1 rover/ollie/step1/main.go
$ sudo ./step1 2B-123E
```

## Code

### step1

This tests that the Sphero SPRK+ or Ollie is connected correctly to your computer, by blinking the built-in LED.

#### OS X

```
$ GODEBUG=cgocheck=0 go run rover/sprkplus/step1/main.go SK-1234
```

OR

```
$ GODEBUG=cgocheck=0 go run rover/ollie/step1/main.go 2B-1234
```

#### Linux

```
$ go build -o step1 rover/sprkplus/step1/main.go
$ sudo ./step1 SK-1234
```

OR

```
$ go build -o step1 rover/ollie/step1/main.go
$ sudo ./step1 2B-1234
```

### step2

Rolls around at random.


#### OS X

```
$ GODEBUG=cgocheck=0 go run rover/sprkplus/step2/main.go SK-1234
```

OR

```
$ GODEBUG=cgocheck=0 go run rover/ollie/step2/main.go 2B-1234
```

#### Linux

```
$ go build -o step2 rover/sprkplus/step2/main.go
$ sudo ./step2 SK-1234
```

OR

```
$ go build -o step2 rover/ollie/step2/main.go
$ sudo ./step2 2B-1234
```

### step3

Gets collision notifications from robot.

#### OS X

```
$ GODEBUG=cgocheck=0 go run rover/sprkplus/step3/main.go SK-1234
```

OR

```
$ GODEBUG=cgocheck=0 go run rover/ollie/step3/main.go 2B-1234
```

#### Linux

```
$ go build -o step3 rover/sprkplus/step3/main.go
$ sudo ./step3 SK-1234
```

OR

```
$ go build -o step3 rover/ollie/step3/main.go
$ sudo ./step3 2B-1234
```

### step4/main.go

This step has us receiving a heartbeat signal from the "base station" using the MQTT machine to machine messaging protocol. No additional hardware needs to be connected. 

We will connect to a MQTT machine to machine messaging server that is maintained by the Eclipse Foundation for public testing.


When the heartbeat data is received from the base station, the built-in LED will change color.

#### OS X

```
$ GODEBUG=cgocheck=0 go run rover/sprkplus/step4/main.go SK-1234 ssl://iot.eclipse.org:8883
```

OR

```
$ GODEBUG=cgocheck=0 go run rover/ollie/step4/main.go 2B-1234 ssl://iot.eclipse.org:8883
```

#### Linux

```
$ go build -o step4 rover/sprkplus/step4/main.go
$ sudo ./step4 SK-1234 ssl://iot.eclipse.org:8883
```

OR

```
$ go build -o step4 rover/ollie/step4/main.go
$ sudo ./step4 2B-1234 ssl://iot.eclipse.org:8883
```

### step5/main.go

Control robot using keyboard arrow keys.

#### OS X

```
$ GODEBUG=cgocheck=0 go run rover/sprkplus/step5/main.go SK-1234 ssl://iot.eclipse.org:8883
```

OR

```
$ GODEBUG=cgocheck=0 go run rover/ollie/step5/main.go 2B-1234 ssl://iot.eclipse.org:8883
```

#### Linux

```
$ go build -o step5 rover/sprkplus/step5/main.go
$ sudo ./step5 SK-1234 ssl://iot.eclipse.org:8883
```

OR

```
$ go build -o step5 rover/ollie/step5/main.go
$ sudo ./step5 2B-1234 ssl://iot.eclipse.org:8883
```

### step6/main.go

Control robot using keyboard to collect data and send to base station.

#### OS X

```
$ GODEBUG=cgocheck=0 go run rover/sprkplus/step6/main.go SK-1234 ssl://iot.eclipse.org:8883
```

OR

```
$ GODEBUG=cgocheck=0 go run rover/ollie/step6/main.go 2B-1234 ssl://iot.eclipse.org:8883
```

#### Linux

```
$ go build -o step6 rover/sprkplus/step6/main.go
$ sudo ./step6 SK-1234 ssl://iot.eclipse.org:8883
```

OR

```
$ go build -o step6 rover/ollie/step6/main.go
$ sudo ./step6 2B-1234 ssl://iot.eclipse.org:8883
```

## License

Copyright (c) 2015-2017 The Hybrid Group. Licensed under the MIT license.
