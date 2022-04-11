package main

import (
	"go.gllm.dev/trackr/cmd/cli"
	"go.gllm.dev/trackr/repo/taskrrepo"
	"go.gllm.dev/trackr/service/taskrsrv"
	"log"
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
