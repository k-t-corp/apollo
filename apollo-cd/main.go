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
	deploymentDirectoryUid, deploymentDirectoryGid, err := parseOwner(config.DeploymentDirectoryOwner.User, config.DeploymentDirectoryOwner.Group)
	if err != nil {
		log.Errorln(err)
		return
	}

	log.Infoln("Configurations")
	log.Infof("NewAppDeployment=%s", NewAppDeployment)
	log.Infof("DeploymentSystemdServices=%s", DeploymentSystemdServices)
	log.Infof("DeploymentDirectory=%s", DeploymentDirectory)
	log.Infof("deploymentDirectoryUid=%d", deploymentDirectoryUid)
	log.Infof("deploymentDirectoryGid=%d", deploymentDirectoryGid)

	for {
		log.Infoln("--- Running loop ---")
		if err := loop(NewAppDeployment, DeploymentSystemdServices, DeploymentDirectory, deploymentDirectoryUid, deploymentDirectoryGid); err != nil {
			log.Errorln(err)
			return
		}

		time.Sleep(10 * time.Second)
	}
}
