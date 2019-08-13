package config

import (
	"fmt"
)

type UserInfo struct {
	User     string `json:"user"`
	Password string `json:"password"`
}
type Service struct {
	Host string `json:"host"`
	Port uint16 `json:"port"`
}

//Database ... mysql DB
type Database struct {
	Service
	UserInfo
	Name        string `json:"name"`
	MaxIdleTime int    `json:"max_idle_time"` //ms
	MaxIdle     int    `json:"max_idle"`
	Pool        int    `json:"pool"`
	MaxOverflow int    `json:"max_overflow"`
}

func (s *Service) GetPort() string {
	return fmt.Sprintf("%d", s.Port)
}

func (s *Service) GetHostPort() string {
	return s.Host + ":" + fmt.Sprintf("%d", s.Port)
}

type EtcdService struct {
	Service
}

type JenkinsService struct {
	Service
	UserInfo
}

type K8sService struct {
	UserInfo
}

type RedisService struct {
	Service
}

type ELKService struct {
	UserInfo
	Service
	IndexMaxResultWindow int //elksearch  index.max_result_window
}

type SonarService struct {
	Service
	UserInfo
}
