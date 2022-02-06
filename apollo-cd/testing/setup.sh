#!/usr/bin/env bash
set -e

# NewAppDeployment
mkdir -p /home/parallels/apollo-cd/app
tee /home/parallels/apollo-cd/app/main.py > /dev/null <<EOT
import time

while True:
    print('dummy service. i am alive')
    time.sleep(1)
EOT
pushd /home/parallels/apollo-cd
tar -czvf new.tar.gz -C app .
popd

# DeploymentSystemdServices
sudo tee /etc/systemd/system/demo.service > /dev/null <<EOT
[Unit]
Description=demo service
After=network.target

[Service]
Type=simple
User=parallels
Group=parallels
WorkingDirectory=/home/parallels/apollo-app
ExecStart=/usr/bin/python3 /home/parallels/apollo-app/main.py
Restart=Always

[Install]
WantedBy=multi-user.target
EOT
sudo systemctl daemon-reload
sudo systemctl enable demo.service
sudo systemctl start demo.service

# DeploymentDirectory
mkdir -p /home/parallels/apollo-app
