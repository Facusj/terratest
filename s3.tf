resource "aws_s3_bucket" "Flugel" {
  bucket = var.bucket_name
  tags = {
    "Name"  = var.name
    "Owner" = var.owner
  }
}