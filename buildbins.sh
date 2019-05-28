#!/bin/bash

# GOOS=windows GOARCH=amd64 go build main.go
# mv main bin/windows/sauceutil
# file bin/windows/sauceutil

echo "Building linux bin..."
GOOS=linux GOARCH=amd64 go build main.go
mv main bin/linux/sauceutil
file bin/linux/sauceutil
echo ""
echo "Building macos/darwin bin..."
GOOS=darwin GOARCH=amd64 go build main.go
mv main bin/macos/sauceutil
file bin/macos/sauceutil
echo ""