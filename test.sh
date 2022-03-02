#!/usr/bin/env bash
start=`date +%s`

set -eu -o pipefail

pids=()

cur_dir=$PWD

echo "Building DB"

cd infra/scylla/dev && docker-compose down --volumes && docker-compose build && docker-compose up -d &
pids+=( $! )

for pid in ${pids[*]}; do
  wait $pid
done

unset pids

ip=$(docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' scylla-node1)
echo $ip

echo "Waiting for scylla to start, retrying every 30 seconds"
while ! cqlsh $ip -u cassandra -p cassandra -e "select rpc_address from system.local" ; do
    sleep 30
done
echo "Connected successfully"

cd db/migrations

for i in *.up.cql; do
  echo "Running migration $i"
  cqlsh $ip -u cassandra -p cassandra -f "$i" & 
  pids+=( $! )
done

for pid in ${pids[*]}; do
  wait $pid
done

cd $cur_dir/backend/users/pkg/scylla && go test -v

end=`date +%s`

runtime=$((end-start))
echo "Duration: ${runtime}s"
