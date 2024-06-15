variable "apiname" {
  type    = string
  default = "rest"
}

variable "stagename" {
  type    = string
  default = "localdev"
}

variable "filename" {
  type    = string
  default = "lambda.zip"
}

variable "functionname" {
  type    = string
  default = "mylambda"
}

variable "lambdarole" {
  type    = string
  default = "testrole"
}
