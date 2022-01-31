package server

import (
	"github.com/gin-gonic/gin"
	"github.com/reshimahendra/lbw-go/internal/app/account"
	"github.com/reshimahendra/lbw-go/internal/config"
	"github.com/reshimahendra/lbw-go/internal/database"
	"github.com/reshimahendra/lbw-go/internal/pkg/logger"
)

func Run() {
    // load configuration
    if err := config.Setup(); err != nil {
        logger.Errorf("fail loading configuration: %v", err)
    }

    // connect to database
    pool, _, err := database.NewDBPool(config.Get().Database)
    if err != nil {
        logger.Errorf("fail connecting to database: %v", err)
    }
    // defer pool.Close()

    // prepare server
    mode, err := config.Get().Server.GetMode()
    if err != nil {
        logger.Errorf("error loading server mode: %v", err)
    }

    // prepare gin engine
    var router *gin.Engine
    if mode == "production" {
        gin.SetMode(gin.ReleaseMode)
        router = gin.New()
	    router.SetTrustedProxies(config.Get().Server.TrustedProxies)
        router.Use(gin.Logger())
        router.Use(gin.Recovery())
    } else {
        router = gin.Default()
    }
    
    // prepare router for account app
    account.Router(pool, router)

    welcome("LotusBW", "http://127.0.0.1:8000", "-", 46) 
    router.Run(":8000")
}
