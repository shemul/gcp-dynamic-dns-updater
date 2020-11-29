#!/bin/bash
set -e

readonly USAGE="bash <project_path/build.sh -v 1.0.X"


while getopts ":v:" opt
do
    case $opt in
        v) readonly VERSION=$OPTARG
        ;;
         \?) echo "Invalid option : $OPTARG"
            echo "See Usage: $USAGE"
            echo "Provided that the command must be run from the project root directory"
            exit 1
        ;;
    esac
done

function error() {
    echo "$(date +'%Y-%m-%dT%H:%M:%S%z') ERROR $@" >&2
    return 1
}


function validate() {
    if [ -z "$VERSION" ]; then
        error "Version must be provided. It can not be empty"
        exit 1
    fi
}

function build() {
    export GOARCH="amd64" GOOS="linux" CGO_ENABLED=0

    if [ -z "$GOPATH" ]; then
       echo "set GOPATH"
       exit
    fi

    hash=$(git rev-parse --short HEAD)
    branch=$(git rev-parse --abbrev-ref HEAD)
    version=$1
    echo 'hash='$hash 'branch='$branch 'version='$version

    go build -o app -v -ldflags="-X main.version=$hash" .
    docker build -t shemul/gcp-dynamic-dns-updater .
    rm ./app
}

validate
build