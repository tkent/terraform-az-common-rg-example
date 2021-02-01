locals {
  #
  # Note, extra_tags that overlap with the predefined "protected" tags below
  # will be silently overwritten. See README.md for why specific validation
  # is not done here.
  #
  tag_set = merge(var.extra_tags, {
    location         = var.location,
    billTo           = var.billTo,
    projectName      = var.projectName,
    projectShortName = var.projectShortName
  })
}

resource "azurerm_resource_group" "resource_group" {
  name     = "${var.projectShortName}ResourceGroup"
  location = var.location
  tags     = local.tag_set
}