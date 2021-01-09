#!/usr/bin/env bash
echo "Building linux..."
env GOOS=linux GOARCH=386 go build -o ./dist/valorant-linux
echo "Building windows..."
env GOOS=windows GOARCH=386 go build -o ./dist/valorant.exe
echo "Building macos..."
env GOOS=darwin GOARCH=386 go build -o ./dist/valorant-macos