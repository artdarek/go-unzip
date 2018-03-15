Package go-unzip
===================

Package go-unzip provides a very simple library to extract zip archive

## Installation
```shell
go get -u github.com/artdarek/go-unzip
```

## Examples

```go
package main

import (
    "github.com/artdarek/go-unzip"
	"fmt"
)

func main() {
	uz := unzip.New("file.zip", "directory/")
	err := uz.Extract()
	if err != nil {
		fmt.Println(err)
	}
}
```

## Contributing

Pull requests, bug fixes and issue reports are welcome.

Before proposing a change, please discuss your change by raising an issue.
