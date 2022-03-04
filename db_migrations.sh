set -eu -o pipefail

pids=()

cd db/migrations

for i in *.up.cql; do
  echo "Running migration $i"
  cqlsh -u cassandra -p cassandra -f "$i" & 
  pids+=( $! )
done

for pid in ${pids[*]}; do
  wait $pid
done
