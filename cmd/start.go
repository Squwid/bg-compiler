package cmd

import (
	"fmt"

	"github.com/Squwid/bg-compiler/docker"
	"github.com/Squwid/bg-compiler/processor"
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

	fmt.Println("Start server!")
}
