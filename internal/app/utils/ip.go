package utils

import (
	"fmt"
	"log"
	"net"
	"strings"
)

const (
	maskBytesLen = "24"
)

// Error codes
const (
	OK                = -1
	UknownIPError int = iota
	WrongMaskPrefix
	NotIPv4
)

// IPError contains information of errors when parsin ip address
type IPError struct {
	ErrorCode int
	Text      string
}

// Error return error message
func (e *IPError) Error() string {
	return fmt.Errorf("Error code: %d\n %s", e.ErrorCode, e.Text).Error()
}

// NewIPError creates a new IPError based on error code
func NewIPError(code int, err error) *IPError {
	text := ""
	if err != nil {
		text = err.Error()
	} else {
		switch code {
		case WrongMaskPrefix:
			text = "Wrong mask prefix. Must be /24"
		case NotIPv4:
			text = "Provided ip address is not a valid IPv4 address"
		}
	}

	return &IPError{code, text}
}

// ExtractSubnet validate IPv4 and extract subnet.
// Only ip addresses with default mask are permitted.
// Default mask is 255.255.255.0 .
// If ip addres does not contain mask prefix: it will be set to default mask
func ExtractSubnet(IPAddr string) (string, error) {
	// checking prefix
	i := strings.Index(IPAddr, "/")
	var addr string
	if i < 0 {
		addr = IPAddr
		IPAddr += "/" + maskBytesLen
	} else {
		if mask := IPAddr[i+1:]; mask != maskBytesLen {
			err := NewIPError(WrongMaskPrefix, nil)
			log.Println(err)
			return "", err
		}
		addr = IPAddr[:i]
	}

	IP := net.ParseIP(addr)

	// Checking if ip is IPv4. IPv6 is not permitted
	if IP == nil || IP.To4() == nil {
		err := NewIPError(NotIPv4, nil)
		log.Println(err)
		return "", err
	}

	_, subnet, err := net.ParseCIDR(IPAddr)
	if err != nil {
		log.Println(err)
		return "", NewIPError(UknownIPError, err)
	}
	return subnet.String(), nil
}
