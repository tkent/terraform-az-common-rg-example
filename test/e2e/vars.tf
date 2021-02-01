#
# Wondering why this is here? Unfortunately, I know of no easy way
# to
#

variable "location" {
  type = string
}
variable "projectName" {
  type = string
}
variable "projectShortName" {
  type = string
}
variable "billTo" {
  type = string
}
variable "extra_tags" {
  type = map(string)
}