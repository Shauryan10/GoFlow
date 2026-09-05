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