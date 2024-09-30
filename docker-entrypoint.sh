#!/bin/bash

set -e

touch .env
echo "Migrating Database..."
./exify-cli migrate up
echo "Migration finished ..."
exec ./exify