package utils

import (
	"errors"
)

func getJudgeStatus(c uint8) (string, error) {
	statusString := []string{"AC", "WA", "RTE", "IR", "OLE", "MLE", "TLE", "IE"}
	if c > uint8(len(statusString)) {
		return "", errors.New("strange judge status code")
	} else {
		return statusString[c], nil
	}
}
