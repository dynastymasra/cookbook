package grpc

import (
	"context"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/dynastymasra/cookbook"

	GRPCAuth "github.com/grpc-ecosystem/go-grpc-middleware/auth"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/**
AuthInterceptor Function for check authentication send in metadata
Example: Bearer token
Ref: https://github.com/grpc-ecosystem/go-grpc-middleware
*/
func AuthInterceptor(expectedToken string) GRPCAuth.AuthFunc {
	return func(ctx context.Context) (context.Context, error) {
		token, err := GRPCAuth.AuthFromMD(ctx, cookbook.Bearer)
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
//LogrusUnaryInterceptor gRPC interceptor to log unary request duration status
this function will set field request_id from metadata
*/
func LogrusUnaryInterceptor(logger *logrus.Entry, reqID string, keys ...string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		startTime := time.Now().UTC()

		for _, key := range keys {
			val := metautils.ExtractIncoming(ctx).Get(key)
			if val == "" {
				return nil, status.Errorf(codes.FailedPrecondition, "metadata key %s not found", key)
			}
		}

		requestID := metautils.ExtractIncoming(ctx).Get(reqID)
		newCtx := context.WithValue(ctx, reqID, requestID)

		log := logger.WithFields(logrus.Fields{
			"method": info.FullMethod,
			"start":  startTime.Format(time.RFC3339),
			reqID:    requestID,
		})

		resp, err := handler(newCtx, req)

		responseTime := time.Now().UTC()
		deltaTime := responseTime.Sub(startTime)

		if err != nil {
			log.WithFields(logrus.Fields{
				"finish": responseTime.Format(time.RFC3339),
				"delta":  deltaTime,
			}).WithError(err).Warnln("gRPC request")

			return resp, err
		}

		log.WithFields(logrus.Fields{
			"finish": responseTime.Format(time.RFC3339),
			"delta":  deltaTime,
		}).Infoln("gRPC request")

		return resp, err
	}
}

/**
LogrusStreamInterceptor gRPC interceptor to log stream request duration status
this function will set field request_id from metadata
*/
func LogrusStreamInterceptor(logger *logrus.Entry, reqID string, keys ...string) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		startTime := time.Now().UTC()

		for _, key := range keys {
			val := metautils.ExtractIncoming(stream.Context()).Get(key)
			if val == "" {
				return status.Errorf(codes.FailedPrecondition, "metadata key not found")
			}
		}

		log := logger.WithFields(logrus.Fields{
			"method": info.FullMethod,
			"start":  startTime.Format(time.RFC3339),
			reqID:    metautils.ExtractIncoming(stream.Context()).Get(reqID),
		})

		err := handler(srv, stream)

		responseTime := time.Now().UTC()
		deltaTime := responseTime.Sub(startTime)

		if err != nil {
			log.WithFields(logrus.Fields{
				"finish": responseTime.Format(time.RFC3339),
				"delta":  deltaTime,
			}).WithError(err).Warnln("gRPC request")

			return err
		}

		log.WithFields(logrus.Fields{
			"finish": responseTime.Format(time.RFC3339),
			"delta":  deltaTime,
		}).Infoln("gRPC request")

		return err
	}
}
