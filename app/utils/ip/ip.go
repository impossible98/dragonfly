package ip

import (
	// import built-in packages
	"net"
)

func Ip() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		// return
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip := ipnet.IP.String()
				// return
				return ip
			}
		}
	}
	return ""
}
