#!/usr/bin/env bash

# get env vars needed for oauth2_proxy
eval $(./scripts/ups_as_envs.py)

# this is what the real app will listen on
ALT_PORT=8081

# runs in background
./scripts/oauth2_proxy -http-address "0.0.0.0:${PORT}" -upstream "http://localhost:${ALT_PORT}" &

# note non-standard port for underlying app, which should also only listen on 127.0.0.1 (not 0.0.0.0)
PORT="${ALT_PORT}" exec ./bin/main
