package main

import (
	"context"
	_ "github.com/lib/pq"
	ChainHandler "github.com/nguyenkhoa0721/go-project-layout/internal/chain/handler"
	"github.com/nguyenkhoa0721/go-project-layout/pkg/common"
	"github.com/nguyenkhoa0721/go-project-layout/pkg/exception"
	"github.com/nguyenkhoa0721/go-project-layout/pkg/kafka"
	pb "github.com/nguyenkhoa0721/grpc/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	grpcServer *grpc.Server
}

func NewApp() *App {
	common.NewCommon()

	grpcLogger := grpc.UnaryInterceptor(exception.GrpcException)
	grpcServer := grpc.NewServer(grpcLogger)

	chainPublicHandler := ChainHandler.NewChainPublicHandler(common.GetCommon())
	pb.RegisterChainServer(grpcServer, chainPublicHandler)

	reflection.Register(grpcServer)

	return &App{
		grpcServer: grpcServer,
	}
}

func (app *App) runGrpcServer() {
	listener, err := net.Listen("tcp", common.GetCommon().Config.App.GrpcServerAddress)
	if err != nil {
		logrus.Error(err)
	}
	logrus.Infof("Start gRPC server at %s", listener.Addr().String())

	go func() {
		err = app.grpcServer.Serve(listener)
		if err != nil {
			logrus.Error("Cannot start gRPC server")
		}
	}()
}

func (app *App) CleanUp() {
	app.grpcServer.GracefulStop()
	common.GetCommon().CleanUp()
	logrus.Info("App: Clean up")
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())

	app := NewApp()
	app.runGrpcServer()

	go kafka.NewConsumer(ctx, common.GetCommon().Config.Kafka.Brokers, "ping", func(bytes []byte) error {
		logrus.Info(string(bytes))

		return nil
	})

	logrus.Info("App is starting. ctrl + C to stop")
	<-c
	cancel()

	logrus.Info("Stopping...")
	time.Sleep(3 * time.Second)

	defer app.CleanUp()
}
