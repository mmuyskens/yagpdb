#!/bin/bash
VERSION=$(git describe --tags)
echo Building version $VERSION
go build -ldflags "-X github.com/mmuyskens/yagpdb/common.VERSION=${VERSION}"