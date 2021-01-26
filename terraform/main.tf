# Cloud
//provider "aws" {
//  version = "~> 3.0"
//  region  = "us-east-1"
//}

provider "archive" {}

# local
 terraform {
   backend "local" {}
 }

 provider "aws" {
   access_key                  = "mock_access_key"
   region                      = "us-east-1"
   s3_force_path_style         = true
   secret_key                  = "mock_secret_key"
   skip_credentials_validation = true
   skip_metadata_api_check     = true
   skip_requesting_account_id  = true

   endpoints {
     dynamodb    = "http://0.0.0.0:4566"
   }
 }