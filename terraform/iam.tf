data "aws_caller_identity" "current" {}

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

resource "aws_iam_role" "goflow_monitoring" {
  name = "GoFlowMonitoringRole"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"

    Statement = [
      {
        Effect = "Allow"
        Principal = {
          AWS = "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"
        }
        Action = "sts:AssumeRole"
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "goflow_monitoring" {
  role       = aws_iam_role.goflow_monitoring.name
  policy_arn = aws_iam_policy.goflow_cloudwatch_read.arn
}