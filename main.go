package main

import (
	"github.com/gin-gonic/gin"

	"shooting-range-server/api"
	"shooting-range-server/middleware"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	authorized := r.Group("/api")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.GET("/getammo", api.GetAmmo)
		authorized.GET("/getclients", api.GetClients)
		authorized.GET("/getemployees", api.GetEmployees)
		authorized.GET("/getguns", api.GetGuns)
		authorized.GET("/getlines", api.GetLines)
		authorized.GET("/getmanufacturers", api.GetManufacturers)
		authorized.GET("/getreservations", api.GetReservations)
		
		authorized.POST("/getammo", api.PostAmmo)
		authorized.POST("/getclients", api.PostClients)
		authorized.POST("/getemployees", api.PostEmployees)
		authorized.POST("/getguns", api.PostGuns)
		authorized.POST("/getlines", api.PostLines)
		authorized.POST("/getmanufacturers", api.PostManufacturers)
		authorized.POST("/getreservations", api.PostReservations)
		
		authorized.POST("/getammo", api.PostAmmo)
		authorized.POST("/getclients", api.PostClients)
		authorized.POST("/getemployees", api.PostEmployees)
		authorized.POST("/getguns", api.PostGuns)
		authorized.POST("/getlines", api.PostLines)
		authorized.POST("/getmanufacturers", api.PostManufacturers)
		authorized.POST("/getreservations", api.PostReservations)

	}

	r.Run(":8080")
}
