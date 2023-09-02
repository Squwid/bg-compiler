package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/Squwid/bg-compiler/docker"
	"github.com/Squwid/bg-compiler/entry"
	"github.com/Squwid/bg-compiler/processor"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	docker.Init()
	processor.InitWorkers()

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(200) }).Methods("GET")
	r.HandleFunc("/compile", compileHandler).Methods("POST")
	logrus.Fatalln(http.ListenAndServe(":8080", r))
}

func compileHandler(w http.ResponseWriter, r *http.Request) {
	logger := logrus.WithField("Action", "CompileHandler")

	bs, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Errorf("Error reading body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var sub *entry.Submission
	if err := json.Unmarshal(bs, &sub); err != nil {
		logger.Errorf("Error unmarshalling submission: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := processor.ProcessSubmission(context.Background(), sub); err != nil {
		logrus.WithError(err).Errorf("Error processing submission")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
