package main

import (
	"github.com/lcaballero/g-ed/cli"
	"github.com/lcaballero/g-ed/serve"
	"os"
)

func main() {
	pr := cli.Processes{
		Serve: serve.Serve,
	}
	cli.New(pr).Run(os.Args)
}
