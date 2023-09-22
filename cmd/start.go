package cmd

import (
	"github.com/Squwid/bg-compiler/docker"
	"github.com/Squwid/bg-compiler/processor"
	"github.com/Squwid/bg-compiler/webserver"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const cmdStartDesc = `Start a compiler webserver that takes HTTP requests and 
runs them in a docker container.`

var cmdStart = &cobra.Command{
	Use:   "start",
	Short: "Start a compiler webserver",
	Long:  cmdStartDesc,
	Args:  cobra.NoArgs,
	Run:   startCmd,
}

func startCmd(cmd *cobra.Command, args []string) {
	docker.Init()
	processor.InitWorkers()

	// TODO: Make port configurable.
	server := webserver.NewServer(8080)
	logrus.WithError(server.Start()).Fatalln("Error starting server")
}
