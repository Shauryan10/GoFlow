data "aws_caller_identity" "current" {}

# ---------------------------------------------------------
# IAM User - GoFlow Monitoring
# ---------------------------------------------------------

resource "aws_iam_user" "goflow_monitoring_user" {
  name = "GoFlowMonitoringUser"

  tags = {
    Project   = "GoFlow"
    Purpose   = "CloudWatch Monitoring"
    ManagedBy = "Terraform"
  }
}

# ---------------------------------------------------------
# CloudWatch / CloudWatch Logs - Read Only Policy
# ---------------------------------------------------------

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

# ---------------------------------------------------------
# IAM Role - GoFlow Monitoring
# ---------------------------------------------------------

resource "aws_iam_role" "goflow_monitoring" {
  name = "GoFlowMonitoringRole"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"

    Statement = [
      {
        Effect = "Allow"

        Principal = {
          AWS = aws_iam_user.goflow_monitoring_user.arn
        }

        Action = "sts:AssumeRole"

        Condition = {
          Bool = {
            "aws:MultiFactorAuthPresent" = "true"
          }
        }
      }
    ]
  })

  tags = {
    Project   = "GoFlow"
    Purpose   = "CloudWatch Monitoring"
    ManagedBy = "Terraform"
  }
}

# ---------------------------------------------------------
# Attach CloudWatch Read-Only Policy to Monitoring Role
# ---------------------------------------------------------

resource "aws_iam_role_policy_attachment" "goflow_monitoring" {
  role       = aws_iam_role.goflow_monitoring.name
  policy_arn = aws_iam_policy.goflow_cloudwatch_read.arn
}

# ---------------------------------------------------------
# Allow Monitoring User to Assume ONLY this Role
# ---------------------------------------------------------

resource "aws_iam_policy" "goflow_assume_monitoring_role" {
  name        = "GoFlowAssumeMonitoringRole"
  description = "Allows GoFlow monitoring user to assume only the GoFlow monitoring role"

  policy = jsonencode({
    Version = "2012-10-17"

    Statement = [
      {
        Effect = "Allow"

        Action = "sts:AssumeRole"

        Resource = aws_iam_role.goflow_monitoring.arn
      }
    ]
  })
}

resource "aws_iam_user_policy_attachment" "goflow_monitoring_user" {
  user       = aws_iam_user.goflow_monitoring_user.name
  policy_arn = aws_iam_policy.goflow_assume_monitoring_role.arn
}