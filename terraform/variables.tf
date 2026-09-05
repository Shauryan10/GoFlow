variable "aws_region" {
  description = "AWS region"
  type        = string
  default     = "ap-south-1"
}

variable "db_name" {
  description = "GoFlow PostgreSQL database name"
  type        = string
  default     = "goflow"
}

variable "db_username" {
  description = "RDS PostgreSQL username"
  type        = string
  default     = "goflow_admin"
}

variable "db_password" {
  description = "RDS PostgreSQL password"
  type        = string
  sensitive   = true //important for password protection

}