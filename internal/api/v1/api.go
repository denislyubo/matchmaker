package v1

import (
	"context"
	"net"
	"net/http"

	schema "github.com/denislyubo/matchmaker"
)

type api struct {
	mux *http.ServeMux
	ctl Controller
}

func New(
	matchService schema.MatchService,
	ctl Controller,

) *api {

	s := api{ctl: ctl, mux: http.NewServeMux()}

	s.applyRoutes()

	return &s
}

func (s *api) Name() string { return "api-v1-match-server" }

func (s *api) Start(ctx context.Context) error {
	errChan := make(chan error, 1)
	var l net.Listener
	go func() {
		var err error
		if l, err = net.Listen("tcp", ":8090"); err != nil {
			errChan <- err
		}
		if err = http.Serve(l, s.mux); err != nil {
			errChan <- err
		}
	}()

	defer func() {
		if l != nil {
			l.Close()
		}
	}()

	select {
	case <-ctx.Done():
		return nil
	case err := <-errChan:
		return err
	}
}

func (s *api) applyRoutes() {
	s.mux.HandleFunc("POST /api/v1/users", s.ctl.AddUserHandler)
	s.mux.HandleFunc("GET /api/v1/match", s.ctl.GetMatchHandler)
}
