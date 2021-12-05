package main

import (
	"fmt"
	"os"
)

func main() {
  pathToConf := os.Args[1]
  fmt.Printf("%s", pathToConf)
  
}
