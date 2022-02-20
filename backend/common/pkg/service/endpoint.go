package service

import (
	"bytes"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	log "github.com/sirupsen/logrus"

	commonerr "github.com/Posrabi/flashy/backend/common/pkg/error"
)

type ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request) error

func ErrorLogger(name string, f ErrorHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		r.Body = io.NopCloser(bytes.NewBuffer(body))

		err := f(w, r)
		if err != nil {
			w.WriteHeader(runtime.HTTPStatusFromCode(commonerr.GetGRPCCode(err)))

			if _, writeErr := io.WriteString(w, err.Error()); writeErr != nil {
				log.Errorf("error writing in errorLogger middleware: %s", err.Error())
			}

			log.WithFields(log.Fields{
				"request":  string(body),
				"endpoint": name,
				"query":    r.URL.Query(),
				"url":      r.URL.Path,
			})
		}
	}
}
