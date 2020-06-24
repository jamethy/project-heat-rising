resource aws_lambda_function daily_data {

  function_name = "${local.prefix}-daily-data"

  s3_bucket         = data.aws_s3_bucket_object.daily_data_lambda_zip.bucket
  s3_key            = data.aws_s3_bucket_object.daily_data_lambda_zip.key
  s3_object_version = data.aws_s3_bucket_object.daily_data_lambda_zip.version_id

  runtime = "go1.x"
  handler = "daily-data-lambda"

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

resource aws_cloudwatch_log_group daily_data {
  name              = "/aws/lambda/${aws_lambda_function.daily_data.function_name}"
  tags              = local.common_tags
  retention_in_days = local.log_retention
}

resource aws_cloudwatch_event_rule daily_data {
  name                = "${local.prefix}-daily-data"
  schedule_expression = "cron(0 10 * * ? *)" // 10 am UTC every day
}

resource aws_cloudwatch_event_target daily_data_target {
  rule = aws_cloudwatch_event_rule.daily_data.name
  arn  = aws_lambda_function.daily_data.arn
}
