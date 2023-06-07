#!/usr/bin/env bash

set -eu

add-apt-repository ppa:longsleep/golang-backports
apt-get update
apt-get install make golang-go=2:1.20~1longsleep1 cgroup-tools -y
