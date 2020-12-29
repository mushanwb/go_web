package middlewares

import "net/http"

type HttpHandlerFunc func(http.ResponseWriter, *http.Request)
