package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jessevdk/go-flags"
	"github.com/rs/cors"
	"github.com/smartatransit/fivepoints/api/v1/schedule"
	"github.com/smartatransit/fivepoints/cmd/scheduleapi/handler"
	"github.com/smartatransit/fivepoints/pkg/authorize"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type options struct {
	JWTSigningSecret string `long:"jwt-signing-secret" env:"JWT_SIGNING_SECRET" description:"the jwt signing secret" required:"true"`
	DynamoTableName  string `long:"dynamo-table-name" env:"DYNAMO_TABLE_NAME" description:"dynamo table name"`
	GRPCPort         int    `long:"grpc-port" env:"GRPC_PORT" description:"port that the grpc server will be started" required:"true"`
	RESTPort         int    `long:"rest-port" env:"PORT" description:"port that the rest server will be started" required:"true"`
}

func main() {
	fmt.Println("Starting fivepoints")
	var opts options
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	logger, _ := zap.NewProduction()
	defer func() {
		_ = logger.Sync() // flushes buffer, if any
	}()

	awsSession := session.Must(session.NewSession())
	svc := dynamodb.New(awsSession)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: false, EmitDefaults: true}))
	s := grpc.NewServer()
	h := &http.Server{
		Addr:         fmt.Sprintf(":%d", opts.RESTPort),
		Handler:      cors.AllowAll().Handler(mux),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	authorizer := authorize.NewClient(opts.JWTSigningSecret)
	grpcAddr := fmt.Sprintf("localhost:%d", opts.GRPCPort)
	scheduleAPI := handler.NewWithDefaultEndpoints(opts.DynamoTableName, svc, authorizer)
	dopts := []grpc.DialOption{grpc.WithInsecure()}
	schedule.RegisterScheduleServiceServer(s, scheduleAPI)
	err = schedule.RegisterScheduleServiceHandlerFromEndpoint(ctx, mux, grpcAddr, dopts)
	if err != nil {
		panic(err)
	}

	errC := make(chan error, 1)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		logger.Info("starting grpc server")
		// Start listening on the specified port
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", opts.GRPCPort))
		if err != nil {
			errC <- err
		}

		reflection.Register(s)

		err = s.Serve(lis)
		if err != nil {
			errC <- err
		}
	}()

	go func() {
		logger.Info("starting http server")
		// Start listening on the specified port
		if err := h.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errC <- err
		}
	}()

	select {
	case err := <-errC:
		logger.Error(err.Error())
		logger.Info("shutting down...")
	case <-quit:
		cancel()
		logger.Info("interrupt signal received")
		logger.Info("shutting down...")
	}
}
