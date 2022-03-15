package main

// Subject to modification, only include function to pass on username
import{
	"github.com/gin-gonic/gin"
	"net/http"
}

func main() {

	r :=gin.Default()
	// resolve all the templates
	r.LoadHTMLGlob("template/*")
	r.GET("/login", func(c *gin.Context){
		c.HTML(http.StatusOK,"login.html",nil)
	})

	// process requests after clicking "Login"
	// Notic: "Incorrect input" function is not actually used
	r.POST("/login", func(c *gin.Context) {
		username :=c.DefaultPostForm("username","Incorrect input")
		password :=c.DefaultPostForm("password","Incorrect input")

	// get username and password from login.html for bahavior.html
	c.HTML(http.StatusOK,"behavior.html",gin.H{
		"Name": username,
		"Password": password
	})
	})

	r.Run(":9090")
}