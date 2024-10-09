package utils

import (
	"fmt"
)

func FormatDSN(user string, password string, host string, port int, dbname string, sslmode string, timezone string) string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s TimeZone=%s", user, password, host, port, dbname, sslmode, timezone)
}
