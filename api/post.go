package api

import (
	"shooting-range-server/model"

	"github.com/gin-gonic/gin"
)

func PostData(c *gin.Context) {
	var newAmmo model.Ammo

	c.BindJSON(&newAmmo)

	//Add to the DB
}
