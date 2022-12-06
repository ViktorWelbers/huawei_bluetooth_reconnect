package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

const filename = "huawei_bluetooth_address.txt"

var headphoneBluetoothAddress string = "00:AD:D5:B4:65:11"

func filterBluetoothDevices(deviceList []string) (string, error) {
	var bluetoothAddress string
	for _, line := range deviceList {
		n := strings.Replace(line, "(", "", -1)
		n = strings.Replace(n, "	", "", 1)
		clean := strings.Split(n, ")")
		if len(clean) > 1 {
			if strings.Contains(clean[1], "HUAWEI FreeBuds Studio") {
				fmt.Println("Found device: ", clean[1])
				fmt.Println("Address: ", clean[0])
				bluetoothAddress = clean[0]
			}
		}
	}
	if bluetoothAddress == "" {
		return "", errors.New("no HUAWEI FreeBuds Studio found")
	}
	return bluetoothAddress, nil
}

func getFreebudsAddress(filename string) string {
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		cmd := exec.Command("cmd", "/c", "btdiscovery")
		cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x08000000}
		output, err := cmd.Output()
		if err != nil {
			fmt.Println("Error: ", err)
		}
		out := string(output)
		outSplit := strings.Split(out, "\n")
		address, err := filterBluetoothDevices(outSplit)
		if err != nil {
			panic(err)
		}
		f, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		_, err = f.WriteString(address)
		return address
	} else {
		fmt.Println("Found file", filename)
		address, err := os.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		return string(address)
	}
}

func reconnectBluetoothDevice(bluetoothAddress *string) {
	disconnect := fmt.Sprintf("btcom -b %s -r -s110b", *bluetoothAddress)
	reconnect := fmt.Sprintf("btcom -b %s -c -s110b", *bluetoothAddress)
	commands := []string{disconnect, reconnect}
	for _, command := range commands {
		cmd := exec.Command("cmd", "/c", command)
		cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x08000000}
		if err := cmd.Run(); err != nil {
			fmt.Println("Error: ", err)
		}
	}
}
func main() {
	if headphoneBluetoothAddress == "" {
		headphoneBluetoothAddress = getFreebudsAddress(filename)
	}
	reconnectBluetoothDevice(&headphoneBluetoothAddress)
}
