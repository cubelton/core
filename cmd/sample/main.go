package main

import (
	"core/pkg/cube"
	"fmt"
)

func main() {
	theCube := cube.NewCube()
	fmt.Printf("State %s|%d", theCube.StateCode(), len(theCube.StateCode()))
}
