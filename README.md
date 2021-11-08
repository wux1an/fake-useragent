# fake-useragent

Provide random user agent 

# Install

```console
go get -u github.com/wux1an/fake-useragent
```

# Usage

```golang
package main

import (
	"fmt"
	"github.com/wux1an/fake-useragent"
)

func main() {
	fmt.Println(ua.Random())
	fmt.Println(ua.RandomType(ua.Desktop))
	fmt.Println(ua.RandomType(ua.Mobile))
	fmt.Println(ua.RandomType(ua.Chrome))
}

// Output:
// Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.0.1 Safari/605.1.15
// Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36
// Mozilla/5.0 (Linux; Android 8.1.0; Moto G (5S)) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.80 Mobile Safari/537.36
// Mozilla/5.0 (Linux; Android 7.1; vivo 1716 Build/N2G47H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.98 Mobile Safari/537.36
```
