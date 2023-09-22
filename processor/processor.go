package processor

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func ProcessSubmission(ctx context.Context, sub *Submission) ([]JobOutput, error) {
	logger := logrus.WithFields(logrus.Fields{
		"Action": "Process",
	})

	def, err := NewDefinition(ctx, sub)
	if err != nil {
		return nil, errors.Wrap(err, "Error creating definition")
	}
	def.Run()
	logger.Infof("Finished processing submission")

	var outputs = []JobOutput{}
	for _, job := range def.Jobs {
		outputs = append(outputs, *job.Output)
	}

	return outputs, nil
}
