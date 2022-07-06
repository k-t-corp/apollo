* Create the DigitalOcean droplet
	* Use Docker community image
	* Region: SFO or SGP
	* Enable backups if the app has a database with source of truth data
	* Enable monitoring
	* Enable IPv6
	* Tag with `monitoring`. This tag is used for CPU/memory/disk over usage alerting to Slack

* [How to Install the DigitalOcean Metrics Agent](https://docs.digitalocean.com/products/monitoring/how-to/install-agent/)
	* `curl -sSL https://repos.insights.digitalocean.com/install.sh | sudo bash`
	* Verify on DigitalOcean dashboard that detailed metrics graphs are present

* [Initial Server Setup with Ubuntu 20.04](https://www.digitalocean.com/community/tutorials/initial-server-setup-with-ubuntu-20-04)
	* `adduser app`
	* `usermod -aG sudo app`
	* `rsync --archive --chown=app:app ~/.ssh /home/app`
    * Remove open Docker ports
		* `ufw status numbered`
		* `ufw delete 2`
		* `ufw status numbered`
		* `ufw delete 2`
		* `ufw status numbered`
		* `ufw delete 3`
		* `ufw status numbered`
		* `ufw delete 3`
        * `ufw enable`
		* `ufw status`
	* Remove login banner
		* `rm -rf /etc/update-motd.d/99-one-click`
	* For non-US region servers such as SGP, install mosh
        `sudo apt install mosh`
        `sudo ufw allow 60000:61000/udp`
	* Try ssh again using the `app` user

* [Manage Docker as a non-root user](https://docs.docker.com/engine/install/linux-postinstall/#manage-docker-as-a-non-root-user)
	* `sudo groupadd docker`
	* `sudo usermod -aG docker $USER`
	* Log out and log in
	* `docker run hello-world`
    * `sudo apt-get install -y haveged` to prevent `docker-compose` from being stuck

* (Optional) [Install MongoDB Community Edition (5.0) on Ubuntu](https://docs.mongodb.com/manual/tutorial/install-mongodb-on-ubuntu/)
	* `wget -qO - https://www.mongodb.org/static/pgp/server-5.0.asc | sudo apt-key add -`
	* `echo "deb [ arch=amd64,arm64 ] https://repo.mongodb.org/apt/ubuntu focal/mongodb-org/5.0 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-5.0.list`
	* `sudo apt update`
	* `sudo apt-get install -y mongodb-org`
	* `sudo systemctl daemon-reload`
	* `sudo systemctl enable mongod`
	* `sudo systemctl start mongod`
	* `sudo systemctl status mongod`

* (Optional) [How to Install and Configure Redis on Ubuntu 20.04](https://linuxize.com/post/how-to-install-and-configure-redis-on-ubuntu-20-04/)
	* `sudo apt update`
	* `sudo apt install -y redis-server`
	* `sudo systemctl daemon-reload`
	* `sudo systemctl enable redis-server`
	* `sudo systemctl start redis-server`
	* `sudo systemctl status redis-server`

* (Optional) [How To Install PostgreSQL 13 on Ubuntu 20.04](https://computingforgeeks.com/how-to-install-postgresql-13-on-ubuntu/)
	* `wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add -`
	* `echo "deb http://apt.postgresql.org/pub/repos/apt/ `lsb_release -cs`-pgdg main" | sudo tee /etc/apt/sources.list.d/pgdg.list`
	* `sudo apt update`
	* `sudo apt install -y postgresql-13 postgresql-client-13`

* Create env
	* `nano /home/app/env`

* (Optional) Generate files with `python-devkit` for proprietary Python applications

* (Optional) Setup `apollo-cd` to continuously deploy proprietary Python applications
	* Find [latest version](https://github.com/k-t-corp/apollo/releases)
	* `curl -LO https://github.com/k-t-corp/apollo/releases/download/v$VERSION/apollo_$VERSION_linux_amd64.tar.gz`
	* `tar -xzvf apollo_$VERSION_linux_amd64.tar.gz`
	* `sudo mv apollo /usr/local/bin/apollo-cd`
	* `rm apollo_*_linux_amd64.tar.gz README.md`
	* `nano /home/app/apollo-cd.json`
    ```
    {
        "NewAppDeployment": "/home/app/new.tar.gz",
        "StopDeploymentScript": "/home/app/code/stop.sh",
        "StartDeploymentScript": "/home/app/code/start.sh",
        "DeploymentDirectory": "/home/app/code",
        "DeploymentDirectoryOwner": {
            "User": "app",
            "Group": "app"
        }
    }
    ```
    * `sudo nano /etc/systemd/system/apollo-cd.service`
    ```
    [Unit]
    Description=apollo-cd
    After=network.target

    [Service]
    Type=simple
    User=app
    Group=app
    ExecStart=/usr/local/bin/apollo-cd /home/app/apollo-cd.json
    Restart=Always

    [Install]
    WantedBy=multi-user.target
    ```
    * `sudo systemctl daemon-reload`
	* `sudo systemctl enable apollo-cd`
	* `sudo systemctl start apollo-cd`

* (Optional) Setup GitLab CI for proprietary Python applications
	* Add `HOST`, `USER` and `DIR` variables to project's GitLab CI vars
	* Add `deploy.pub` to server's `~/.ssh/authorized_keys`
	* Reboot machine for `journalctl` to show logs
	* Try to push a small change and verify on server apollo-cd works
		* `sudo journalctl -u apollo-cd -f`

* (Optional) Setup `watchtower` to continuously deploy self-hosted applications
    * `mkdir -p /home/app/watchtower`
	* `nano /home/app/watchtower/docker-compose.yml`
    ```
    services:
    watchtower:
        image: containrrr/watchtower
        restart: always
        volumes:
        - /var/run/docker.sock:/var/run/docker.sock
        environment:
        - WATCHTOWER_POLL_INTERVAL=<seconds>
    ```
	* `docker-compose up -d`

* [Install Caddy](https://caddyserver.com/docs/install#debian-ubuntu-raspbian)
	* `sudo apt install -y debian-keyring debian-archive-keyring apt-transport-https`
	* `curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/gpg.key' | sudo tee /etc/apt/trusted.gpg.d/caddy-stable.asc`
	* `curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/debian.deb.txt' | sudo tee /etc/apt/sources.list.d/caddy-stable.list`
	* `sudo apt update`
	* `sudo apt install caddy`
	* `sudo ufw allow proto tcp from any to any port 80,443`
	* `sudo vi /etc/caddy/Caddyfile`
    ```
    <domain>

    reverse_proxy 127.0.0.1:<port>
    ```
    * `sudo systemctl enable caddy.service`
    * Point DNS
    * `sudo systemctl start caddy.service`
    * Verify can ssh into machine using DNS name
	* Verify that the application is accessible using DNS name

* (Optional) Configure Tailscale to allow direct DB access and/or cross-machine communication
	* [Install Tailscale on Ubuntu 20.04](https://tailscale.com/download/linux)
		* `sudo tailscale up` could hang for a while. Be patient.
	* Mark the machine as "No expiry" in Tailscale admin console
	* `sudo ufw allow in on tailscale0`
	* `sudo ufw status`
	* To make MongoDB accessible on tailnet
		* `sudo nano /etc/mongod.conf`
			* Change `bindIp: 127.0.0.1` to `bindIp: 127.0.0.1,<tailscale IP>`
		* Make MongoDB systemd depend on Tailscale interface up
			* `sudo systemctl edit mongod.service`
			```
			[Unit]
			After=tailscaled.service
			Requires=tailscaled.service

			[Service]
			ExecStartPre=/usr/lib/systemd/systemd-networkd-wait-online --interface=tailscale0
			```

* Do a final reboot to make sure everything works after a crash
