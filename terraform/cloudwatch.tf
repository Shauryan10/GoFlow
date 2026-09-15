data "aws_sns_topic" "goflow_alerts" {
  name = "GoFlowMonitoringAlerts"
}

resource "aws_cloudwatch_metric_alarm" "goflow_rds_cpu" {
  alarm_name        = "GoFlow-RDS-High-CPU"
  alarm_description = "Triggers when GoFlow RDS CPU utilization exceeds 80% for 5 minutes"

  namespace           = "AWS/RDS"
  metric_name         = "CPUUtilization"
  comparison_operator = "GreaterThanThreshold"

  evaluation_periods = 1
  period             = 300
  statistic          = "Average"
  threshold          = 80

  dimensions = {
    DBInstanceIdentifier = aws_db_instance.goflow.identifier
  }

  alarm_actions = [
    data.aws_sns_topic.goflow_alerts.arn
  ]

  treat_missing_data = "notBreaching"

  tags = {
    Project   = "GoFlow"
    ManagedBy = "Terraform"
  }
}

resource "aws_cloudwatch_metric_alarm" "goflow_rds_storage" {
  alarm_name        = "GoFlow-RDS-Low-Storage"
  alarm_description = "Triggers when GoFlow RDS free storage falls below 2 GB"

  namespace           = "AWS/RDS"
  metric_name         = "FreeStorageSpace"
  comparison_operator = "LessThanThreshold"

  evaluation_periods = 1
  period             = 300
  statistic          = "Average"
  threshold          = 2147483648

  dimensions = {
    DBInstanceIdentifier = aws_db_instance.goflow.identifier
  }

  alarm_actions = [
    data.aws_sns_topic.goflow_alerts.arn
  ]

  treat_missing_data = "notBreaching"

  tags = {
    Project   = "GoFlow"
    ManagedBy = "Terraform"
  }
}


resource "aws_cloudwatch_metric_alarm" "goflow_rds_connections" {
  alarm_name        = "GoFlow-RDS-High-Connections"
  alarm_description = "Triggers when GoFlow RDS database connections exceed 80"

  namespace           = "AWS/RDS"
  metric_name         = "DatabaseConnections"
  comparison_operator = "GreaterThanThreshold"

  evaluation_periods = 1
  period             = 300
  statistic          = "Average"
  threshold          = 80

  dimensions = {
    DBInstanceIdentifier = aws_db_instance.goflow.identifier
  }

  alarm_actions = [
    data.aws_sns_topic.goflow_alerts.arn
  ]

  treat_missing_data = "notBreaching"

  tags = {
    Project   = "GoFlow"
    ManagedBy = "Terraform"
  }
}


resource "aws_cloudwatch_metric_alarm" "goflow_rds_memory" {
  alarm_name        = "GoFlow-RDS-Low-Memory"
  alarm_description = "Triggers when GoFlow RDS freeable memory falls below 256 MB"

  namespace           = "AWS/RDS"
  metric_name         = "FreeableMemory"
  comparison_operator = "LessThanThreshold"

  evaluation_periods = 1
  period             = 300
  statistic          = "Average"
  threshold          = 268435456

  dimensions = {
    DBInstanceIdentifier = aws_db_instance.goflow.identifier
  }

  alarm_actions = [
    data.aws_sns_topic.goflow_alerts.arn
  ]

  treat_missing_data = "notBreaching"

  tags = {
    Project   = "GoFlow"
    ManagedBy = "Terraform"
  }
}

