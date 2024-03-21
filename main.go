package main

import (
    "flag"
    "fmt"
    "log"
    "os/exec"
    "strings"
)

func main() {
  
    networkPtr := flag.String("network", "192.168.1.1/24", "Specify the network to scan")
    flag.Parse()


    devices, err := scanNetwork(*networkPtr)
    if err != nil {
        log.Fatal(err)
    }


    fmt.Println("Discovered Devices:")
    for _, device := range devices {
        fmt.Printf("IP: %s, MAC: %s\n", device.IP, device.MAC)
    }
}

type Device struct {
    IP  string
    MAC string
}

func scanNetwork(network string) ([]Device, error) {
    var devices []Device


    cmd := exec.Command("arp", "-a", network)
    output, err := cmd.Output()
    if err != nil {
        return nil, err
    }

    lines := strings.Split(string(output), "\n")
    for _, line := range lines {
        fields := strings.Fields(line)
        if len(fields) >= 2 {
            ip := fields[0]
            mac := fields[1]
            devices = append(devices, Device{IP: ip, MAC: mac})
        }
    }

    return devices, nil
}
