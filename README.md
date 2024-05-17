# Duration Parser

A small golang library to parse YouTube video duration format (eg `DD:HH:MM:SS`)

## How to get?

```
go get github.com/ppalone/duration-parser
```

## Format

YouTube video duration format is `DD:HH:MM:SS` (For example `01:10:45:01`)

```
DD: Days
HH: Hours (must be less than 23)
MM: Minutes (must be less than 59)
SS: Seconds (must be less than 59)
```

## Usage

```go
package main

import (
	"fmt"

	parser "github.com/ppalone/duration-parser"
)

func main() {
	s := "01:30"
	d, _ := parser.Parse(s)
	fmt.Println("Duration in seconds:", d.Seconds())
}
```