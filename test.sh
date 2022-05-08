#!/usr/bin/env bash
start=`date +%s`

set -eu -o pipefail

pids=()

cur_dir=$PWD

echo "Building DB"

cd infra/scylla/dev && docker-compose down --volumes && docker-compose build --pull && docker-compose up -d &
pids+=( $! )

for pid in ${pids[*]}; do
  wait $pid
done

unset pids

echo "Waiting for scylla to start, retrying every 30 seconds"

while [ "$(docker exec scylla-node1 nodetool status | grep -c UN)" -lt 3 ] ; do
    echo "Retrying in 30 seconds"
    sleep 30
done

echo "Connected successfully"

bash db_migrations.sh & wait $!

cd $cur_dir/backend/users/pkg/scylla && go test -race -v
cd $cur_dir/backend/versus/pkg/api && go test -race -v

end=`date +%s`

runtime=$((end-start))
echo "Duration: ${runtime}s"
