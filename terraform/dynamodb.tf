resource "aws_dynamodb_table" "basic-dynamodb-table" {
  name           = "ParameterAPI"
  billing_mode   = "PROVISIONED"
  read_capacity  = 20
  write_capacity = 20
  hash_key       = "Key"
//  range_key      = "Key"

  attribute {
    name = "Key"
    type = "S"
  }

//  attribute {
//    name = "Key"
//    type = "S"
//  }

  ttl {
    attribute_name = "TimeToExist"
    enabled        = false
  }

  tags = {
    Name        = "dynamodb-table-ParameterAPI"
    Environment = "development"
  }
}