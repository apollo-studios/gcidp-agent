package traefik

import (
	"fmt"
	"strings"
)

const (
	True    = "true"
	Enable  = "traefik.enable"
	Network = "traefik.docker.network"
)

func Rule(routerName string) string {
	return fmt.Sprintf("traefik.http.routers.%s.rule", strings.ToLower(routerName))
}

func Host(host string) string {
	return fmt.Sprintf("Host(`%s`)", strings.ToLower(host))
}

func TLS(routerName string) string {
	return fmt.Sprintf("traefik.http.routers.%s.tls", strings.ToLower(routerName))
}

func TLSResolver(routerName string) string {
	return fmt.Sprintf("traefik.http.routers.%s.tls.certresolver", strings.ToLower(routerName))
}

func DefineService(routerName string) string {
	return fmt.Sprintf("traefik.http.routers.%s.service", strings.ToLower(routerName))
}

func LoadBalancerPort(routerName string) string {
	return fmt.Sprintf("traefik.http.services.%s.loadbalancer.server.port", strings.ToLower(routerName))
}

func Wildcard(rn, t string) string {
	return fmt.Sprintf("traefik.http.routers.%s.tls.domains[0].%s", strings.ToLower(rn), t)
}
