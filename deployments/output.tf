
output "rds_endpoint" {
  value = aws_db_instance.postgres_rds.address
  description = "RDS Endpoint"
}

output "rds_username" {
  value = aws_db_instance.postgres_rds.username
    description = "RDS Username"
}

output "rds_password" {
  value = aws_db_instance.postgres_rds.password
    description = "RDS Password"
}

output "rds_db_name" {
  value = var.db_name
    description = "RDS Database Name"
}