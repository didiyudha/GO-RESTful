package main

import (
	"github.com/julienschmidt/httprouter"
)

// NewRouter return all router
func NewRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/users", FindAllUsers)
	router.POST("/users", SaveUsers)
	router.GET("/users/:id", FindUserByID)
	router.PUT("/users/:id", UpdateUser)
	router.DELETE("/users/:id", DeleteUser)
	return router
}
