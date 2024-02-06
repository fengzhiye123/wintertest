package router

import (
	"github.com/gin-gonic/gin"
	"winter_test/app/api/internal/service"
)

func InitRouter(port string) error {
	r := gin.Default()
	r.POST("/register", service.Register)

	err := r.Run(":" + port)
	if err != nil {
		return err
	}
	return err
}
