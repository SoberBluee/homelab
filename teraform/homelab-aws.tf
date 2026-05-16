terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "6.45.0"
    }
  }
}

provider "aws" {
    # Configuration options
    region = "eu-west-2"

    endpoints { 
        s3       = "http://aws.local"
        ec2      = "http://aws.local"
        dynamodb = "http://aws.local"
    }
}

resource "aws_iam_group" "read_only_s3" { 
  name = "ready_only_s3"
  path = "/read_only/s3"
}

resource "aws_"

resource "aws_s3_bucket" { 

    bucket = "my_test_s3"
    region = "eu-west-2"

    tags = { 
      Name = "Heartbeat bucket"
      Environment = "Dev"
    }
}