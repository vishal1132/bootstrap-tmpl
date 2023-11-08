package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"{{ module_name }}/config"
	common_http "{{ module_name }}/http"
	"{{ module_name }}/observability/log"
	{% if redis_enabled %}
	"{{ module_name }}/redis"
	{% endif %}
	"{{ module_name }}/status"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	conf := config.LoadConfig()

	router := common_http.GetRouter()

	log.ReplaceDefaultLogger(log.WithLevel(log.LevelFromString(conf.LogConfig.Level)),
		log.WithNamespace(conf.LogConfig.Namespace),
	)

	// init dependencies here.
	{% if db_enabled %}
	db := postgres.InitDatabase(ctx, conf.DBConfig)
	{% endif %}

	{% if redis_enabled %}
	cache := redis.InitRedis(ctx, conf.RedisConfig)
	_ = redis.NewRedisDaoImpl(cache)
	{% endif %}
	// Routes here.
	status.NewStatusController(router, 
		{% if db_enabled %}
		status.NewPGMonitor("maindb", db),
		{% endif %}
		{% if redis_enabled %}
		status.NewRedisMonitor("maincache", cache),
		{% endif %}
	)

	srv := common_http.InitServer(router, conf.AppConfig.ServerWriteTimeout, conf.AppConfig.ServerReadTimeout)
	log.Info(ctx, "Starting server")
	go srv.ListenAndServe()
	waitForShutdownSignal(ctx, conf.AppConfig.GracefulShutDownTimeout, srv, cancel)
}

func waitForShutdownSignal(ctx context.Context, gracefulShutDownTimeout int64, srv *http.Server, cancel context.CancelFunc) {
	var gracefulStop = make(chan os.Signal, 1)

	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	signal.Notify(gracefulStop, syscall.SIGQUIT)

	select {
	case <-gracefulStop:
		cancel()
		// if stop signal is received, wait for some time so that background workers get time to exit
		<-time.After(time.Duration(gracefulShutDownTimeout) * time.Millisecond)
	case <-ctx.Done():
		// shutdown if context was cancelled by something else before shutdown signal
	}
	srv.Shutdown(ctx)
}
