resource aws_lambda_function weather {

  function_name = "${local.prefix}-weather"

  s3_bucket         = data.aws_s3_bucket_object.weather_lambda_zip.bucket
  s3_key            = data.aws_s3_bucket_object.weather_lambda_zip.key
  s3_object_version = data.aws_s3_bucket_object.weather_lambda_zip.version_id

  runtime = "go1.x"
  handler = "weather-lambda"

  timeout     = "30"  // seconds
  memory_size = "128" // mb

  role = aws_iam_role.general_lambda.arn

  tags = local.common_tags

  environment {
    variables = {
      OPEN_WEATHER_BASE_URL = var.open_weather_base_url
      OPEN_WEATHER_API_KEY  = var.open_weather_api_key
      OPEN_WEATHER_LAT      = var.open_weather_lat
      OPEN_WEATHER_LON      = var.open_weather_lon
      DATABASE_URL          = var.database_url
      DATABASE_USERNAME     = var.database_username
      DATABASE_PASSWORD     = var.database_password
    }
  }
}

resource aws_cloudwatch_log_group weather {
  name              = "/aws/lambda/${aws_lambda_function.weather.function_name}"
  tags              = local.common_tags
  retention_in_days = local.log_retention
}

resource aws_cloudwatch_event_rule weather {
  name                = "${local.prefix}-weather"
  schedule_expression = "cron(*/2 * * * ? *)" // Every minute
}

resource aws_cloudwatch_event_target weather_target {
  rule = aws_cloudwatch_event_rule.weather.name
  arn  = aws_lambda_function.weather.arn
}

resource aws_lambda_permission weather_cloudwatch {
  statement_id  = "AllowExecutionFromCloudWatch"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.weather.function_name
  principal     = "events.amazonaws.com"
  source_arn    = aws_cloudwatch_event_rule.weather.arn
}
