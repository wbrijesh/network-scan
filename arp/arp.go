package arp

import (
	"fmt"
	"network-scan/definitions"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

func GetDevices() ([]definitions.Device, error) {
	if runtime.GOOS == "windows" {
		return nil, fmt.Errorf("this program is not supported on Windows")
	}

	cmd := exec.Command("arp", "-a")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	return parseARPTable(string(output)), nil
}

func parseARPTable(arpTable string) []definitions.Device {
	var devices []definitions.Device

	pattern := regexp.MustCompile(`(\d+\.\d+\.\d+\.\d+).*?(\S+:\S+:\S+:\S+:\S+:\S+)`)

	lines := strings.Split(arpTable, "\n")
	for _, line := range lines {
		matches := pattern.FindStringSubmatch(line)
		if len(matches) == 3 {
			device := definitions.Device{
				IP:  matches[1],
				MAC: matches[2],
			}
			devices = append(devices, device)
		}
	}

	return devices
}
