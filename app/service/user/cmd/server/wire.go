//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/Fl0rencess720/Ayana/app/service/user/internal/biz"
	"github.com/Fl0rencess720/Ayana/app/service/user/internal/conf"
	"github.com/Fl0rencess720/Ayana/app/service/user/internal/data"
	"github.com/Fl0rencess720/Ayana/app/service/user/internal/server"
	"github.com/Fl0rencess720/Ayana/app/service/user/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Registry, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
