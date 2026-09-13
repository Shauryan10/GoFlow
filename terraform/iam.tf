resource "aws_iam_policy" "goflow_cloudwatch_read" {
  name        = "GoFlowCloudWatchReadOnly"
  description = "Read-only access to CloudWatch metrics and logs for GoFlow monitoring"

  policy = jsonencode({
    Version = "2012-10-17"

    Statement = [
      {
        Effect = "Allow"

        Action = [
          "cloudwatch:GetMetricData",
          "cloudwatch:GetMetricStatistics",
          "cloudwatch:ListMetrics",
          "logs:DescribeLogGroups",
          "logs:DescribeLogStreams",
          "logs:GetLogEvents",
          "logs:FilterLogEvents"
        ]

        Resource = "*"
      }
    ]
  })
}