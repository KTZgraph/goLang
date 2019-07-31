package main

//Uruchomienie na windowsie
// go run main.go connector_couchdb.go connector_mysql.go models.go views.go
//go run main.go connector_couchdb.go connector_mysql.go models.go views.go tokens.go
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	//https://razil.cc/post/2018/10/go-cors-request-on-gin/
	return func(c *gin.Context) {
		method := c.Request.Method

		origin := c.Request.Header.Get("Origin")

		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization") // Header
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")

		}

		if method == "OPTIONS" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization") // Header
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}

func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()
	router.Use(Cors())

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	api.GET("/getUsersView", getUsersView)                    //widok pobrania wszystkich użytkowników
	api.POST("/addNewUserView", addNewUserView)               // json z danymi nowego uzytkownika
	api.POST("/loginUserView", loginUserView)                 //logowanie
	api.POST("/updateUserPointsView", updateUserPointsView)   // inkrementacja punktow
	api.POST("/addNewCardView", addNewCardView)               // [couchdb] dodawanie nowej karty
	api.POST("/getAllLoggedUsersView", getAllLoggedUsersView) // [mysql] pobieranie listy zalogowanych uzytkownikow
	api.POST("/updateLoggedUserView", updateLoggedUserView)   //aktualizacja ze uzytkownik zalogowany
	api.POST("/updateLogoutUserView", updateLogoutUserView)   //aktualizacja stanu jak uzytkownik sie wyloguje
	api.POST("/getAllCardsView", getAllCardsView)             //pobieranie wszystkich kart z couchd

	// Start and run the server
	router.Run(":3000")
}
