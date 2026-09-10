output "rds_endpoint" {
  description = "RDS PostgreSQL endpoint"
  value       = aws_db_instance.goflow.address
}

output "rds_port" {
  description = "RDS PostgreSQL port"
  value       = aws_db_instance.goflow.port
}

output "rds_database_name" {
  description = "GoFlow database name"
  value       = aws_db_instance.goflow.db_name
}

output "goflow_monitoring_role_arn" {
  description = "ARN of the GoFlow CloudWatch monitoring IAM role"
  value       = aws_iam_role.goflow_monitoring.arn
}

output "goflow_monitoring_policy_arn" {
  description = "ARN of the GoFlow CloudWatch read-only policy"
  value       = aws_iam_policy.goflow_cloudwatch_read.arn
}