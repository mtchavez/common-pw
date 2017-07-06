package main

func main() {
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
