package main

import "runtime"

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	buildFilters()
	startServer()
}

func buildFilters() {
	createPasswordFilters()
}

func startServer() {
	router := setupRoutes()
	router.Run(":3000")
}
