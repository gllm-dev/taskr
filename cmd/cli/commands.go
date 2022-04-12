package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"go.gllm.dev/taskr/ports"
)

func NewCmdAdd(s ports.TaskrService) *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "add task",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("no args provided")
			} else if len(args) == 1 {
				if err := s.AddTask(args[0]); err != nil {
					fmt.Println(err.Error())
				}
				return nil
			}
			if err := s.AddTask(args[0], args[1:]...); err != nil {
				fmt.Println(err.Error())
			}
			return nil
		},
	}
}

func NewCmdPause(s ports.TaskrService) *cobra.Command {
	return &cobra.Command{
		Use:   "pause",
		Short: "pause task",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("no args provided")
			}
			if err := s.PauseTask(args[0]); err != nil {
				fmt.Println(err.Error())
			}
			return nil
		},
	}
}

func NewCmdResume(s ports.TaskrService) *cobra.Command {
	return &cobra.Command{
		Use:   "resume",
		Short: "resume task",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("no args provided")
			}
			if err := s.ResumeTask(args[0]); err != nil {
				fmt.Println(err.Error())
			}
			return nil
		},
	}
}

func NewCmdFinish(s ports.TaskrService) *cobra.Command {
	return &cobra.Command{
		Use:   "finish",
		Short: "finish task",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("no args provided")
			}
			if err := s.FinishTask(args[0]); err != nil {
				fmt.Println(err.Error())
			}
			return nil
		},
	}
}

func NewCmdGet(s ports.TaskrService) *cobra.Command {
	return &cobra.Command{
		Use:   "get",
		Short: "get task",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("no args provided")
			}
			t, err := s.GetTask(args[0])
			if err != nil {
				fmt.Println(err.Error())
				return nil
			}
			cmd.Println(t.Fmt())
			return nil
		},
	}
}

func NewCmdList(s ports.TaskrService) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "list tasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			t, err := s.ListTasks()
			if err != nil {
				fmt.Println(err.Error())
				return nil
			}
			for _, tt := range t {
				cmd.Println(tt.Fmt())
			}
			return nil
		},
	}
}
