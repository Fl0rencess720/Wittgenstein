package server

import (
	"context"

	roleV1 "github.com/Fl0rencess720/Wittgenstein/api/gateway/role/v1"
	seminarV1 "github.com/Fl0rencess720/Wittgenstein/api/gateway/seminar/v1"
	userV1 "github.com/Fl0rencess720/Wittgenstein/api/gateway/user/v1"
	"github.com/Fl0rencess720/Wittgenstein/app/gateway/interface/internal/conf"
	"github.com/Fl0rencess720/Wittgenstein/app/gateway/interface/internal/service"
	"github.com/Fl0rencess720/Wittgenstein/pkgs/jwtc"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
)

func NewWhiteListMatcher() selector.MatchFunc {

	whiteList := make(map[string]struct{})
	whiteList["/Wittgenstein.v1.User/Login"] = struct{}{}
	whiteList["/Wittgenstein.v1.User/Register"] = struct{}{}
	whiteList["/Wittgenstein.v1.User/RefreshToken"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

func NewHTTPServer(c *conf.Server, role *service.RoleService, user *service.UserService, seminar *service.SeminarService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			selector.Server(jwtc.Auth()).Match(NewWhiteListMatcher()).Build(),
		),
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		)),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	roleV1.RegisterRoleManagerHTTPServer(srv, role)
	userV1.RegisterUserHTTPServer(srv, user)
	seminarV1.RegisterSeminarHTTPServer(srv, seminar)
	route := srv.Route("/seminar")
	seminarRouter := route.Group("/topic")
	seminarRouter.POST("starting", service.StartTopic)
	seminarRouter.POST("resuming", service.ResumeTopic)
	return srv
}
