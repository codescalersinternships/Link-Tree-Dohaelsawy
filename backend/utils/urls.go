package utils

import (
	"fmt"
)

func GenerateLinkTreeUrl(baseURL string, linkTreePath string,username string) string {
	return fmt.Sprintf("http://%s/%s/%s", baseURL,linkTreePath, username)
}
