package http

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"
)

type HttpServer struct {
	httpServer *http.Server
	logger     *zerolog.Logger
}

func (server *HttpServer) setRouter(router *httprouter.Router) {
	server.logger.Info().Msg("http router list")
	server.logger.Info().Msg("GET /ping")
	server.logger.Info().Msg("POST /claim")
	server.logger.Info().Msg("POST /registerToken")

	router.GET("/ping", server.Ping)
	router.POST("/claim", server.GetClaimApproval)
	router.POST("/registerToken", server.GetRegisterTokenApproval)
}

func NewHttpServer(logger *zerolog.Logger) *HttpServer {
	return &HttpServer{logger: logger}
}

func (server *HttpServer) Run(addr string) error {
	router := httprouter.New()
	server.httpServer = &http.Server{
		Addr:    addr,
		Handler: router,
	}
	server.setRouter(router)
	return server.httpServer.ListenAndServe()
}

func (server *HttpServer) Shutdown() error {
	return server.httpServer.Close()
}
