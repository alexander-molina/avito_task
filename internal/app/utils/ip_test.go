package utils_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/alexander-molina/avito_task/internal/app/utils"
)

type data struct {
	value       string
	result      string
	errCode     int
	description string
}

func Test_ExtractSubnet(t *testing.T) {
	var testData = [...]data{
		{"127.0.0.1", "127.0.0.0/24", utils.OK, "Correct IPv4 address"},
		{"127.0.0.1/24", "127.0.0.0/24", utils.OK, "Correct IPv4 address"},
		{"127.0.0.1/20", "", utils.WrongMaskPrefix, "Correct IPv4 address"},
		{"127.0.0.1/a", "", utils.WrongMaskPrefix, "Correct IPv4 address"},
		{"wrong", "", utils.NotIPv4, "Incorrect IPv4 address"},
		{"2001:0db8:85a3:0000:0000:8a2e:0370:7334", "", utils.NotIPv4, "Incorrect IPv4 address"},
	}

	for i, d := range testData {
		r, err := utils.ExtractSubnet(d.value)

		if r != d.result {
			t.Error("Test wrong result. " + d.description)
		}
		if err != nil {
			code := extractErrCode(err.Error())
			if d.errCode != code {
				t.Errorf("Test wrong error. Case №%d\n", i)
			}
		}
	}
}

func extractErrCode(s string) int {
	l := strings.Index(s, ":")
	r := strings.Index(s, "\n")

	res, _ := strconv.Atoi(s[l+2 : r])
	return res
}
