package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	serviceName        = "finance-tracker-api"
	serviceVersion     = "1.0.0"
	serviceDescription = "Personal finance tracker REST API"
	serviceDocs        = "https://github.com/filipbabicdev/finance-tracker-api"
)

type RouteInfo struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}

// RootHandler closes over r so Routes() is evaluated per-request, not at
// registration time -- routes added after this handler is wired up still
// show up in the response.
func RootHandler(r *gin.Engine, env string) gin.HandlerFunc {
	return func(c *gin.Context) {
		ginRoutes := r.Routes()
		routes := make([]RouteInfo, 0, len(ginRoutes))
		for _, rt := range ginRoutes {
			routes = append(routes, RouteInfo{Method: rt.Method, Path: rt.Path})
		}

		c.JSON(http.StatusOK, gin.H{
			"service":     serviceName,
			"version":     serviceVersion,
			"status":      "ok",
			"environment": env,
			"description": serviceDescription,
			"docs":        serviceDocs,
			"timestamp":   time.Now().UTC().Format(time.RFC3339),
			"routes":      routes,
		})
	}
}

func NoRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "route not found"})
	}
}
