terraform {
  required_providers {
    scaffolding = {
      version = "0.1"
      source  = "kocasarac.com/edu/scaffolding"
    }
  }
}

provider "scaffolding" {
  profile = "asd"
}

output "ec2" {
  value = scaffolding_resource.ec2
}