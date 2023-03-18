# Usage

## Start Development Environment

```shell
# start databases
docker-compose -f ./db-docker-compose.yaml up -d
# apis
# run `go install github.com/cosmtrek/air@latest` if air not installed
air
# ui
cd ./web && tyarn dev

# shut down databases
docker-compose -f ./db-docker-compose.yaml down
```

## Run Executable In Production
