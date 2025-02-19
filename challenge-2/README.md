# Eldoria Scrolls Decoder

## Overview

This project is a console application that retrieves and deciphers ancient scrolls from the enchanted land of Eldoria. The scrolls contain secrets hidden by the Elders using powerful spells. The application fetches the scroll content from a specified URL, extracts the secrets using a specific pattern, and displays them.

## Features

- Fetches scroll content from a given URL.
- Extracts secrets surrounded by `{*` and `*}` from the scroll content.
- Displays the extracted secrets in a structured manner.

## Requirements

- Go 1.16 or later
- Internet connection to fetch the scroll content

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/SwatiShin/copilot-hackathon-challenges.git
    git switch vortex
    cd challenge-2
    ```

2. Build the application:
    ```sh
    go build -o eldoria-scrolls-decoder
    ```

## Usage

1. Run the application:
    ```sh
    ./eldoria-scrolls-decoder
    ```

2. The application will fetch the scroll content from the specified URL and display the extracted secrets.

## Code Structure

- `main.go`: Contains the main logic for fetching and extracting secrets from the scroll content.
- `main_test.go`: Contains test cases for the functions in `main.go`.

## Functions

### `fetchScrollContent(url string) (string, error)`

Retrieves the content of the scroll from the given URL. Makes an HTTP GET request to the URL and returns the response body as a string. Returns an error if there is an issue during the HTTP request or reading the response body.

**Complexity:** O(n), where n is the size of the response body.

### `extractSecrets(content string) []string`

Extracts the secrets from the given content string. Looks for secrets surrounded by the symbols `{*` and `*}` and returns them as a slice of strings.

**Complexity:** O(m), where m is the length of the content string.

## Testing

Run the tests using the following command:
```sh
go test