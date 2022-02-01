package main

import (
	"errors"
	"flag"
	"github.com/golang/glog"
	"os"
	"time"
)

func loop(NewAppDeployment, DeploymentSystemdService, DeploymentDirectory string) error {
	if _, err := os.Stat(NewAppDeployment); errors.Is(err, os.ErrNotExist) {
		glog.Infof("New application deployment %s does not exist\n", NewAppDeployment)
		return nil
	} else if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	if err := systemctlStop(DeploymentSystemdService); err != nil {
		glog.Infof("Failed to stop systemd service %s\n", DeploymentSystemdService)
		return err
	}

	for i := 0; i < 12; i++ {
		isInactive, err := systemctlIsInactive(DeploymentSystemdService)
		if err != nil {
			return err
		}
		if isInactive {
			break
		}
		time.Sleep(5 * time.Second)
	}

	if err := os.RemoveAll(DeploymentDirectory); err != nil {
		return err
	}

	if err := untargz(NewAppDeployment, DeploymentDirectory); err != nil {
		glog.Infof("Failed to untar new application deployment %s to %s\n", NewAppDeployment, DeploymentDirectory)
		return err
	}

	if err := systemctlStart(DeploymentSystemdService); err != nil {
		glog.Infof("Failed to start systemd service %s\n", DeploymentSystemdService)
		return err
	}

	if err := os.Remove(NewAppDeployment); err != nil {
		return err
	}

	glog.Infoln("Deployment finished")
	return nil
}

func main() {
	flag.Parse()

	NewAppDeployment := os.Getenv("NewAppDeployment")
	DeploymentSystemdService := os.Getenv("DeploymentSystemdService")
	DeploymentDirectory := os.Getenv("DeploymentDirectory")

	if NewAppDeployment == "" || DeploymentSystemdService == "" || DeploymentDirectory == "" {
		glog.Warningln("Overwriting all environment variables with default values")
		NewAppDeployment = "/home/apollo/apollo-cd/new.tar.gz"
		DeploymentSystemdService = "apollo-app"
		DeploymentDirectory = "/home/apollo/app"
	}
	glog.Infof("NewAppDeployment=%s", NewAppDeployment)
	glog.Infof("DeploymentSystemdService=%s", DeploymentSystemdService)
	glog.Infof("DeploymentDirectory=%s", DeploymentDirectory)

	for {
		glog.Infoln("--- Running loop ---")
		if err := loop(NewAppDeployment, DeploymentSystemdService, DeploymentDirectory); err != nil {
			glog.Error(err)
			return
		}

		time.Sleep(10 * time.Second)
	}
}
