start=`date +%s`

set -eu -o pipefail

pids=()

services="users"

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
