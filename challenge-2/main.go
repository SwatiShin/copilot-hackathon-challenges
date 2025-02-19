package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// fetchScrollContent retrieves the content of the scroll from the given URL.
// It makes an HTTP GET request to the URL and returns the response body as a string.
// If there is an error during the HTTP request or reading the response body, it returns an error.
func fetchScrollContent(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// extractSecrets extracts the secrets from the given content string.
// It looks for secrets surrounded by the symbols `{*` and `*}` and returns them as a slice of strings.
func extractSecrets(content string) []string {
	var secrets []string
	start := 0
	for {
		start = strings.Index(content[start:], "{*")
		if start == -1 {
			break
		}
		start += 2
		end := strings.Index(content[start:], "*}")
		if end == -1 {
			break
		}
		secrets = append(secrets, content[start:start+end])
		start += end + 2
	}
	return secrets
}

func main() {
	url := "https://raw.githubusercontent.com/sombaner/copilot-hackathon-challenges/main/Data/Scrolls.txt"
	content, err := fetchScrollContent(url)
	if err != nil {
		fmt.Printf("Error fetching scroll content: %v\n", err)
		return
	}

	secrets := extractSecrets(content)
	fmt.Println("Extracted Secrets:")
	for _, secret := range secrets {
		fmt.Println(secret)
	}
}
