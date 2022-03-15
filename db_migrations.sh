set -eu -o pipefail

cd db/migrations

for i in *.up.cql; do
  echo "Running migration $i"
  cqlsh -u cassandra -p cassandra -f "$i" & 
  wait $! 
done
