variable "location" {
  type        = string
  default     = "westus2"
  description = "The location for the resource group"
}

variable "projectName" {
  type        = string
  description = "The project name"
}

variable "projectShortName" {
  type        = string
  description = "A project identifier, must comply with resource group naming constraints"
}

variable "billTo" {
  type        = string
  description = "The bill-to department for the resource group"
}

variable "extra_tags" {
  type        = map(string)
  description = "An option set of additional tags for the resource."
  default     = {}
}