#!/bin/bash

set -ex

sd "@BFC_APP_DOMAIN@" "$BFC_APP_DOMAIN" go.mod
sd "@BFC_APP_DOMAIN@" "$BFC_APP_DOMAIN" handlers.go
sd "@BFC_APP_DOMAIN@" "$BFC_APP_DOMAIN" main.go 
sd "@BFC_APP_DOMAIN@" "$BFC_APP_DOMAIN" ./templates/index.gohtml
