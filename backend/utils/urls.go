package utils

import (
	"fmt"
)

func GenerateLinkTreeUrl(username string, baseURL string) string {
	return fmt.Sprintf("%s/%s", baseURL, username)
}
