set dotenv-load := true
set export := true
set shell := ["bash", "-euc"]

default:
    just --list

clean:
    @echo "🧹 Cleaning generated files..."
    @rm -rf etc internal

# @just gen-api [src_api] - Generates API service. Defaults to user.api
# The output directory will be the parent of the api file's directory.
# Example: `just gen-api desc/user.api` will output to `.`
# Example: `just gen-api apis/v1/user.api` will output to `apis`
gen-api src_api: clean
    @echo "🚀 Generating API service from {{src_api}}..."
    @echo "Output directory will be: $(dirname $(dirname {{src_api}}))"
    @goctl api go -api {{src_api}} -dir $(dirname $(dirname {{src_api}})) --style go_zero
    @echo "✅ API generated successfully."

# @just gen-rpc [src_proto] - Generates RPC service. Defaults to user.proto
# The output directory will be the parent of the proto file's directory.
gen-rpc src_proto: clean
    @echo "🚀 Generating RPC service from {{src_proto}}..."
    @echo "Output directory will be: $(dirname $(dirname {{src_proto}}))"
    @SERVICE_NAME=$(basename {{src_proto}} | sed 's/\.proto$//')
    @goctl rpc protoc {{src_proto}} --go_out=$(dirname $(dirname {{src_proto}})) --go-grpc_out=$(dirname $(dirname {{src_proto}})) --zrpc_out=$(dirname $(dirname {{src_proto}})) --style go_zero
    @find $(dirname $(dirname {{src_proto}}))/pb -type f -name "*.pb.go" -exec sed -i 's/,omitempty//g' {} +
    @echo "✅ RPC generated successfully."
# @just gen-model <db_suffix> <table_name> - Generates a model into a generic directory.
# Example: `just gen-model usercenter user`
gen-model db_suffix table_name:
    @echo "🚀 Generating model for table '{{table_name}}' from database {{env('DB_NAME_PREFIX')}}{{db_suffix}}..."

    @goctl model mysql datasource \
        -url="{{env('DB_USER')}}:{{env('DB_PASS')}}@tcp({{env('DB_HOST')}}:{{env('DB_PORT')}})/{{env('DB_NAME_PREFIX')}}{{db_suffix}}" \
        -table="{{table_name}}" \
        -dir="{{env('MODEL_GEN_DIR')}}" \
        -cache=true \
        --style=goZero

    @echo "✅ Model generated in '{{env('MODEL_GEN_DIR')}}'. You can now move the files to the desired service directory."
