# stun: Go Email Automation Library

[![Go Report Card](https://goreportcard.com/badge/github.com/C-o-m-o-n/stun)](https://goreportcard.com/report/github.com/C-o-m-o-n/stun)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

`stun` is a simple and flexible Go library for email automation. It provides a convenient way to send emails using the Go programming language.

## Features

- **Basic Email Sending:** Send plain text emails easily.
- **Environment Variable Support:** Securely read email credentials from environment variables.
- **Configurable SMTP Server:** Easily configure the SMTP server details.
- _[You can add more features here as you enhance the library]_

## Installation

```bash
go get github.com/yourusername/stun

```

## Usage

```go
package main

import (
	"github.com/yourusername/stun"
)

func main() {
	// Create a new Email instance
	email := stun.NewEmail("recipient@example.com", "Test Subject", "This is the email body.")

	// Send the email
	err := email.Send()
	if err != nil {
		panic(err)
	}
}
```

## Contributing

Contributions are welcome! Please feel free to open issues or pull requests.
