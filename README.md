# tinyid
A tiny, URL-friendly, unique string ID generator for golang.

## import
```golang
import "github.com/abhisheksoni27/tinyid"
```

## example

```golang
package main

import (
	"log"

	"github.com/abhisheksoni27/tinyid"
)

func main() {
	id := tinyid.New(10)
	log.Printf("tinyid = %s", id)
}
```