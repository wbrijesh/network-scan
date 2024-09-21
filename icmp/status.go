package icmp

import (
	"net"
	"os"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func CheckICMPReachability(ipAddress string) (bool, error) {
	icmpConn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		return false, err
	}
	defer icmpConn.Close()

	icmpMessage := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.Echo{
			ID: os.Getpid() & 0xffff, Seq: 1,
			Data: []byte("ping"),
		},
	}

	messageBytes, err := icmpMessage.Marshal(nil)
	if err != nil {
		return false, err
	}

	destinationIP := &net.IPAddr{IP: net.ParseIP(ipAddress)}
	if destinationIP.IP == nil {
		return false, net.InvalidAddrError("invalid IP address")
	}

	_, err = icmpConn.WriteTo(messageBytes, destinationIP)
	if err != nil {
		return false, err
	}

	responseBuffer := make([]byte, 1500)
	err = icmpConn.SetReadDeadline(time.Now().Add(2 * time.Second))
	if err != nil {
		return false, err
	}

	bytesRead, _, err := icmpConn.ReadFrom(responseBuffer)
	if err != nil {
		return false, nil // this means device is offline this isn't an error condition
	}

	responseMessage, err := icmp.ParseMessage(1, responseBuffer[:bytesRead])
	if err != nil {
		return false, err
	}

	return responseMessage.Type == ipv4.ICMPTypeEchoReply, nil
}
