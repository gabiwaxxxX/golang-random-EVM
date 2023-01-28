package wifi_managment

import (
	"fmt"
	"os/exec"
)

func GetCurrentWifiSSID1() string {

	cmd := "networksetup -getairportnetwork en0 | cut -c 24-"
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return fmt.Sprintf("Failed to execute command: %s", cmd)
	}
	// fmt.Println("out", string(out))
	return string(out)

}
