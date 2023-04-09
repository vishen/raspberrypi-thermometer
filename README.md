# Raspberry PI Thermometer

A thermometer using the Raspberry PI and [Wire Digital Temperature Sensor](https://raspberry.piaustralia.com.au/products/1-wire-digital-temperature-sensor-for-raspberry-pi-unassembled-1m).

## Prerequisites

- [Enable 1-wire on Raspberry PI](https://www.raspberrypi-spy.co.uk/2018/02/enable-1-wire-interface-raspberry-pi/)

## Building for Raspberry PI

	$ GOOS=linux GOARCH=arm GOARM=7 go build -o ./build/armv7-linux/thermometer

