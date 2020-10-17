terraform {
  required_providers {
    tss = {
      source  = "terraform.thycotic.com/thycotic/tss"
      version = "~> 1.0"
    }
  }
}

variable "tss_username" {
  type = string
}

variable "tss_password" {
  type = string
}

variable "tss_server_url" {
  type = string
}

variable "tss_secret_id" {
  type = string
}

provider "tss" {
  username   = var.tss_username
  password   = var.tss_password
  server_url = var.tss_server_url
}

data "tss_secret" "my_username" {
  id    = var.tss_secret_id
  field = "username"
}

data "tss_secret" "my_password" {
  id    = var.tss_secret_id
  field = "password"
}

output "username" {
  value     = data.tss_secret.my_username.value
}

output "password" {
  value     = data.tss_secret.my_password.value
}
