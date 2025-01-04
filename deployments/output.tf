
output "rds_endpoint" {
  value = aws_db_instance.postgres_rds.address
  description = "RDS Endpoint"
}
