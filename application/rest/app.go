package rest

import (
	"antrein/dd-dashboard-analytic/application/common/resource"
	"antrein/dd-dashboard-analytic/application/common/usecase"
	"antrein/dd-dashboard-analytic/internal/handler/grpc/analytic"
	"antrein/dd-dashboard-analytic/model/config"
	"compress/gzip"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func setupCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE, PATCH")
	w.Header().Set("Agccess-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
}

type gzipResponseWriter struct {
	http.ResponseWriter
	Writer *gzip.Writer
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func compressHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}
		w.Header().Set("Content-Encoding", "gzip")
		gz, _ := gzip.NewWriterLevel(w, gzip.BestSpeed)
		defer gz.Close()

		gzw := gzipResponseWriter{ResponseWriter: w, Writer: gz}
		next.ServeHTTP(gzw, r)
	})
}

func ApplicationDelegate(cfg *config.Config, uc *usecase.CommonUsecase, rsc *resource.CommonResource) (http.Handler, error) {
	router := mux.NewRouter()

	router.HandleFunc("/dd/dashboard/analytic/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong!")
	})

	// routes

	// analytic
	analyticRouter := analytic.New(cfg, rsc.GRPC)
	analyticRouter.RegisterRoute(router)

	handlerWithMiddleware := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setupCORS(w)

		if r.Method == "OPTIONS" {
			return
		}

		compressHandler(router).ServeHTTP(w, r)
	})

	return handlerWithMiddleware, nil
}
