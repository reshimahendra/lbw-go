package account

import (
	"github.com/gin-gonic/gin"
	ds "github.com/reshimahendra/lbw-go/internal/app/account/datastore"
	h "github.com/reshimahendra/lbw-go/internal/app/account/handler"
	s "github.com/reshimahendra/lbw-go/internal/app/account/service"
	db "github.com/reshimahendra/lbw-go/internal/database"
	"github.com/reshimahendra/lbw-go/internal/middleware"
)


func Router(dbPool db.IDatabase, router *gin.Engine) {
    // user.status layer setup
    userStatusDatastore := ds.NewUserStatusStore(dbPool)
    userStatusService := s.NewUserStatusService(userStatusDatastore)
    userStatusHandler := h.NewUserStatusHandler(userStatusService)

    // user.role layer setup
    userRoleDatastore   := ds.NewUserRoleStore(dbPool)
    userRoleService     := s.NewUserRoleService(userRoleDatastore)
    userRoleHandler     := h.NewUserRoleHandler(userRoleService)

    // user layer setup
    userDatastore       := ds.NewUserStore(dbPool)
    userService         := s.NewUserService(userDatastore)
    userHandler         := h.NewUserHandler(userService)

    // app router group
    user := router.Group("/account")
    user.Use(middleware.CORS())
    user.Use(middleware.Security())

    user.POST("/signup", userHandler.SignupHandler)
    user.POST("/signin", userHandler.SigninHandler)

    // need authorization
    userAuth := router.Group("/account")
    userAuth.Use(middleware.CORS())
    userAuth.Use(middleware.Security())
    userAuth.Use(middleware.Authorize())

    // Router for User
    userAuth.POST("/", userHandler.UserCreateHandler)
    user.PUT("/:id", userHandler.UserUpdateHandler)
    userAuth.DELETE("/:id", userHandler.UserDeleteHandler)
    userAuth.GET("/:id", userHandler.UserGetHandler)
    userAuth.GET("/", userHandler.UserGetsHandler)
    userAuth.POST("/refresh-token", userHandler.RefreshTokenHandler)
    userAuth.POST("/check-token", userHandler.CheckTokenHandler)

    // router for user.status
    userStatus := user.Group("/status")
    userStatus.GET("/", userStatusHandler.UserStatusGetsHandler)
    userStatus.GET("/:id", userStatusHandler.UserStatusGetHandler)
    
    // router for user.role
    // userRole := user.Group("/role")
    userRoleAuth := userAuth.Group("/role")
    userRoleAuth.POST("/", userRoleHandler.UserRoleCreateHandler)
    userRoleAuth.PUT("/:id", userRoleHandler.UserRoleUpdateHandler)
    userRoleAuth.DELETE("/:id", userRoleHandler.UserRoleDeletesHandler)
    userRoleAuth.GET("/:id", userRoleHandler.UserRoleGetHandler)
    userRoleAuth.GET("/", userRoleHandler.UserRoleGetsHandler)
}
