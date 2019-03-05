variable "region" {
  description = "AWS region to run the standalone kubernaut vm"
  default     = "us-east-1"
}

variable "disk_size" {
  default = 20
}

variable "iam_role" {
  description = "IAM role configuration for the machine"
  default = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Effect": "Allow",
      "Principal": {
        "Service": "ec2.amazonaws.com"
      }
    }
  ]
}
EOF
}

variable "iam_policy" {
  description = "IAM policy configuration for the machine"
  default = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Action": "*",
            "Effect": "Allow",
            "Resource": "*"
        }
    ]
}
EOF
}

variable "machine_type" {
  description = "The size of the EC2 virtual machine"
  default     = "m3.medium"
}

variable "tags" {
  description = "Tags added to the Kubernaut cluster infrastructure"
  type        = "map"
  default     = {}
}