start=`date +%s`

set -eu -o pipefail

pids=()

BUILD_TS=false

for arg in "$@"; do 
  if [ $arg = "ts" ]; then
    BUILD_TS=true
  fi
done

if [ $BUILD_TS = true ]; then
  make ts=1 pb &
else
  make pb &
fi
pids+=( $! )

make gen &
pids+=( $! )

for pid in ${pids[*]}; do
  wait $pid
done

unset pids

services="users versus"

echo "[STARTING] - Building Go services"

for service in $services; do
  if [ -f "build/${service}" ]; then
    rm build/$service
  fi

  (go build -o build/$service "backend/${service}/main.go") & 
  pids+=( $! )
done

for pid in ${pids[*]}; do
  wait $pid
done

echo "[FINISHED] - Building Go services"

end=`date +%s`

runtime=$((end-start))
echo "Duration: ${runtime}s"
