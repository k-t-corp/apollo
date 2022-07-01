package main

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"os"
)

func loop(NewAppDeployment string, stopDeploymentScript, startDeploymentScript, DeploymentDirectory string, uid, gid int) error {
	if _, err := os.Stat(NewAppDeployment); errors.Is(err, os.ErrNotExist) {
		log.Infof("New application deployment %s does not exist\n", NewAppDeployment)
		return nil
	} else if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	if err := executeShellScript(stopDeploymentScript, DeploymentDirectory); err != nil {
		log.Infof("Failed to execute stop deployment script %s\n", stopDeploymentScript)
		return err
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

	if err := executeShellScript(startDeploymentScript, DeploymentDirectory); err != nil {
		log.Infof("Failed to execute start deployment script %s\n", startDeploymentScript)
		return err
	}

	if err := os.Remove(NewAppDeployment); err != nil {
		return err
	}

	log.Infoln("Deployment finished")
	return nil
}
