resource aws_lambda_function thermostat {

  function_name = "${local.prefix}-thermostat"

  s3_bucket         = data.aws_s3_bucket_object.thermostat_lambda_zip.bucket
  s3_key            = data.aws_s3_bucket_object.thermostat_lambda_zip.key
  s3_object_version = data.aws_s3_bucket_object.thermostat_lambda_zip.version_id

  runtime = "go1.x"
  handler = "thermostat-lambda"

  timeout     = "30"  // seconds
  memory_size = "128" // mb

  role = aws_iam_role.general_lambda.arn

  tags = local.common_tags

  environment {
    variables = {
      CARRIER_USERNAME  = var.carrier_username
      CARRIER_PASSWORD  = var.carrier_password
      DATABASE_URL      = var.database_url
      DATABASE_USERNAME = var.database_username
      DATABASE_PASSWORD = var.database_password
    }
  }
}

resource aws_cloudwatch_log_group thermostat {
  name              = "/aws/lambda/${aws_lambda_function.thermostat.function_name}"
  tags              = local.common_tags
  retention_in_days = local.log_retention
}

resource aws_cloudwatch_event_rule thermostat {
  name                = "${local.prefix}-thermostat"
  schedule_expression = "cron(* * * * ? *)" // Every minute
}

resource aws_cloudwatch_event_target thermostat_target {
  rule = aws_cloudwatch_event_rule.thermostat.name
  arn  = aws_lambda_function.thermostat.arn
}

resource aws_lambda_permission thermostat_cloudwatch {
  statement_id  = "AllowExecutionFromCloudWatch"
  action        = "lambda:InvokeFunction"
  function_name = "${aws_lambda_function.thermostat.function_name}"
  principal     = "events.amazonaws.com"
  source_arn    = aws_cloudwatch_event_rule.thermostat.arn
}
