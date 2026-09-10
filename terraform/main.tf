data "http" "myip" {
  url = "https://checkip.amazonaws.com"
}

resource "aws_security_group" "goflow_rds" {
  name        = "goflow-rds-sg"
  description = "Security group for GoFlow RDS PostgreSQL"

  ingress {
    description = "PostgreSQL from current public IP"
    protocol    = "tcp"
    from_port   = 5432
    to_port     = 5432
    cidr_blocks = ["${chomp(data.http.myip.response_body)}/32"]
  }

  egress {
    description = "Allow outbound traffic"
    protocol    = "-1"
    from_port   = 0
    to_port     = 0
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name      = "goflow-rds-sg"
    Project   = "GoFlow"
    ManagedBy = "Terraform"
  }
}


resource "aws_db_instance" "goflow" {
  identifier = "goflow-postgres"

  engine         = "postgres"
  engine_version = "17"

  instance_class        = "db.t4g.micro"
  allocated_storage     = 20
  max_allocated_storage = 50
  storage_type          = "gp3"
  storage_encrypted     = true

  db_name  = var.db_name
  username = var.db_username
  password = var.db_password
  port     = 5432

  vpc_security_group_ids = [aws_security_group.goflow_rds.id]

  publicly_accessible = true

  backup_retention_period = 1
  deletion_protection     = false
  skip_final_snapshot     = true

  auto_minor_version_upgrade = true

  tags = {
    Name        = "goflow-postgres"
    Project     = "GoFlow"
    Environment = "development"
    ManagedBy   = "Terraform"
  }
}