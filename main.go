package main

import (
	"github.com/Squwid/bg-compiler/cmd"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := cmd.Execute(); err != nil {
		logrus.WithError(err).Fatalf("Error executing flags")
	}
}
