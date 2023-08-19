package main

import (
	"fmt"

	"github.com/AndrewSav/gosamplebuild/internal/version"
)

func main() {

	fmt.Printf("gosamplebuild %s\n", version.BuildVersion())
}
