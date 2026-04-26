package interceptor

import (
	"context"
	"ginskeleton/app/global/variable"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func GrpcRequestLog() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		startTime := time.Now()

		clientIP := "unknown"
		if p, ok := peer.FromContext(ctx); ok {
			clientIP = p.Addr.String()
		}

		reqJson, _ := protojson.Marshal(req.(proto.Message))

		resp, err := handler(ctx, req)

		costTime := time.Since(startTime).Milliseconds()

		if err != nil {
			variable.ZapLog.Error("grpc - ",
				zap.String("method", info.FullMethod),
				zap.String("client_ip", clientIP),
				zap.String("request", string(reqJson)),
				zap.Int64("cost_ms", costTime),
				zap.Error(err),
			)
		} else {
			//respJson, _ := protojson.Marshal(resp.(proto.Message))
			variable.ZapLog.Info("grpc - ",
				zap.String("method", info.FullMethod),
				zap.String("client_ip", clientIP),
				zap.String("request", string(reqJson)),
				//zap.String("response", string(respJson)),
				zap.Int64("cost_ms", costTime),
			)
		}

		return resp, err
	}
}
