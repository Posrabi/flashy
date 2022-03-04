package utils

import (
	"net"

	log "github.com/sirupsen/logrus"
)

func GetNodeIPAddress() string {
	// Does not actually establish a connection because it is udp.
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}
