package ip

import (
	"errors"
	"net"
)

func InterfaceIP(name string) (string, error) {
	ifi, err := net.InterfaceByName(name)
	if err != nil {
		return "", err
	}
	addrs, err := ifi.Addrs()
	if err != nil {
		return "", err
	}
	for _, addr := range addrs {
		return addr.(*net.IPNet).IP.String(), nil
	}
	return "", errors.New("interface had no addresses")
}
