
resource aws_iam_role general_lambda {
  name_prefix        = local.prefix
  tags               = local.common_tags
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource aws_iam_policy lambda_logging {
  name_prefix = "${local.prefix}-lambda-logging"
  path        = "/"
  description = "IAM policy for logging from a lambda"
  policy      = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": "logs:CreateLogGroup",
            "Resource": "arn:aws:logs:${local.region}:*:*"
        },
        {
            "Effect": "Allow",
            "Action": [
                "logs:CreateLogStream",
                "logs:PutLogEvents"
            ],
            "Resource": [
                "${aws_cloudwatch_log_group.daily_data.arn}",
                "${aws_cloudwatch_log_group.weather.arn}",
                "${aws_cloudwatch_log_group.thermostat.arn}"
            ]
        }
    ]
}
EOF
}

resource aws_iam_role_policy_attachment lambda_logging {
  policy_arn = aws_iam_policy.lambda_logging.arn
  role       = aws_iam_role.general_lambda.name
}


