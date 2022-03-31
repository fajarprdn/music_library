package main

import (
	"github.com/gin-gonic/gin"
	"music_library/config"
	"music_library/delivery/api"
	"music_library/delivery/middleware"
	"music_library/logger"
)

type AppServer interface {
	Run()
}

type appServer struct {
	routerEngine *gin.Engine
	cfg          config.Config
}

func (p *appServer) initHandlers() {
	p.routerEngine.Use(middleware.ErrorMiddleware())
	p.v1()
}

func (p *appServer) v1() {
	trackApiGroup := p.routerEngine.Group("/track")
	api.NewTrackApi(trackApiGroup, p.cfg.UseCaseManager.TrackRegistrationUseCase())
}

func (p *appServer) Run() {
	p.initHandlers()
	logger.Log.Info().Msgf("Server run on %s", p.cfg.ApiConfig.Url)
	err := p.routerEngine.Run(p.cfg.ApiConfig.Url)
	if err != nil {
		logger.Log.Fatal().Msg("Server failed to run")
	}

}

func Server() AppServer {
	//gin.SetMode(gin.ReleaseMode)
	//r := gin.New()
	//r.Use(gin.Recovery())
	r := gin.Default()

	c := config.NewConfig(".", "config")
	return &appServer{
		routerEngine: r,
		cfg:          c,
	}
}
