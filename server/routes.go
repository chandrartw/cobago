package server

import (
//	"os"

//	"github.com/go-redis/redis"
	//"gitlab.myih.telkom.co.id/bpd/nprm/nprm-backend/-/tree/development/auth"
	//auth "github.com/putriapriandi/cobago/auth"
	//"gitlab.myih.telkom.co.id/bpd/nprm/nprm-backend/-/tree/development/middleware"
	//middleware "github.com/putriapriandi/cobago/middleware"

	"github.com/gin-gonic/gin"
	//"gitlab.myih.telkom.co.id/bpd/nprm/nprm-backend/-/tree/development/controller"
	controller "github.com/putriapriandi/cobago/controller"

	// swagger embed files
	// gin-swagger middleware
)

// func NewRedisDB(host string, port string) *redis.Client {
// 	redisClient := redis.NewClient(&redis.Options{
// 		Addr: host + ":" + port,
// 		DB:   0,
// 	})
// 	return redisClient
// }

func Routes(router *gin.Engine, inDB *controller.InDB) {
	// //redis details
	// redis_host := os.Getenv("KVS_HOST")
	// redis_port := os.Getenv("REDIS_PORT")

	// redisClient := NewRedisDB(redis_host, redis_port)

	// var rd = auth.NewAuth(redisClient)
	// var tk = auth.NewToken()
//	var service = middleware.NewProfile(rd, tk)

	// router.POST("/login", service.Login)
	// router.POST("/logout", service.Logout)
	// router.GET("/captcha", middleware.CaptchaHandler)
	// router.POST("/captcha", middleware.CaptchaSolver)
	// router.POST("/forget/request", inDB.RequestForgetPassword)
	// router.GET("/verify/activation/:code", inDB.VerifyUser)
	// router.GET("/verify/forget/:code", inDB.VerifyForgetPassword)

	//routeU := router.Group("/users")
	//routeUser := routeU.Group("/", service.AuthMiddleware)
		router.POST("/user", inDB.CreateUser)// buat akun
		router.GET("/user",inDB.GetAllUser)// get all 
		router.GET("/user/:id", inDB.GetUser)//bikin baruu
		router.PUT("/user/:id", inDB.UpdateUser)
		router.DELETE("/user/:id",inDB.DeleteUser)
		// routeUser.GET("/get/:id", inDB.GetUser)
		// routeUser.GET("/get-all", inDB.GetAllUser)
		// routeUser.PUT("/update", inDB.UpdateUser)
		// routeUser.DELETE("/delete/:id", inDB.DeleteUser)
		// routeUser.POST("/question/create", inDB.CreateQuestion)
		// routeUser.POST("/question/update", inDB.CreateQuestion)
	
		
	//router.POST("/users/create", inDB.CreateUser) telegram

	// router.POST("/users/create", inDB.CreateUser)

}
