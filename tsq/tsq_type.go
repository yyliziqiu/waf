package tsq

import (
	"strconv"
	"strings"
)

type ServiceConfig struct {
	Name     string
	Protocol string
	Host     string
	Port     int
	Path     string
}

func (s ServiceConfig) ToUrl() string {
	if s.Protocol == "" {
		s.Protocol = "http"
	}
	if s.Host == "" {
		s.Host = "127.0.0.1"
	}
	if s.Port == 0 {
		s.Port = 80
	}

	sb := strings.Builder{}
	sb.WriteString(s.Protocol)
	sb.WriteString("://")
	sb.WriteString(s.Host)
	sb.WriteString(":")
	sb.WriteString(strconv.Itoa(s.Port))
	if s.Path != "" {
		sb.WriteString("/")
		sb.WriteString(strings.TrimLeft(s.Path, "/"))
	}
	return sb.String()
}

func (s ServiceConfig) JoinUrl(postfix string) string {
	return strings.TrimRight(s.ToUrl(), "/") + "/" + strings.TrimLeft(postfix, "/")
}
