terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.27"
    }
  }

  backend "remote" {
    organization = "nickforero"

    workspaces {
      name = "flugel_actions_pipeline"
    }
  }

  required_version = ">= 0.14.9"
}

provider "aws" {
  profile = "default"
  region  = "us-west-2"
}

resource "aws_s3_bucket" "flugel_test_bucket" {
  bucket = "flugel-tf-test-bucket"

  tags = {
    Name  = "Flugel"
    Owner = "InfraTeam"
  }
}

resource "aws_instance" "flugel_test_instance" {
  ami           = "ami-830c94e3"
  instance_type = "t2.micro"

  tags = {
    Name  = "Flugel"
    Owner = "InfraTeam"
  }
}
