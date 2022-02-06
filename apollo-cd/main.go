package main

import (
	"encoding/json"
	"flag"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	user2 "os/user"
	"time"
)

func isRoot() bool {
	user, err := user2.Current()
	if err != nil {
		log.Errorln(err)
		return false
	}
	return user.Name == "root"
}

type ConfigOwner struct {
	User  string `json:"User"`
	Group string `json:"Group"`
}

type Config struct {
	NewAppDeployment          string      `json:"NewAppDeployment"`
	DeploymentSystemdServices []string    `json:"DeploymentSystemdServices"`
	DeploymentDirectory       string      `json:"DeploymentDirectory"`
	DeploymentDirectoryOwner  ConfigOwner `json:"DeploymentDirectoryOwner"`
}

func main() {
	flag.Parse()

	if !isRoot() {
		log.Errorln("apollo-cd must be run under root")
		return
	}

	if len(os.Args) != 2 {
		log.Errorln("Usage: apollo-cd config.json")
		return
	}

	configBytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Errorln(err)
		return
	}
	var config Config
	if err := json.Unmarshal(configBytes, &config); err != nil {
		log.Errorln(err)
		return
	}

	NewAppDeployment := config.NewAppDeployment
	DeploymentSystemdServices := config.DeploymentSystemdServices
	DeploymentDirectory := config.DeploymentDirectory
	deploymentDirectoryUser := config.DeploymentDirectoryOwner.User
	deploymentDirectoryGroup := config.DeploymentDirectoryOwner.Group
	if err != nil {
		log.Errorln(err)
		return
	}

	log.Infoln("Configurations")
	log.Infof("NewAppDeployment=%s", NewAppDeployment)
	log.Infof("DeploymentSystemdServices=%s", DeploymentSystemdServices)
	log.Infof("DeploymentDirectory=%s", DeploymentDirectory)
	log.Infof("deploymentDirectoryUser=%s", deploymentDirectoryUser)
	log.Infof("deploymentDirectoryGroup=%s", deploymentDirectoryGroup)

	for {
		log.Infoln("--- Running loop ---")
		if err := loop(NewAppDeployment, DeploymentSystemdServices, DeploymentDirectory, deploymentDirectoryUser, deploymentDirectoryGroup); err != nil {
			log.Errorln(err)
			return
		}

		time.Sleep(10 * time.Second)
	}
}
