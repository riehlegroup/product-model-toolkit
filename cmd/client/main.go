package main

import (
	"fmt"

	"github.com/osrgroup/product-model-toolkit/pkg/version"
)

func main() {
	fmt.Println("Hello, I'm the client.")
	fmt.Println("Version: " + version.Name())
}
