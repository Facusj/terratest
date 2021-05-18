variable "region" {
  default     = "us-east-1"
  description = "AWS region"
}

variable "ami" {
  default     = "ami-09e67e426f25ce0d7"
  description = "EC2 Image ID"
}

variable "instance_type" {
  default     = "t2.micro"
  description = "EC2 instance Size"
}

variable "name" {
  default     = "Flugel"
  description = "Tag Name"
}

variable "owner" {
  default     = "InfraTeam"
  description = "Tag Owner"
}


variable "bucket_name" {
  default     = "flugel"
  description = "S3 bucket name"
}