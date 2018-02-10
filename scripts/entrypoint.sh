#!/bin/bash

set -o pipefail
set -o errtrace
set -o nounset
set -o errexit


if ! [[ -f db.yml ]];then
	echo "Warning: Didn't find a users db (db.yml)"
fi

if [[ -f app.yml ]];then
	./go389 server -c app.yml
else
	./go389 server -c "${APP_CONFIG}"
fi