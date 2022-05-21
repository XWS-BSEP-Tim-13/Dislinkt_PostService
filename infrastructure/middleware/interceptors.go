package middleware

import (
	"context"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/jwt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_tags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var (
	CommonInterceptors grpc.ServerOption
)

func init() {
	logger := log.WithFields(log.Fields{
		"goapi": "server",
	})

	CommonInterceptors = grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc_logrus.UnaryServerInterceptor(logger),
		AuthInterceptor(),
		grpc_tags.UnaryServerInterceptor(),
	))
}

func AuthInterceptor() grpc.UnaryServerInterceptor {

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		tokenStr, err := grpc_auth.AuthFromMD(ctx, "Bearer")

		if err != nil {
			return req, grpc.Errorf(codes.Unauthenticated, err.Error())
		}
		token, _, err := jwt.ParseJwt(tokenStr)
		if err != nil || token == nil {
			return req, grpc.Errorf(codes.Unauthenticated, err.Error())
		} else if !token.Valid {
			return req, grpc.Errorf(codes.Unauthenticated, "Invalid Token")
		}

		//role := claims.Role
		//username := claims.Username

		newCtx := context.TODO()
		return handler(newCtx, req)
	}
}
