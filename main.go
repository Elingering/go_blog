package main

import (
	_ "bolg/app/Providers"
	"bolg/routes"
)

func main() {
	r := routes.ApiRoutes()
	// Listen and Server in 0.0.0.0:8080
	r.Run("go_blog.com:8080")
}
