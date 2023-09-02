package processor

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Squwid/bg-compiler/entry"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func ProcessSubmission(ctx context.Context, sub *entry.Submission) error {
	logger := logrus.WithFields(logrus.Fields{
		"Action": "Process",
	})

	def, err := NewDefinition(ctx, sub)
	if err != nil {
		return errors.Wrap(err, "Error creating definition")
	}
	def.Run()
	logger.Infof("Finished processing submission")

	var outputs = []JobOutput{}
	for _, job := range def.Jobs {
		outputs = append(outputs, *job.Output)
	}
	bs, err := json.MarshalIndent(outputs, "", "  ")
	if err != nil {
		return errors.Wrap(err, "Error marshalling outputs")
	}
	fmt.Println(string(bs))

	return nil

}
