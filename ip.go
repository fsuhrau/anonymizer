package anonymizer

import (
	"net"
	"strings"
)

func AnonymizeIP(originalIP string) string {
	var ip string
	index := strings.LastIndex(originalIP, ".")
	if index >= 0 {
		// v4
		ip = originalIP[:index]
		ip += ".0"
	} else {
		// v6
		index = strings.LastIndex(originalIP, ":")
		findings := 0
		for i := range originalIP {
			if originalIP[i:i+1] == ":" {
				findings++
				if index == i || findings == 4 {
					ip = originalIP[:i] + ":"
					break
				}
			}
		}
	}
	return ip
}

func AnonymizeNetIP(originalIP net.IP) string {
	return AnonymizeIP(originalIP.String())
}
