# Installation

`go get github.com/aubm/interval`

# Exemple usage

```golang
package main

import (
    "fmt"
    "time"

    "github.com/aubm/interval"
)

func main() {
	stop := interval.Start(func() {
		fmt.Println("Hello, World !")
	}, 5*time.Second)

    // ...

    stop() // stops the interval
}
```
