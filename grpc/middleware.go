package grpc

import (
	"context"
	"time"

	"github.com/dynastymasra/cookbook"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/**
AuthHandler Function for check authentication send by client
Ref: https://github.com/grpc-ecosystem/go-grpc-middleware
*/
func AuthHandler(expectedToken string) grpc_auth.AuthFunc {
	return func(ctx context.Context) (context.Context, error) {
		token, err := grpc_auth.AuthFromMD(ctx, cookbook.Bearer)
		if err != nil {
			return nil, err
		}

		if token != expectedToken {
			return nil, status.Errorf(codes.Unauthenticated, "invalid token")
		}

		return ctx, nil
	}
}

/**
LogrusHandler Log duration with logrus in gRPC call
Ref: https://github.com/grpc-ecosystem/go-grpc-middleware
*/
func LogrusHandler(duration time.Duration) (string, interface{}) {
	return "duration", duration.Seconds()
}
