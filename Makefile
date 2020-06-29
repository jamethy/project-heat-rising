# todo learn how to write make properly
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

build-upstairs-lambda:
	GOARCH=amd64 GOOS=linux go build -v -o bin/upstairs-lambda -ldflags "-w -s" ./cmd/upstairs_lambda/main.go
	zip -j -qq ./bin/upstairs-lambda.zip ./bin/upstairs-lambda

build-all: build-daily-data-lambda build-thermostat-lambda build-weather-lambda build-upstairs-lambda

# todo use goreleaser
deploy-daily-data-lambda:
	AWS_PROFILE=personal aws s3 cp bin/thermostat-lambda.zip s3://project-rising-heat-infra/lambdas/daily-data-lambda.zip
	AWS_PROFILE=personal aws lambda update-function-code --function-name prh-daily-data --s3-bucket project-rising-heat-infra --s3-key lambdas/daily-data-lambda.zip

deploy-thermostat-lambda:
	AWS_PROFILE=personal aws s3 cp bin/thermostat-lambda.zip s3://project-rising-heat-infra/lambdas/thermostat-lambda.zip
	AWS_PROFILE=personal aws lambda update-function-code --function-name prh-thermostat --s3-bucket project-rising-heat-infra --s3-key lambdas/thermostat-lambda.zip

deploy-weather-lambda:
	AWS_PROFILE=personal aws s3 cp bin/weather-lambda.zip s3://project-rising-heat-infra/lambdas/weather-lambda.zip
	AWS_PROFILE=personal aws lambda update-function-code --function-name prh-weather --s3-bucket project-rising-heat-infra --s3-key lambdas/weather-lambda.zip

deploy-upstairs-lambda:
	AWS_PROFILE=personal aws s3 cp bin/upstairs-lambda.zip s3://project-rising-heat-infra/lambdas/upstairs-lambda.zip
	AWS_PROFILE=personal aws lambda update-function-code --function-name prh-upstairs --s3-bucket project-rising-heat-infra --s3-key lambdas/upstairs-lambda.zip

deploy-all: deploy-daily-data-lambda deploy-thermostat-lambda deploy-weather-lambda deploy-upstairs-lambda

generate:
	rm `grep -l SQLBoiler internal/db/*` || true
	sqlboiler -c internal/db/sqlboiler.toml --no-tests -o internal/db -p db psql


test:
	go test -v ./...
clean:
	go clean
	rm -r bin/*
