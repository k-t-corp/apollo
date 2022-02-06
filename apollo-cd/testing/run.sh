#!/usr/bin/env bash
set -ex

pushd ..
sudo go run . ./testing/config.json
popd
