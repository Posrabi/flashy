echo "the following files will be built:"
buf ls-files

buf lint

echo "[STARTING] - Generating go types"
buf generate -o protos --template buf.gen.yaml
echo "[FINISHED] - Generating go types"
