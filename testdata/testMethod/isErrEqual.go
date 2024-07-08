package testMethod

import "strings"

func IsErrEqual(
	getError error,
	wantsError error,
) bool {
	if getError == nil && wantsError == nil {
		return true
	}
	if getError == nil || wantsError == nil {
		return false
	}
	getErrMessage := strings.TrimSpace(getError.Error())
	wantErrMessage := strings.TrimSpace(wantsError.Error())
	return IsConEqual(getErrMessage, wantErrMessage)
}
