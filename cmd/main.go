package main

import (
	"go.gllm.dev/trackr/cmd/cli"
	repo2 "go.gllm.dev/trackr/repo"
	"go.gllm.dev/trackr/services/taskrsrv"
	"log"
)

func main() {
	repo, err := repo2.NewRepo()
	if err != nil {
		log.Fatal(err)
	}

	service := taskrsrv.NewService(repo)
	rootCmd := cli.NewCmdRoot(service)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}