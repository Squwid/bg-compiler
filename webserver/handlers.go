package webserver

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/Squwid/bg-compiler/processor"
	"github.com/sirupsen/logrus"
)

func compileHandler(w http.ResponseWriter, r *http.Request) {
	logger := logrus.WithField("Action", "CompileHandler")

	bs, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Errorf("Error reading body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var sub *processor.Submission
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
