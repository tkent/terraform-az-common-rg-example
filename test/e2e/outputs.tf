#
# Outputs have to be pulled through here to make them easily accessible to the
# terratest code.
#
output "tag_set" {
  value = module.common_rg.tag_set
}

output "resource_group" {
  value = module.common_rg.resource_group
}