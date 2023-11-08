package status

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"{{ module_name }}/observability/log"
)

type statusHandler struct {
	dependencyMonitors []DependencyMonitor
}

type DependencyMonitor interface {
	GetName() string
	Ping(ctx context.Context) error
}

func NewStatusController(r *chi.Mux, dependencyMonitors ...DependencyMonitor) {
	s := &statusHandler{
		dependencyMonitors: dependencyMonitors,
	}
	r.Get("/status", s.healthCheckHandler)
}

func (s *statusHandler) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("deepcheck") == "" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("imok"))
		return
	}

	dependencyStatus := make(map[string]string, len(s.dependencyMonitors))
	ctx := r.Context()
	for _, monitor := range s.dependencyMonitors {
		if err := monitor.Ping(ctx); err != nil {
			dependencyStatus[monitor.GetName()] = err.Error()
			continue
		}
		dependencyStatus[monitor.GetName()] = "ok"
	}
	json.NewEncoder(w).Encode(dependencyStatus)

}

func logDependencyError(ctx context.Context, dependency, name string, err error) error {
	if err != nil {
		log.Error(ctx, err, "error pinging dependency", zap.String("dependency", dependency),
			zap.String("name", name))
	}
	return err
}
