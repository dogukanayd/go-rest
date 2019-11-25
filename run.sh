#!/usr/bin/env bash

cd "${APP_PATH}" || /
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

dep ensure

chmod 777 -R vendor

go run app/main.go
