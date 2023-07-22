package main

import (
	"fmt"

	"${ghes_host}/${ghes_org}/ghes-version-update-test-with-private-go-module-on-ghes-instance/hello"
)

func main() {
	fmt.Printf("Hello %s\n", hello.Who())
}
