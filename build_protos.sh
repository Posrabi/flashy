echo "the following files will be built:"
buf ls-files

buf lint

echo "-- installing ts-proto --"
yarn

echo "[STARTING] - Generating go and ts types"
buf generate -o protos --template buf.gen.yaml
buf generate -o protos --template buf_java.gen.yaml
cp protos/users/proto/users.ts frontend/Flashy/src/types/flashy.ts
cp protos/com/flashy/Users.java frontend/Flashy/android/app/src/main/java/com/flashy/Users.java
echo "[FINISHED] - Generating go and ts types"
