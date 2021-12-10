package server

import (
	"blog/core/setting"
	"fmt"
	"net/http"
	"time"
)

func NewServer(setting *setting.Server, routes http.Handler) (s *http.Server) {
	h := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.Port),
		Handler:        routes,
		ReadTimeout:    setting.ReadTimeout * time.Second,
		WriteTimeout:   setting.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return h
}
