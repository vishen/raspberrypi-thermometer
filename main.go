package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var (
	devicesDirFlag = flag.String("devices-dir", "/sys/bus/w1/devices", "devices directory")
)

func main() {
	flag.Parse()

	deviceID, err := findDeviceID(*devicesDirFlag)
	if err != nil {
		log.Fatalf("unable to find device id: %v", err)
	}

	for {
		temperature, err := deviceTemperature(*devicesDirFlag, deviceID)
		if err != nil {
			log.Fatalf("unable to get device temperature: %v", err)
		}
		fmt.Println("Temperature:", temperature)
		time.Sleep(time.Second * 5)
	}
}

func findDeviceID(devicesDir string) (string, error) {
	fis, err := os.ReadDir(devicesDir)
	if err != nil {
		return "", err
	}
	for _, fi := range fis {
		if fi.IsDir() && strings.Contains(fi.Name(), "-") {
			return fi.Name(), nil
		}
	}
	return "", errors.New("unable to find device id")
}

func deviceTemperature(devicesDir, deviceID string) (float32, error) {
	temp, err := os.ReadFile(filepath.Join(devicesDir, deviceID, "temperature"))
	if err != nil {
		return 0, err
	}
	t, err := strconv.Atoi(string(temp))
	if err != nil {
		return 0, err
	}
	return float32(t) / 1000, nil
}
