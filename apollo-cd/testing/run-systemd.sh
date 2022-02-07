#!/usr/bin/env bash
set -e

sudo tee /etc/systemd/system/apollo-cd.service > /dev/null <<EOT
[Unit]
[Unit]
Description=apollo-cd
After=network.target

[Service]
Type=simple
User=root
Group=root
ExecStart=/usr/local/bin/apollo-cd /home/app/apollo-cd.json
Restart=Always

[Install]
WantedBy=multi-user.target
EOT
sudo systemctl daemon-reload
sudo systemctl enable apollo-cd.service
sudo systemctl start apollo-cd.service
journalctl -u apollo-cd -xf
