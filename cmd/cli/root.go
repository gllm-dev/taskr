package cli

import (
	"github.com/spf13/cobra"
	"go.gllm.dev/taskr/ports"
)

func NewCmdRoot(service ports.TaskrService) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "taskr",
		Short: "Root command",
	}
	cmd.AddCommand(NewCmdAdd(service))
	cmd.AddCommand(NewCmdGet(service))
	cmd.AddCommand(NewCmdResume(service))
	cmd.AddCommand(NewCmdFinish(service))
	cmd.AddCommand(NewCmdList(service))
	cmd.AddCommand(NewCmdPause(service))
	return cmd
}
