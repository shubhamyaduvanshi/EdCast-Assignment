provider "aws" {
  region = "us-east-1"  # Adjust the region as per your AWS setup
}

resource "aws_api_gateway_rest_api" "example_api" {
  name        = "my-api"
  description = "API Gateway for my application"
}

resource "aws_api_gateway_resource" "add_resource" {
  rest_api_id = aws_api_gateway_rest_api.example_api.id
  parent_id   = aws_api_gateway_rest_api.example_api.root_resource_id
  path_part   = "add"
}

resource "aws_api_gateway_method" "post_method" {
  rest_api_id   = aws_api_gateway_rest_api.example_api.id
  resource_id   = aws_api_gateway_resource.add_resource.id
  http_method   = "POST"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "app_integration" {
  rest_api_id             = aws_api_gateway_rest_api.example_api.id
  resource_id             = aws_api_gateway_resource.add_resource.id
  http_method             = aws_api_gateway_method.post_method.http_method
  integration_http_method = "POST"
  type                    = "HTTP_PROXY"
  uri                     = "http://localhost:8080/add"  # Replace with your Go app endpoint
  request_parameters = {
    "integration.request.header.Content-Type" = "'application/json'"
  }
}

resource "aws_api_gateway_method_response" "post_method_response" {
  rest_api_id = aws_api_gateway_rest_api.example_api.id
  resource_id = aws_api_gateway_resource.add_resource.id
  http_method = aws_api_gateway_method.post_method.http_method
  status_code = "200"
}

resource "aws_api_gateway_integration_response" "app_integration_response" {
  rest_api_id = aws_api_gateway_rest_api.example_api.id
  resource_id = aws_api_gateway_resource.add_resource.id
  http_method = aws_api_gateway_method.post_method.http_method
  status_code = aws_api_gateway_method_response.post_method_response.status_code
  response_templates = {
    "application/json" = ""
  }
}

