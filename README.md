# Apollo
Linux provisioning scripts + application deployment system. Suitable for self-hosting and hobby-scale application deployments.

## Philosophy
* Linux-native: Apollo tools delegate logic to Linux-native tools as much as possible.
* Small and easy to understand: As a collary, apollo tools are small, have minimal logic, and are easy to understand.
* Modular: Apollo tools can each be used as a standalone tool and run independently.
* Collaborating: Apollo tools can also collaborate via small, well-defined and Linux-native interfaces.
* Single-tenant: Apollo tools, even collaborating, only host one logical application per machine. Tools such as Dokku and CapRover focuses on hosting multiple applications per machine, which I think makes their tools more complex, their applications more resource-competing and less fault-tolerant. 
* Self-hosted software friendly: By basing on Linux instead of higher level of abstractions such as Kubernetes, apollo tools are more suitable for self-hosted software which sometimes does not provide Kubernetes or even Docker deployment options.
* Open for customization: By basing on Linux, it is also easier to make ad-hoc customizations to your machines and applications if needed.

## Features
* Linux Provisioning and Security Setup (one-time)
* Linux Monitoring (continuous)
* Application Deployment (one-time) and Continuous Deployment (continuous)
  * Application types
    * Private non-Docker workload
    * Public Docker workload
    * Public non-Docker workload
  * Workload types
    * Web
    * Pre-deploy
    * Background
* Environment and Secret Injection (continuous)
* Logging (continuous)
* Database Provisioning
  * Types
    * MongoDB
    * Postgres
    * Redis
    * ElasticSearch
  * Tasks
    * Provisioning (one-time)
    * Auto injection (continuous)
    * Backup (continuous)
    * Restore (one-time)
* Reverse Proxy Provisioning
  * Types
    * Caddy
    * Nginx
  * Tasks
    * Provisioning (one-time)
* Application Monitoring (continuous)
