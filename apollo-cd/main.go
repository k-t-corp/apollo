package main

import (
	"encoding/json"
	"flag"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	user2 "os/user"
	"strconv"
	"time"
)

type ConfigOwner struct {
	User  string `json:"User"`
	Group string `json:"Group"`
}

type Config struct {
	NewAppDeployment         string      `json:"NewAppDeployment"`
	StopDeploymentScript     string      `json:"StopDeploymentScript"`
	StartDeploymentScript    string      `json:"StartDeploymentScript"`
	DeploymentDirectory      string      `json:"DeploymentDirectory"`
	DeploymentDirectoryOwner ConfigOwner `json:"DeploymentDirectoryOwner"`
}

func parseOwner(username, groupName string) (int, int, error) {
	user, err := user2.Lookup(username)
	if err != nil {
		return -1, -1, err
	}
	uid, err := strconv.Atoi(user.Uid)
	if err != nil {
		return -1, -1, err
	}
	group, err := user2.LookupGroup(groupName)
	if err != nil {
		return -1, -1, err
	}
	gid, err := strconv.Atoi(group.Gid)
	if err != nil {
		return -1, -1, err
	}
	return uid, gid, nil
}

func main() {
	flag.Parse()

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
	StopDeploymentScript := config.StopDeploymentScript
	StartDeploymentScript := config.StartDeploymentScript
	DeploymentDirectory := config.DeploymentDirectory
	deploymentDirectoryUser := config.DeploymentDirectoryOwner.User
	deploymentDirectoryGroup := config.DeploymentDirectoryOwner.Group
	uid, gid, err := parseOwner(deploymentDirectoryUser, deploymentDirectoryGroup)
	if err != nil {
		log.Errorln(err)
		return
	}

	log.Infoln("Configurations")
	log.Infof("NewAppDeployment=%s", NewAppDeployment)
	log.Infof("StopDeploymentScript=%s", StopDeploymentScript)
	log.Infof("StartDeploymentScript=%s", StartDeploymentScript)
	log.Infof("DeploymentDirectory=%s", DeploymentDirectory)
	log.Infof("deploymentDirectoryUser=%s", deploymentDirectoryUser)
	log.Infof("deploymentDirectoryGroup=%s", deploymentDirectoryGroup)

	for {
		log.Infoln("--- Running loop ---")
		if err := loop(NewAppDeployment, StopDeploymentScript, StartDeploymentScript, DeploymentDirectory, uid, gid); err != nil {
			log.Errorln(err)
			return
		}

		time.Sleep(10 * time.Second)
	}
}
