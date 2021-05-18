terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
    }
    random = {
      source = "hashicorp/random"
    }
  }
  // required_version = "~> 0.15"
  backend "remote" {
    hostname     = "app.terraform.io"
    organization = "Test-terraform-Arg"

    workspaces {
      name = "gh-actions"
    }
  }
}

provider "aws" {
  profile = "default"
  region  = var.region
}
resource "aws_instance" "Flugel" {
  ami           = var.ami
  instance_type = var.instance_type
  tags = {
    "Name"  = var.name
    "Owner" = var.owner
  }
}