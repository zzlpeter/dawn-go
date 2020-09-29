package api

import "github.com/gin-gonic/gin"


// @Summary Get Auth
// @Accept x-www-form-urlencoded
// @Produce json
// @Param username query string true "userName"
// @Param password query string true "password"
// @Success 200 {object} SwaggerResponse
// @Failure 500 {object} SwaggerResponse
// @Router /assassin/cas/ [get]
func GetLogIn(c *gin.Context)  {
	//ticket := c.Query("ticket")
	//callbackUrl := ""

}

func setUserCookie() {

}