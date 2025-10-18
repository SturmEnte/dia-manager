#!/usr/bin/env bash
set -euo pipefail

: "${BACKEND_HOST:=backend}"
: "${BACKEND_PORT:=8369}"

envsubst '${BACKEND_HOST} ${BACKEND_PORT}' \
  < /etc/nginx/conf.d/default.conf.template \
  > /etc/nginx/conf.d/default.conf

exec "$@"