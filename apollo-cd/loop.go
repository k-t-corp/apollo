package main

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"os"
)

func loop(NewAppDeployment string, deploymentSystemdServices []string, DeploymentDirectory string, uid, gid int) error {
	if _, err := os.Stat(NewAppDeployment); errors.Is(err, os.ErrNotExist) {
		log.Infof("New application deployment %s does not exist\n", NewAppDeployment)
		return nil
	} else if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	for _, u := range deploymentSystemdServices {
		if err := systemctlStop(u); err != nil {
			log.Infof("Failed to stop systemd service %s\n", u)
			return err
		}
	}

	if err := os.RemoveAll(DeploymentDirectory); err != nil {
		return err
	}

	if err := untargz(NewAppDeployment, DeploymentDirectory); err != nil {
		log.Infof("Failed to untar new application deployment %s to %s\n", NewAppDeployment, DeploymentDirectory)
		return err
	}

	if err := chownR(DeploymentDirectory, uid, gid); err != nil {
		log.Infof("Failed to own %s to %d,%d\n", DeploymentDirectory, uid, gid)
		return err
	}

	for _, u := range deploymentSystemdServices {
		if err := systemctlStart(u); err != nil {
			log.Infof("Failed to start systemd service %s\n", u)
			return err
		}
	}

	if err := os.Remove(NewAppDeployment); err != nil {
		return err
	}

	log.Infoln("Deployment finished")
	return nil
}
