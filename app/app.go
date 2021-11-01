package app

import (
	"github.com/gin-gonic/gin"
	"thoughts/app/api"
	"thoughts/app/exceptions"
)

// Server defines a server object
// this Server object will have a port and app variables
// app will be of type gin.Engine
// this can also be set directly in a function to start the server
type Server struct {
	port string
	app  *gin.Engine
}

// App returns an object of Server
// this function does not need to return anything
// it can be set to start the server on its own
func App() Server {
	return Server{
		port: "80",
		app:  gin.Default(),
	}
}

// Start is used to start the server
// an object of Server is passed as an argument
// then we use its values to initialize the routes and start the server
func (s *Server) Start() {
	api.Routes(s.app)
	err := s.app.Run(":" + s.port)
	if err != nil {
		exceptions.Error("Some error occurred and was not possible to start the server.", err)
		return
	}
}
