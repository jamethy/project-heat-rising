data aws_s3_bucket_object daily_data_lambda_zip {
  bucket = local.infra_bucket
  key    = "lambdas/daily-data-lambda.zip"
}

data aws_s3_bucket_object thermostat_lambda_zip {
  bucket = local.infra_bucket
  key    = "lambdas/thermostat-lambda.zip"
}

data aws_s3_bucket_object weather_lambda_zip {
  bucket = local.infra_bucket
  key    = "lambdas/weather-lambda.zip"
}

data aws_s3_bucket_object upstairs_lambda_zip {
  bucket = local.infra_bucket
  key    = "lambdas/upstairs-lambda.zip"
}
