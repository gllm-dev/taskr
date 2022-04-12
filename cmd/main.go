package main

import (
	"log"

	"go.gllm.dev/taskr/cmd/cli"
	"go.gllm.dev/taskr/repo/taskrrepo"
	"go.gllm.dev/taskr/service/taskrsrv"
)

func main() {
	repo, err := taskrrepo.New()
	if err != nil {
		log.Fatal(err)
	}

	service := taskrsrv.New(repo)
	rootCmd := cli.NewCmdRoot(service)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
