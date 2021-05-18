output "tags" {
  description = "List of tags"
  value       = aws_instance.Flugel.tags
}

output "instance_id" {
  description = "EC2 instance ID"
  value       = aws_instance.Flugel.id
}