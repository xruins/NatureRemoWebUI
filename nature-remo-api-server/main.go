package main

import (
	"os"

	"github.com/xruins/NatureRemoWebUI/nature-remo-api-server/cmd"
)

func main() {
	cmd.Execute()
	os.Exit(0)
}
