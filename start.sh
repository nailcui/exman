#!/bin/bash

export GIN_MODE=release
mkdir -p /logs/exman
touch /logs/exman/info.log
nohup ./exman >> /logs/exman/info.log 2>&1 &