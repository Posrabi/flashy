echo "the following files will be built:"
buf ls-files

buf lint

echo "-- installing ts-proto --"
yarn

echo "[STARTING] - Generating go and ts types"
buf generate -o protos --template buf.gen.yaml
cp protos/users/proto/users.ts frontend/flashy/src/types
cp protos/users/proto/users.proto frontend/flashy/android/app/src/main/proto
echo "[FINISHED] - Generating go and ts types"
