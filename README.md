# GO CLI PACKAGE
> by [Dottics (PTY) LTD](https://dottics.com)

**CLI** is a minimal Go package created as an open source project to make it
easy to create basic CLI applications.

### Getting Started
Add the package to your project:
```bash
go get github.com/dottics/cli
```

Now getting started with a blank project:
```go
package main

import (
	"flag"
	"github.com/dottics/cli"
	"log"
	"os"
)

func main() {
	// instantiate the root level command, this can be named anything
	// as only sub commands require specific names.
	root := cli.NewCommand("main", flag.ExitOnError)
	// now we pass all the args when we run the command, note that the
	// first element is os.Args is the executable command.
	// A similar design is used for sub commands see below for more detail.
	err := root.Run(os.Args[1:])
	if err != nil {
		log.Fatalln(err)
	}
}
```
That is all that is required to get started. This implementation is obviously
pretty useless since at this time it does nothing, but will not break.

*Examples*
```bash
go run main.go
# Output: 
```
and 
```bash
# here the 'get' command does not exist
go run main.go get
# Output: 
```
and
```bash
# here the 'get' command does not exist
go run main.go -help
# Output: "
# 
# ** WIP **
# 
# "
```