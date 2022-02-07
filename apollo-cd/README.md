# apollo-cd
`apollo-cd` is a daemon that performs rudimentary continuous deployment.

## Usage
It must be run under root

It expects a JSON configuration as the first argument. The JSON object should have the following key/value pairs

| Key                         | Usage                                                               | Example value                           |
|-----------------------------|---------------------------------------------------------------------|-----------------------------------------|
| `NewAppDeployment`          | A file location for new application deployment **gzipped** tarballs | `/home/apollo/apollo-cd/new.tar.gz`     |
| `DeploymentSystemdServices` | List of systemd services that runs the application                  | `["apollo-app", "apollo-background"]`   |
| `DeploymentDirectory`       | A directory that stores code for the running application            | `/home/apollo/app`                      |
| `DeploymentDirectoryOwner`  | `DeploymentDirectory`'s owning user and group                       | `{"User": "apollo", "Group": "apollo"}` |

## Functionality
It does the following things in a loop:
* Polls `NewAppDeployment`
* If there is a tarball
  * Stop `DeploymentSystemdServices`
  * Delete `DeploymentDirectory`
  * Untar the tarball to `DeploymentDirectory`
  * Own `DeploymentDirectory` to `DeploymentDirectoryOwner`
  * Start `DeploymentSystemdServices`
  * Delete `NewAppDeployment`
