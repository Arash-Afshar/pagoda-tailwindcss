#!/bin/bash

# This script is meant to be run inside the docker container
set -e

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

cd $SCRIPT_DIR/..

set -a
source .env

make setup
make dev-css & make dev

