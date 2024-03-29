#!/bin/bash
# this builds all the svcs and bundle them into base_image, the each dockerfile will extract from the base_image as an individual svc.
pids=()

make pb &
pids+=( $! )
make gen &
pids+=( $! )

for pid in ${pids[*]}; do
  wait $pid
done

unset pids

services="users versus"

echo "[STARTING] - Building go services"

for service in $services; do
  CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/${service} backend/${service}/main.go &
  pids+=( $! )
done

for pid in ${pids[*]}; do
  wait $pid
done

echo "[FINISHED] - Building go services"
