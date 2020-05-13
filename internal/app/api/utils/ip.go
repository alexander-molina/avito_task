package utils

import (
	"fmt"
	"log"
	"net"
)

const (
	maskBytesLen = "24"
)

// Error codes
const (
	OK                = -1
	UknownIPError int = iota
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
		case UknownIPError:
			text = "Unknown error"
		case NotIPv4:
			text = "Provided ip address is not a valid IPv4 address"
		}
	}

	return &IPError{code, text}
}

// ExtractSubnet validate IPv4 and extract subnet using mask 255.255.255.0.
func ExtractSubnet(IPAddr string) (string, error) {
	IP := net.ParseIP(IPAddr)
	IPAddr += "/" + maskBytesLen

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
