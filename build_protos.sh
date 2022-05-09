echo "the following files will be built:"
buf ls-files

buf lint

if [ $# -eq 1 ]; then
  if [ $1 = "ts" ]; then
    echo "-- installing ts-proto --"
    yarn

    echo "[STARTING] - Generating go and ts types"
    buf generate -o protos --template buf.ts.gen.yaml
    cp protos/users/proto/users.ts frontend/flashy/src/types
    cp protos/users/proto/users.proto frontend/flashy/android/app/src/main/proto
    cp protos/versus/proto/versus.proto frontend/flashy/android/app/src/main/proto
    echo "[FINISHED] - Generating go and ts types"
  else
    echo "Typescript types not generated."
  fi
else
  echo "[STARTING] - Generating go types"
  buf generate -o protos --template buf.gen.yaml
  echo "[FINISHED] - Generating go types"
  echo "Typescript types not generated."
fi



