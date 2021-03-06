package server

import (
	"fmt"
	"html"

	"github.com/gin-gonic/gin"
	"github.com/mtchavez/common-pw/filters"
)

// ValidateForm to deserialize POST /validate form
type ValidateForm struct {
	Password string `form:"password" json:"password" binding:"required"`
}

func setupRoutes() *gin.Engine {
	router := gin.Default()
	router.POST("/validate", validatePOST)
	return router
}

func validatePOST(c *gin.Context) {
	var json ValidateForm
	c.BindJSON(&json)
	validPass := len(json.Password) > 0 && len(json.Password) < 200
	if !validPass {
		c.JSON(400, gin.H{
			"status": "failed",
			"error":  "a password must be provided",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  "OK",
		"top196":  html.EscapeString(fmt.Sprintf("%v", filters.PWFilters.Top196.Lookup([]byte(json.Password)))),
		"top3575": html.EscapeString(fmt.Sprintf("%v", filters.PWFilters.Top3575.Lookup([]byte(json.Password)))),
		"top95k":  html.EscapeString(fmt.Sprintf("%v", filters.PWFilters.Top95k.Lookup([]byte(json.Password)))),
		"top32m":  html.EscapeString(fmt.Sprintf("%v", filters.PWFilters.Top32m.Lookup([]byte(json.Password)))),
	})
}
