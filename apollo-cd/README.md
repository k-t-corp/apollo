# apollo-cd
`apollo-cd` is a daemon that performs rudimentary continuous deployment.

## Usage
It expects a JSON configuration as the first argument. The JSON object should have the following key/value pairs

| Key                        | Usage                                                               | Example value                           |
|----------------------------|---------------------------------------------------------------------|-----------------------------------------|
| `NewAppDeployment`         | A file location for new application deployment **gzipped** tarballs | `/home/apollo/apollo-cd/new.tar.gz`     |
| `StopDeploymentScript`     | A file location for shell script that stops the deployment          | `/home/apollo/app/stop.sh`              |
| `StartDeploymentScript`    | A file location for shell script that starts the deployment         | `/home/apollo/app/start.sh`             |
| `DeploymentDirectory`      | A directory that stores code for the running application            | `/home/apollo/app`                      |
| `DeploymentDirectoryOwner` | `DeploymentDirectory`'s owning user and group                       | `{"User": "apollo", "Group": "apollo"}` |

## Functionality
It does the following things in a loop:
* Polls `NewAppDeployment`
* If there is a tarball
  * Run `StopDeploymentScript`
  * Delete `DeploymentDirectory`
  * Untar the tarball to `DeploymentDirectory`
  * Own `DeploymentDirectory` to `DeploymentDirectoryOwner`
  * Run `StartDeploymentScript`
  * Delete `NewAppDeployment`
