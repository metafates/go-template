package main

import (
	"github.com/metafates/go-template/cmd"
	"github.com/metafates/go-template/config"
	"github.com/metafates/go-template/log"
	"github.com/samber/lo"
)

func main() {
	// prepare config and logs
	lo.Must0(config.Init())
	lo.Must0(log.Init())

	// run the app
	cmd.Execute()
}
