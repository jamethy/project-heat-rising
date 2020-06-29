// must create "pi" user with aws credentials and SQS access
resource aws_lambda_function upstairs {

  function_name = "${local.prefix}-upstairs"

  s3_bucket         = data.aws_s3_bucket_object.upstairs_lambda_zip.bucket
  s3_key            = data.aws_s3_bucket_object.upstairs_lambda_zip.key
  s3_object_version = data.aws_s3_bucket_object.upstairs_lambda_zip.version_id

  runtime = "go1.x"
  handler = "upstairs-lambda"

  timeout = "30"
  // seconds
  memory_size = "128"
  // mb

  role = aws_iam_role.general_lambda.arn

  tags = local.common_tags

  environment {
    variables = {
      DATABASE_URL      = var.database_url
      DATABASE_USERNAME = var.database_username
      DATABASE_PASSWORD = var.database_password
    }
  }
}

resource aws_cloudwatch_log_group upstairs {
  name              = "/aws/lambda/${aws_lambda_function.upstairs.function_name}"
  tags              = local.common_tags
  retention_in_days = local.log_retention
}

resource aws_sqs_queue upstairs_input_dlq {
  name = "${local.prefix}-upstairs-input-dlq"
  tags = local.common_tags
}

// copy url to sense_hat_collector/config.json sqsUrl
resource aws_sqs_queue upstairs_input {
  name = "${local.prefix}-upstairs-input"
  redrive_policy = jsonencode({
    deadLetterTargetArn = aws_sqs_queue.upstairs_input_dlq.arn
    maxReceiveCount     = 3
  })
  tags = local.common_tags
}

resource aws_iam_policy queue_read_write_policy {
  name        = "${local.prefix}-lambda-sqs"
  path        = "/"
  description = "IAM policy for sending and reading messages from SQS"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "sqs:SendMessage",
        "sqs:ReceiveMessage",
        "sqs:DeleteMessage",
        "sqs:GetQueueAttributes"
        ],
      "Resource": "${aws_sqs_queue.upstairs_input.arn}"
    }
  ]
}
EOF
}

resource aws_iam_role_policy_attachment lambda_sqs {
  policy_arn = aws_iam_policy.queue_read_write_policy.arn
  role       = aws_iam_role.general_lambda.name
}

resource aws_lambda_event_source_mapping upstairs_input {
  event_source_arn = aws_sqs_queue.upstairs_input.arn
  function_name    = aws_lambda_function.upstairs.function_name
}
