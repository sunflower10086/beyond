package cache

import "fmt"

func GetSMSCodeKey(phone string) string {
	return fmt.Sprintf("sms_%s", phone)
}
