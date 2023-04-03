package cmd

import (
	"context"
	"fmt"
	"log"
	"mage/cmd/api/leaderboard"
	"mage/cmd/api/user"
	"mage/pkg/configs"
	"mage/pkg/db"
	"mage/pkg/utils"
	"net"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Server struct {
	userService  *user.UserService
	boardService *leaderboard.LeaderBoardService
}

func Run() {
	configs.LoadConfig()
	db.ConnectDatabases()

	zapLogger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer zapLogger.Sync()
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(authInterceptor))

	serv := Server{
		userService:  user.NewUserService(),
		boardService: leaderboard.NewLeaderBoardService(),
	}
	//UserServiceServer
	user.RegisterUserServiceServer(grpcServer, serv.userService)
	//LeaderBoardServiceServer
	leaderboard.RegisterLeaderBoardServiceServer(grpcServer, serv.boardService)

	lis, err := net.Listen("tcp", ":"+configs.Cfg.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()

	fmt.Printf("Starting server on port %v\n", configs.Cfg.Port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func authInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	defer log.Printf("Method: %s\tDuration: %v", info.FullMethod, time.Since(start))

	if info.FullMethod == "/LeaderBoardService/Endgame" || info.FullMethod == "/LeaderBoardService/LeaderBoard" {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "Authorization token not found")
		}
		authHeader := md.Get("token")
		if len(authHeader) == 0 {
			return nil, status.Errorf(codes.Unauthenticated, "Authorization token not found")
		}

		claims := utils.Claims{}
		token, err := jwt.ParseWithClaims(authHeader[0], &claims, func(token *jwt.Token) (interface{}, error) {
			if jwt.SigningMethodHS256 != token.Method {
				return nil, status.Errorf(codes.InvalidArgument, "Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(configs.Cfg.SecretKey), nil
		})
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "Could not authenticate: %v", err)
		}
		if !token.Valid {
			return nil, status.Errorf(codes.Unauthenticated, "Could not authenticate: %v", err)
		}
		ctx = context.WithValue(ctx, "userId", claims.UserId)
	}

	return handler(ctx, req)
}
