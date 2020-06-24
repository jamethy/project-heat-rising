// NOTE, before running you need the secrets file.
// get file: AWS_PROFILE=personal aws s3 cp s3://project-rising-heat-infra/terraform/secrets.tfvars .
// upload file: AWS_PROFILE=personal aws s3 cp secrets.tfvars s3://project-rising-heat-infra/terraform/secrets.tfvars
// Run with: terraform apply -var-file=secrets.tfvars
// =====================================================================================================================

provider aws {
  region  = "us-east-1"
  profile = "personal"
  version = "2.67.0"
}

terraform {
  backend s3 {
    bucket  = "project-rising-heat-infra"
    key     = "terraform/state"
    region  = "us-east-1"
    profile = "personal"
  }
}

locals {
  prefix = "prh"
  common_tags = {
    "project" : "Project Rising Heat"
  }
  infra_bucket  = "project-rising-heat-infra"
  region        = "us-east-1"
  log_retention = 7 // days
}

