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

    // need authorization
    userAuth := router.Group("/account")
    userAuth.Use(middleware.CORS())
    userAuth.Use(middleware.Security())
    userAuth.Use(middleware.Authorize())

    // Router for User
    userAuth.POST("/", userHandler.UserCreateHandler)
    userAuth.PUT("/:id", userHandler.UserUpdateHandler)
    userAuth.DELETE("/:id", userHandler.UserDeleteHandler)
    user.GET("/:id", userHandler.UserGetHandler)
    userAuth.GET("/", userHandler.UserGetsHandler)

    // router for user.status
    userStatus := user.Group("/status")
    userStatus.GET("/", userStatusHandler.UserStatusGetsHandler)
    userStatus.GET("/:id", userStatusHandler.UserStatusGetHandler)
    
    // router for user.role
    userRole := user.Group("/role")
    userRole.POST("/", userRoleHandler.UserRoleCreateHandler)
    userRole.PUT("/:id", userRoleHandler.UserRoleUpdateHandler)
    userRole.DELETE("/:id", userRoleHandler.UserRoleDeletesHandler)
    userRole.GET("/:id", userRoleHandler.UserRoleGetHandler)
    userRole.GET("/", userRoleHandler.UserRoleGetsHandler)
}
