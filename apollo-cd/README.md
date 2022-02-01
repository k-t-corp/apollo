# apollo-cd
`apollo-cd` is a daemon that performs rudimentary continuous deployment.

## Requirements
It expects the following locations/names as environment variables

If any **one** of the environment variables is not provided, the daemon will use the set of default values for **all** environment variables.

| Name                       | Usage                                                               | Default value                       |
|----------------------------|---------------------------------------------------------------------|-------------------------------------|
| `NewAppDeployment`         | A file location for new application deployment **gzipped** tarballs | `/home/apollo/apollo-cd/new.tar.gz` |
| `DeploymentSystemdService` | An **user instance** systemd service that runs the application      | `apollo-app`                        |
| `DeploymentDirectory`      | A directory that stores code for the running application            | `/home/apollo/app`                  |

## Functionality
It does the following things in a loop:
* Polls `NewAppDeployment`
* If there is a tarball
  * Fully stop `DeploymentSystemdService`
  * Delete `DeploymentDirectory`
  * Untar the tarball to `DeploymentDirectory`
  * Start `DeploymentSystemdService`
  * Delete `NewAppDeployment`
