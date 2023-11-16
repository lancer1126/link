package service

import (
	"github.com/Masterminds/semver/v3"
	"github.com/alimy/tryst/cfg"
	"log"
)

type Service interface {
	Name() string
	Version() *semver.Version
	OnInit() error
	OnStart() error
	OnStop() error
}

func InitService() []Service {
	ss := startService()
	for _, s := range ss {
		if err := s.OnInit(); err != nil {
			log.Fatalf("Initial %s service error: %s", s.Name(), err)
		}
	}
	return ss
}

func startService() (ss []Service) {
	cfg.In(cfg.Actions{
		"Web": func() {
			ss = append(ss, startWebService())
		},
	})
	return
}
