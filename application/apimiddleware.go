package application

import (
	"databaselineservice/sdk/cervello"
	"databaselineservice/utils/httperror"
	"github.com/gin-gonic/gin"
)

// CORSMiddleware ....
// allow cors origin for browsers requests
// authenticate and decode incoming token

func AuthMiddleWare(authorizationMethod string) gin.HandlerFunc {
	//cervello.Login()
	return func(c *gin.Context) {
		var user *cervello.User
		var err error

		if authorizationMethod == "token" {
			//authorize the incoming token
			//log.Println(c.GetHeader("Authorization"))
			user, err = cervello.ValidateBarearToken(c.GetHeader("Authorization"))
			if err != nil {
				c.AbortWithStatusJSON(401, httperror.NewUnAuthenticatedServerError("لا يوجد صلاحية لهذا المستخدم"))
				return
			}
			c.Set("USERINFO", user)
		} else {
			headerUserName, headerPassword, HasAuth := c.Request.BasicAuth()
			if (!HasAuth) || (headerUserName != USERNAME) || (headerPassword != PASSWORD) {
				c.AbortWithStatusJSON(401, httperror.NewUnAuthenticatedServerError("لا يوجد صلاحية لهذا المستخدم"))
				return
			}
			c.Set("USERINFO", &cervello.User{
				Token: cervello.GetCervelloToken(),
				TokenClaim: cervello.TokenClaim{
					Email: envAuthUsername,
					Sub:   envAuthPassword,
				},
			})
			user.Token = cervello.GetCervelloToken()
			user.TokenClaim = cervello.TokenClaim{
				Email: "super@iotblue.net",
				Sub:   "0d48a197-cedb-4bc6-8c05-de9c91d69209",
			}
		}
		//_, err = cervello.OrganizationService.GetOrganizationById(cervello.GetOrgID(), user.Token)
		//if err != nil {
		//	logger.LogMessage("error", "Organization not found")
		//	c.AbortWithStatusJSON(http.StatusNotFound, httperror.NewNotFoundError("لم يتم إيجاد المنظومة"))
		//	return
		//}

		// pass user object to the context
		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-Correlation-Id")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,PUT,PATCH,POST,DELETE")

		// pass browser options requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

	}
}
