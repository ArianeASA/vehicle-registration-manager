
variable "aws_region" {
  description = "AWS region for all resources."

  type    = string
  default = "us-east-2"
}

variable "account_id" {
  type    = string
}

variable "access_key" {
  type    = string
}

variable "secret_key" {
  type    = string
}

variable "db_username" {
  type    = string
  description = "The username for the database."
}

variable "db_password" {
  type    = string
  description = "The password for the database."
}

variable "db_name" {
  type    = string
  description = "The name of the database."
}

variable "db_driver" {
  type    = string
  description = "The driver for the database."
}

variable "db_schema" {
  type    = string
  description = "The schema for the database."
}

variable "scope" {
  type    = string
  description = "The scope for the database."
}