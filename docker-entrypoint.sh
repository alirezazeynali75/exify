#!/bin/bash

set -e

echo "Migrating Database..."
./exify-cli migrate up

exec ./exify