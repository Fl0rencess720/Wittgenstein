// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/Fl0rencess720/Wittgenstein/app/gateway/user/internal/biz"
	"github.com/Fl0rencess720/Wittgenstein/app/gateway/user/internal/conf"
	"github.com/Fl0rencess720/Wittgenstein/app/gateway/user/internal/data"
	"github.com/Fl0rencess720/Wittgenstein/app/gateway/user/internal/server"
	"github.com/Fl0rencess720/Wittgenstein/app/gateway/user/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confService *conf.Service, confData *conf.Data, registry *conf.Registry, logger log.Logger) (*kratos.App, func(), error) {
	client := data.NewRedis(confData)
	discovery := server.NewDiscovery(registry)
	userClient := data.NewUserServiceClient(confService, discovery)
	dataData, cleanup, err := data.NewData(confData, logger, client, userClient)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	userUsecase := biz.NewUserUsecase(userRepo, logger, userClient)
	userService := service.NewUserService(userUsecase)
	httpServer := server.NewHTTPServer(confServer, userService, logger)
	registrar := server.NewRegistrar(registry)
	app := newApp(logger, httpServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
