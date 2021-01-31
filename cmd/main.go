package main

import (
	"context"
	"fmt"
	"github.com/boshnyakovich/news-aggregator/config"
	"github.com/boshnyakovich/news-aggregator/internal/handlers"
	"github.com/boshnyakovich/news-aggregator/internal/repository"
	"github.com/boshnyakovich/news-aggregator/internal/service"
	"github.com/boshnyakovich/news-aggregator/pkg/fasthttpserver"
	"github.com/boshnyakovich/news-aggregator/pkg/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const dbType = "postgres"

func main() {
	var (
		conf config.Config
		quit = make(chan os.Signal, 1)
	)

	if err := conf.Parse(); err != nil {
		log.Fatalf("error parsing config: %s", err.Error())
	}

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	log, err := initLogger(conf.App.Name, conf.Logger, conf.App.Debug)
	if err != nil {
		log.Fatal("error initializing logger", err)
	}

	log.Debugf("%+v", conf)

	db, err := initDB(conf.Database)
	if err != nil {
		log.Fatal("error initializing db", err)
	}

	defer db.Close()

	repo := repository.NewRepo(db, log)
	service := service.NewService(repo, log)
	handlers := handlers.NewHandlers(service, db, log)

	server, err := fasthttpserver.New().
		WithConfig(&fasthttpserver.Config{
			ReadBufferSize:  4 * 1024,
			WriteBufferSize: 4 * 1024,
			Addr:            fmt.Sprintf(":%d", conf.HTTPConfig.Port),
		}).
		WithLivenessHandler("/health", handlers.LivenessHandler).
		Build()
	if err != nil {
		log.Fatalf("error while init server: %v \n", err)
	}

	server.Router().GET("/habr", handlers.GetTopHabrNews)
	server.Router().GET("/four_pda", handlers.GetFourPDANewsByDate)

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		if err := server.Start(ctx); err != nil {
			cancel()
		}
	}()

	select {
	case <-ctx.Done():
		log.Info("context was canceled")
	case s := <-quit:
		log.Info("signal was provided: ", s)
		cancel()
	}

	wg.Wait()

	log.Info("news-aggregator is stopped")

}

func initDB(conf config.Database) (*sqlx.DB, error) {
	db, err := sqlx.Connect(dbType, fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Name))
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(conf.MaxLifetime)
	db.SetMaxOpenConns(conf.MaxConns)

	return db, nil
}

func initLogger(appName string, conf config.Logger, debug bool) (*logger.Logger, error) {
	fields := make(map[string]interface{})

	if !debug {
		fields = map[string]interface{}{
			"servicename": appName,
			"pods":        conf.PodName,
			"kubenode":    conf.PodNode,
			"namespace":   conf.PodNamespace,
		}
	}

	log := logger.New(fields, conf.Address, conf.Level, debug)

	if !debug {
		if err := log.Connect(); err != nil {
			return nil, err
		}
	}

	return log, nil
}
