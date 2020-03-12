package main

import (
	"errors"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/koyuta/manifest-updater/pkg/registry"
	"github.com/koyuta/manifest-updater/pkg/repository"
	"github.com/urfave/cli"
)

func execute(c *cli.Context) error {
	if c.NArg() != 0 {
		if err := cli.ShowAppHelp(c); err != nil {
			return err
		}
		// Return empty error to set 1 to exit status.
		return errors.New("")
	}

	var stop = make(chan struct{})
	go func() {
		sigch := make(chan os.Signal, 1)
		signal.Notify(sigch, syscall.SIGTERM)
		<-sigch
		stop <- struct{}{}
	}()

	checkInterval := 1 * time.Minute

	updater := NewUpdater(
		registry.NewDockerHubRegistry(
			c.String(registryDockerHubFlag.Name),
			c.String(registryFilterFlag.Name),
		),
		repository.NewGitHubRepository(
			c.String(repositoryGitFlag.Name),
			c.String(repositoryBranchFlag.Name),
			c.String(repositoryPathFlag.Name),
			c.String(registryDockerHubFlag.Name),
		),
	)

	looper := NewUpdateLooper(updater, checkInterval)
	looper.Loop(stop)

	return nil
}
