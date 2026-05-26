provider "aws" {
  region                      = "us-east-1"
  access_key                  = "test"
  secret_key                  = "test"
  
  # Crucial for LocalStack S3 routing
  s3_use_path_style           = true
  
  # Skip validations that require a real AWS connection
  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true

  # Point the services you are using to LocalStack
  endpoints {
    s3         = "http://localhost:4566"
    dynamodb   = "http://localhost:4566"
    lambda     = "http://localhost:4566"
    sts        = "http://localhost:4566"
    iam        = "http://localhost:4566"
    ec2        = "http://localhost:4566"
  }
}

resource "aws_s3_bucket" "my_local_bucket" {
  bucket = "localstack-terraform-bucket"
}

variable "bucket_name" { 
    description = "A test bucket"
    type        = string
    default     = "bucket"
}

variable "ec2name" { 
    description = "a test ec2 instance"
    type        = string
    default     = "ec2-default-name"
}

variable "ec2type" { 
    description = "the ec2 name"
    type        =  string
    default     = "t2.micro"
}


# data "aws_ami" "ubuntu" {
#   most_recent = true

#   filter {
#     name = "name"
#     values = ["ubuntu/images/hvm-ssd-gp3/ubuntu-noble-24.04-amd64-server-*"]
#   }

#   owners = ["099720109477"] # Canonical
# }


resource "aws_instance" "app_server" {
  ami             = "ami-1234567890abcdef0" # Hardcoded dummy AMI for LocalStack
  instance_type   = var.ec2type

  tags = {
    Name = var.ec2name
  }
}

module "aws_vpn" { 
  
}


