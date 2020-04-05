# anonymizer
Golang anonymizer for ip addresses.

You can use it to anonymize your logs.

Use it at your own risk i will not give any legal correctness.

# usage
## install
```shell
go get github.com/fsuhrau/anonymizer
```

## example
[Try it out in Go Playground](https://play.golang.org/p/EE7afMz8wvo)
```go
package main

import (
	"fmt"
	"net"
	"github.com/fsuhrau/anonymizer"
)

func main() {
	ipV4 := "192.168.0.12"
	out := anonymizer.AnonymizeIP(ipV4)
	// prints out 192.168.0.0
	fmt.Println(out)

	ipV6 := "2001:db8:85a3:8d3:1319:8a2e:370:7348"
	out = anonymizer.AnonymizeIP(ipV6)
	// prints out 2001:db8:85a3:8d3:
	fmt.Println(out)
    
	netIP := net.ParseIP("216.58.207.67")
	out = anonymizer.AnonymizeNetIP(netIP)
	// prints out 216.58.207.0
	fmt.Println(out)
}
```
