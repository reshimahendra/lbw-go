package account

import (
	"github.com/gin-gonic/gin"
	ds "github.com/reshimahendra/lbw-go/internal/app/account/datastore"
	h "github.com/reshimahendra/lbw-go/internal/app/account/handler"
	s "github.com/reshimahendra/lbw-go/internal/app/account/service"
	db "github.com/reshimahendra/lbw-go/internal/database"
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

    // Router for User
    user.POST("/", userHandler.UserCreateHandler)
    user.PUT("/:id", userHandler.UserUpdateHandler)
    user.DELETE("/:id", userHandler.UserDeleteHandler)
    user.GET("/:id", userHandler.UserGetHandler)
    user.GET("/", userHandler.UserGetsHandler)

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
