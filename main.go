package main

import (
	"os"

	"github.com/y13i/awsprice/lib"
)

func main() {
	awsprice.GetCLIApp().Run(os.Args)
}
