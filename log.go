package gonaut

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func NowDateTimeString() string {
	return time.Now().Format(time.DateTime)
}

func LogFatal(s string, p ...any) {
	panic(fmt.Sprintf("[FATAL] %s: %s\n", NowDateTimeString(), fmt.Sprintf(s, p...)))
}

func LogInfo(s string, p ...any) {
	fmt.Printf("[INFO] %s: %s\n", NowDateTimeString(), fmt.Sprintf(s, p...))
}

func LogDebug(s string, p ...any) {
	if strings.ToLower(os.Getenv("DEBUG")) == "on" {
		fmt.Printf("[DEBUG] %s: %s\n", NowDateTimeString(), fmt.Sprintf(s, p...))
	}
}
