package server

// StartServer will setup routes and start a Gin server
func StartServer() {
	router := setupRoutes()
	router.Run(":3000")
}
