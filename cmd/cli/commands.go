package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"go.gllm.dev/trackr/ports"
)

func NewCmdAdd(s ports.Service) *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "add task",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("no args provided")
			} else if len(args) == 1 {
				return s.AddTask(args[0])
			}
			return s.AddTask(args[0], args[1:]...)
		},
	}
}

func NewCmdPause(s ports.Service) *cobra.Command {
	return &cobra.Command{
		Use:   "pause",
		Short: "pause task",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("no args provided")
			}
			return s.PauseTask(args[0])
		},
	}
}

func NewCmdResume(s ports.Service) *cobra.Command {
	return &cobra.Command{
		Use:   "resume",
		Short: "resume task",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("no args provided")
			}
			return s.ResumeTask(args[0])
		},
	}
}

func NewCmdFinish(s ports.Service) *cobra.Command {
	return &cobra.Command{
		Use:   "finish",
		Short: "finish task",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("no args provided")
			}
			return s.FinishTask(args[0])
		},
	}
}

func NewCmdGet(s ports.Service) *cobra.Command {
	return &cobra.Command{
		Use:   "get",
		Short: "get task",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("no args provided")
			}
			t, err := s.GetTask(args[0])
			if err != nil {
				return err
			}
			cmd.Println(t.Fmt())
			return nil
		},
	}
}

func NewCmdList(s ports.Service) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "list tasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			t, err := s.ListTasks()
			if err != nil {
				return err
			}
			for _, tt := range t {
				cmd.Println(tt.Fmt())
			}
			return nil
		},
	}
}
