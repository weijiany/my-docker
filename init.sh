#!/usr/bin/env bash

set -eu

add-apt-repository ppa:longsleep/golang-backports
apt-get update
apt-get install make golang-go cgroup-tools -y
