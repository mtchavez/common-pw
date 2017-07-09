package main

import (
	"runtime"

	"github.com/mtchavez/common-pw/filters"
	"github.com/mtchavez/common-pw/server"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	go filters.BuildFilters()
	server.StartServer()
}
