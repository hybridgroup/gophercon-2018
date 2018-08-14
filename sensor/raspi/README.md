# Raspberry Pi Sensor Station

## What you need

    - Raspberry Pi running latest Raspbian OS
    - Grove IoT Starter Kit parts
    - Personal computer with Go 1.10+ installed

## Installation

```
go get -d -u gobot.io/x/gobot/...
```

## Running the code

When you run any of these examples, you will compile the code on your computer, but then execute the binary file on the Raspberry Pi itself. 

This will require moving the compiled code onto the Raspberry Pi, and then execute the code on the Raspberry Pi itself, not on your own computer.

We have included a shell script to make this process easier. You can run it like this:

```
$ ./runner.sh step1 192.168.1.42
```

Note: You'll use the IP Address you get during the setup process

The `runner.sh` script performs the following steps for you:

For example, to compile the code for step 1:

```
$ GOARCH=arm GOOS=linux go build -o step1 step1/main.go
```

Then to move the code to the Raspberry Pi, it uses the `scp` command:

```
$ scp step1 pi@10.0.0.23:/home/pi/
```

Lastly, to execute it on your Raspberry Pi, it uses the `ssh` command:

```
$ ssh -t root@<IP of your device> ./step1
```


### step1.go - LED

### step2.go - LED, Button

### step3.go - LED, Button, RGB LCD Display

### step4.go - LED, Button, RGB LCD Display, Rotary

### step5.go - LED, Button, RGB LCD Display, Rotary, Gobot API

### step6.go - LED, Button, RGB LCD Display, Rotary, Gobot API, Light Sensor
