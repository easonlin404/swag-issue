package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/easonlin404/swag-issue/docs"
)

// @title API
// @version 1.0
// @description fdsfds
// @termsOfService asfsadf

// @contact.name API Support
// @contact.url dsgdfsg
// @contact.email fdssgfdsfd

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost
// @BasePath /v1
func main() {
	r := gin.New()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}

// @Summary fdsgfdsgfd
// @Description dfsgdfsgfd
// @ID get-all-tags
// @Accept  json
// @Produce  json
// @Param   some_id     path    int     true        "Some ID"
// @Success 200 {array} models.Tag
// @Router /v1/tags [get]
func GetTags(c *gin.Context) {}


