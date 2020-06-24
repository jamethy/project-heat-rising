all: test build-all

build-daily-data-lambda:
	GOARCH=amd64 GOOS=linux go build -v -o bin/daily-data-lambda -ldflags "-w -s" ./cmd/daily_data_lambda/main.go
	zip -j -qq ./bin/daily-data-lambda.zip ./bin/daily-data-lambda

build-thermostat-lambda:
	GOARCH=amd64 GOOS=linux go build -v -o bin/thermostat-lambda -ldflags "-w -s" ./cmd/thermostat_lambda/main.go
	zip -j -qq ./bin/thermostat-lambda.zip ./bin/thermostat-lambda

build-weather-lambda:
	GOARCH=amd64 GOOS=linux go build -v -o bin/weather-lambda -ldflags "-w -s" ./cmd/weather_lambda/main.go
	zip -j -qq ./bin/weather-lambda.zip ./bin/weather-lambda

build-all: build-daily-data-lambda build-thermostat-lambda build-weather-lambda

generate:
	rm `grep -l SQLBoiler internal/db/*` || true
	sqlboiler -c internal/db/sqlboiler.toml --no-tests -o internal/db -p db psql


test:
	go test -v ./...
clean:
	go clean
	rm -r bin/*
