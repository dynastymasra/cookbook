package grpc

import (
	"context"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/dynastymasra/cookbook"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/**
AuthUnaryInterceptor Function for check authentication send by client
Example: Bearer token
Ref: https://github.com/grpc-ecosystem/go-grpc-middleware
*/
func AuthUnaryInterceptor(expectedToken string) grpc_auth.AuthFunc {
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
ValidateMetadataKeyInterceptor Function used to check metadata key included in request
*/
func ValidateMetadataKeyInterceptor(keys ...string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		for _, key := range keys {
			val := metautils.ExtractIncoming(ctx).Get(key)
			if val == "" {
				return nil, status.Errorf(codes.FailedPrecondition, "metadata key not found")
			}
		}

		resp, err := handler(ctx, req)

		return resp, err
	}
}

/**
LogrusUnaryInterceptor gRPC interceptor to log unary request duration status
*/
func LogrusUnaryInterceptor(logger *logrus.Entry) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		startTime := time.Now()

		resp, err := handler(ctx, req)

		log := logger.WithFields(logrus.Fields{
			"full_method":  info.FullMethod,
			"request_time": startTime.Format(time.RFC3339),
		})

		responseTime := time.Now()
		deltaTime := responseTime.Sub(startTime)

		log.WithFields(logrus.Fields{
			"response_time": responseTime.Format(time.RFC3339),
			"delta_time":    deltaTime,
		}).Infoln("gRPC request")

		return resp, err
	}
}

/**
LogrusStreamInterceptor gRPC interceptor to log stream request duration status
*/
func LogrusStreamInterceptor(logger *logrus.Entry) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		startTime := time.Now()

		err := handler(srv, stream)

		log := logger.WithFields(logrus.Fields{
			"full_method":  info.FullMethod,
			"request_time": startTime.Format(time.RFC3339),
		})

		responseTime := time.Now()
		deltaTime := responseTime.Sub(startTime)

		log.WithFields(logrus.Fields{
			"response_time": responseTime.Format(time.RFC3339),
			"delta_time":    deltaTime,
		}).Infoln("gRPC request")

		return err
	}
}
