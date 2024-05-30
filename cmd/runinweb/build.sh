echo $1
cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./cmd/runinweb/
env GOOS=js GOARCH=wasm go build -o ./cmd/runinweb/yourGame.wasm $1
go run ./cmd/runinweb