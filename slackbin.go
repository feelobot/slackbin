package main

import "github.com/gin-gonic/gin"
import "net/http"

func main() {
    r := gin.Default()
    //token=ToMNl0zEGBeW4E8aJwiGWFOO
    //team_id=T0001
    //team_domain=example
    //channel_id=C2147483705
    //channel_name=test
    //user_id=U2147483697
    //user_name=Steve
    // gin.H is a shortcut for map[string]interface{}
    r.POST("/", func(c *gin.Context) {
      action := c.Param("action")
      c.String(http.StatusOK, "Hello %s", action)
    })
    // Listen and server on 0.0.0.0:8080
    r.Run(":8080")
}


func ls() {

}
